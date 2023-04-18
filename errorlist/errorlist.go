package errorlist

import "bytes"

// ------ Error Type ------ //
// (Works with json.Marshal)

type Error string

func NewError(input string) Error {
	return Error(input)
}

func (me Error) String() string {
	return string(me)
}

func (me Error) Error() string {
	return me.String()
}

// ----- Errors Map ----- //

type Errors map[string]Error

func New() Errors {
	return make(map[string]Error)
}

func NewSingleError(name string, err error) Errors {
	var me = make(map[string]Error)
	me[name] = NewError(err.Error())
	return me
}

func (me Errors) String() string {
	var buffer bytes.Buffer

	for k, v := range me {
		buffer.WriteString("\n")
		buffer.WriteString("          - ")
		buffer.WriteString(k)
		buffer.WriteString(": ")
		buffer.WriteString(v.Error())
	}

	return buffer.String()
}

func (me Errors) Error() string {
	return me.String()
}
