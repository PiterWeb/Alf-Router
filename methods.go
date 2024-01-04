package alf

import (
	"strings"
)

const (
	GET     method = "GET"
	POST    method = "POST"
	DELETE  method = "DELETE"
	PUT     method = "PUT"
	PATCH   method = "PATCH"
	HEAD    method = "HEAD"
	OPTIONS method = "OPTIONS"
)

func (m method) valid() bool { // return if the method is valid

	switch m.string() {
	case "GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS":
		return true
	default:
		return false
	}

}

func (m method) string() string {

	return strings.ToUpper(string(m))

}
