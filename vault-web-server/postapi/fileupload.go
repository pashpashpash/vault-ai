package postapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	openai "github.com/sashabaranov/go-openai"
)

type UploadResponse struct {
	Message             string            `json:"message"`
	NumFilesSucceeded   int               `json:"num_files_succeeded"`
	NumFilesFailed      int               `json:"num_files_failed"`
	SuccessfulFileNames []string          `json:"successful_file_names"`
	FailedFileNames     map[string]string `json:"failed_file_names"`
}

const MAX_FILE_SIZE int64 = 3 << 20         // 3 MB
const MAX_TOTAL_UPLOAD_SIZE int64 = 3 << 20 // 3 MB

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	// Limit the request body size
	r.Body = http.MaxBytesReader(w, r.Body, MAX_TOTAL_UPLOAD_SIZE)

	err := r.ParseMultipartForm(MAX_TOTAL_UPLOAD_SIZE) // Maximum upload of 3 MB
	if err != nil {
		if err == http.ErrMissingBoundary || err == http.ErrNotMultipart || err == http.ErrNotSupported {
			log.Println("[UploadHandler ERR] Error parsing multipart form:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("[UploadHandler ERR] Request body size exceeds the limit:", err)
		http.Error(w, "Request body size exceeds the limit", http.StatusRequestEntityTooLarge)
		return
	}

	files := r.MultipartForm.File["files"]
	uuid := r.FormValue("uuid") // Get the UUID from the form data
	userProvidedOpenApiKey := r.FormValue("apikey")

	log.Println("[UploadHandler] UUID=", uuid)

	clientToUse := DEFAULT_OPENAI_CLIENT
	if userProvidedOpenApiKey != "" {
		log.Println("[UploadHandler] Using provided custom API key:", userProvidedOpenApiKey)
		clientToUse = openai.NewClient(userProvidedOpenApiKey)
	}

	responseData := UploadResponse{
		SuccessfulFileNames: make([]string, 0),
		FailedFileNames:     make(map[string]string),
	}

	for _, file := range files {
		fileName := file.Filename

		if file.Size > MAX_FILE_SIZE {
			errMsg := fmt.Sprintf("File size exceeds the %d bytes limit", MAX_FILE_SIZE)
			log.Println("[UploadHandler ERR]", errMsg, fileName)
			responseData.NumFilesFailed++
			responseData.FailedFileNames[fileName] = errMsg
			continue
		}

		// Read the file in memory
		f, err := file.Open()
		if err != nil {
			errMsg := "Error opening file"
			log.Println("[UploadHandler ERR]", errMsg, err)
			responseData.NumFilesFailed++
			responseData.FailedFileNames[fileName] = errMsg
			continue
		}
		defer f.Close()

		// Get the file name, MIME type, and first 32 characters of the contents
		fileType := file.Header.Get("Content-Type")
		fileContent := ""
		filePreview := ""

		// Check if the file is a PDF
		if fileType == "application/pdf" {
			fileContent, err = extractTextFromPDF(f, file.Size)
			if err != nil {
				errMsg := "Error extracting text from PDF"
				log.Println("[UploadHandler ERR]", errMsg, err)
				responseData.NumFilesFailed++
				responseData.FailedFileNames[fileName] = errMsg
				continue
			}
		} else {
			fileContent, err = getTextFromFile(f)
			if err != nil {
				errMsg := "Error reading file"
				log.Println("[UploadHandler ERR]", errMsg, err)
				responseData.NumFilesFailed++
				responseData.FailedFileNames[fileName] = errMsg
				continue
			}
		}

		if len(fileContent) > 32 {
			filePreview = fileContent[:32]
		}
		log.Printf("File Name: %s, File Type: %s, File Content (first 32 characters): %s\n", fileName, fileType, filePreview)

		// Process the fileBytes into embeddings and store in Pinecone here
		chunks, err := CreateChunks(fileContent, fileName)
		if err != nil {
			errMsg := "Error chunking file"
			log.Println("[UploadHandler ERR]", errMsg, err)
			responseData.NumFilesFailed++
			responseData.FailedFileNames[fileName] = errMsg
			continue
		}

		embeddings, err := getEmbeddings(clientToUse, chunks, 100, openai.AdaEmbeddingV2)
		if err != nil {
			errMsg := fmt.Sprintf("Error getting embeddings: %v", err)
			log.Println("[UploadHandler ERR]", errMsg)
			responseData.NumFilesFailed++
			responseData.FailedFileNames[fileName] = errMsg
			continue
		}
		fmt.Printf("Total chunks: %d\n", len(chunks))
		fmt.Printf("Total embeddings: %d\n", len(embeddings))
		fmt.Printf("Embeddings length: %d\n", len(embeddings[0]))

		// Call the upsertEmbeddingsToPinecone function
		err = upsertEmbeddingsToPinecone(embeddings, chunks, uuid)
		if err != nil {
			errMsg := fmt.Sprintf("Error upserting embeddings to Pinecone: %v", err)
			log.Println("[UploadHandler ERR]", errMsg)
			responseData.NumFilesFailed++
			responseData.FailedFileNames[fileName] = errMsg
			continue
		}

		log.Println("Successfully added pinecone embeddings!")

		responseData.NumFilesSucceeded++
		responseData.SuccessfulFileNames = append(responseData.SuccessfulFileNames, fileName)
	}

	if responseData.NumFilesFailed > 0 {
		responseData.Message = "Some files failed to upload and process"
	} else {
		responseData.Message = "All files uploaded and processed successfully"
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(responseData)
	if err != nil {
		log.Println("[UploadHandler ERR] Error writing json response", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
