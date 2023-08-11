package json

import (
	"testing"

	alf "github.com/PiterWeb/Alf-Router"
)

func TestMarshall(t *testing.T) {

	value := map[string]string{
		"test": "success",
	}

	_, err := Marshall(&alf.Ctx{}, value)

	if err != nil {
		t.Error("JSON Test Failed with err: " + err.Error())
	}

	t.Log(value)

}

func TestUnmarshall(t *testing.T) {

	value := []byte("{\"test\": \"success\"}")

	var valueParsed map[string]string

	err := Unmarshall(value, &valueParsed)

	if err != nil {
		t.Error("JSON Test Failed with err: " + err.Error())
	}

	t.Log(valueParsed)

}