package form

import "github.com/pashpashpash/vault/errorlist"

type Form interface {
	Validate() errorlist.Errors
	String() string
}
