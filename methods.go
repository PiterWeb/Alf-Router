package alf

import (
	"strings"
)

func (m Method) valid() bool { // return if the method is valid

	switch m.string() {
	case "GET", "POST", "PUT", "DELETE":
		return true
	default:
		return false
	}

}

func (m Method) string() string {

	return strings.ToUpper(string(m))

}