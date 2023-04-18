package validator

import (
	"strings"

	"github.com/pashpashpash/vault/errorlist"
)

type Validator interface {
	Validate(map[string]error)
}

func CheckNotEmpty(input, name string, errs errorlist.Errors) {
	if len(strings.TrimSpace(input)) == 0 {
		errs[name] = errorlist.NewError("cannot be blank")
	}
}
