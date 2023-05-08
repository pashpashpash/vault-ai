package postapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/pashpashpash/vault/errorlist"
	"github.com/pashpashpash/vault/form"
)

// Util: does grunt work of decoding + verifying form + writing errors back
func FormParseVerify(form form.Form, name string, w http.ResponseWriter, r *http.Request) errorlist.Errors {
	if err := r.ParseMultipartForm(80000); err != nil {
		log.Printf("[%s] Error parsing form POST\n", name)
		errs := errorlist.NewSingleError("parsePost", err)

		bytes, err := json.Marshal(errs)
		if err != nil {
			log.Println("Dev fucked up bad, errors didn't JSON encode")
			return nil
		}

		http.Error(w, fmt.Sprintf("%s", bytes), http.StatusBadRequest)
		return errs
	}

	if err := schema.NewDecoder().Decode(form, r.Form); err != nil {
		log.Printf("[%s] Decode FAIL: %s %s\n", name, form, err)
		errs := errorlist.NewSingleError("decode", err)

		bytes, err := json.Marshal(errs)
		if err != nil {
			log.Println("Dev fucked up bad, errors didn't JSON encode")
			return nil
		}

		http.Error(w, fmt.Sprintf("%s", bytes), http.StatusBadRequest)
		return errs
	}

	errs := form.Validate()
	if len(errs) > 0 {
		log.Printf("[%s] Validation FAIL: %s %s\n", name, form, errs)

		bytes, err := json.Marshal(errs)
		if err != nil {
			log.Println("Dev fucked up bad, errors didn't JSON encode")
		}

		http.Error(w, fmt.Sprintf("%s", bytes), http.StatusBadRequest)
		return errs
	}

	log.Printf("[%s] Validated: %s\n", name, form)
	return nil
}
