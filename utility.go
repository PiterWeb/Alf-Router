package alf

import (
	"fmt"

	"github.com/pquerna/ffjson/ffjson"
)

func JSON(data interface{}) []byte { // utility function to convert a struct or map to json([]byte)

	json, err := ffjson.Marshal(data)

	if err != nil {
		go showInternalError(fmt.Sprintf("alf.JSON utility (ffjson): %s", err.Error()))
	}

	return json

}
