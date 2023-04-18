package validator

import (
	"net/mail"
	"regexp"

	"github.com/pashpashpash/vault/errorlist"
)

const SPECIAL_DELETE_ASCII string = "\x10"

type Email struct {
	EmailAddr string `schema:"email"`
}

func (me *Email) Validate(errs errorlist.Errors) {
	if me.EmailAddr == SPECIAL_DELETE_ASCII {
		return
	}
	var emailRegex = regexp.MustCompile("^.+@.+\\..+$")

	// Just does basic parser validation, note that foo@localhost is valid
	if _, err := mail.ParseAddress(me.EmailAddr); err != nil {
		errs["email"] = errorlist.NewError("email did not parse")
	}

	// Does SUPER basic regex validation that email must have @ and .
	if !emailRegex.MatchString(me.EmailAddr) {
		errs["email"] = errorlist.NewError("email format invalid")
	}

	if len(me.EmailAddr) > 320 {
		errs["email"] = errorlist.NewError("email format invalid: too long")
	}
}

func ValidateEmail(errs errorlist.Errors, me *Email) {
	var emailRegex = regexp.MustCompile("^.+@.+\\..+$")

	// Just does basic parser validation, note that foo@localhost is valid
	if _, err := mail.ParseAddress(me.EmailAddr); err != nil {
		errs["email"] = errorlist.NewError("email did not parse")
	}

	// Does SUPER basic regex validation that email must have @ and .
	if !emailRegex.MatchString(me.EmailAddr) {
		errs["email"] = errorlist.NewError("email format invalid")
	}

	if len(me.EmailAddr) > 320 {
		errs["email"] = errorlist.NewError("email format invalid: too long")
	}
}
