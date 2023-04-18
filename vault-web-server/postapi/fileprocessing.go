package postapi

import "strings"

type Chunk struct {
	Start int
	End   int
	Title string
	Text  string
}

func CreateChunks(fileContent string, window int, stride int, title string) []Chunk {
	sentences := strings.Split(fileContent, ".") // assuming sentences end with a period
	newData := make([]Chunk, 0)

	for i := 0; i < len(sentences)-window; i += stride {
		iEnd := i + window
		text := strings.Join(sentences[i:iEnd], ". ")
		start := 0
		end := 0

		if i > 0 {
			start = len(strings.Join(sentences[:i], ". ")) + 2 // +2 for the period and space
		}

		end = len(strings.Join(sentences[:iEnd], ". "))

		newData = append(newData, Chunk{
			Start: start,
			End:   end,
			Title: title,
			Text:  text,
		})
	}

	return newData
}
