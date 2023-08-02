package alf

import (
	"testing"
)

func TestValid(t *testing.T) {

	value := "hello world"

	valid := Method(value).valid()

	if valid {
		t.Errorf("%s cannot be a Method", value)
	} else {

		t.Logf("%s skipped successfully", value)
	}

	value = "get"

	valid = Method(value).valid()

	if !valid {
		t.Errorf("%s is a valid method but it fails", value)
	} else {

		t.Logf("%s didn't had to skip", value)
	}


}
