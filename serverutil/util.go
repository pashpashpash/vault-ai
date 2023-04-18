package serverutil

import (
	"path/filepath"
)

func WebAbs(relpath string) string {
	abspath, err := filepath.Abs("./static/" + relpath)

	if err != nil {
		panic("err converting to absolute path " + relpath)
	}

	return abspath
}
