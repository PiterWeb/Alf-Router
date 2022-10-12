package alf

import (
	"github.com/pquerna/ffjson/ffjson"
)

func JSON(data interface{}) []byte { // utility function to convert a struct or map to json([]byte)

	json, err := ffjson.Marshal(data)

	if err != nil {
		showInternalError("alf.JSON utility (ffjson): " + err.Error())
	}

	return json

}