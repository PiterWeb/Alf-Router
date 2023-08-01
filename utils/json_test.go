package utils

import (
	"testing"

	alf "github.com/PiterWeb/Alf-Router"
)

func TestJSON(t *testing.T) {

	value := map[string]string{
		"test": "success",
	}

	_ , err := JSON(&alf.Ctx{}, value)

	if (err != nil) {
		t.Error("JSON Test Failed with err: " + err.Error())
	}

	t.Log(value)

}
