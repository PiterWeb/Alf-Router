package json

import (
	"fmt"

	alf "github.com/PiterWeb/Alf-Router"
	misc "github.com/PiterWeb/Alf-Router/errors"
	"github.com/pquerna/ffjson/ffjson"
)

func Marshall(ctx *alf.Ctx, data interface{}) (int, error) { // utility function to convert a struct or map to json([]byte)

	json, err := ffjson.Marshal(data)

	if err != nil {
		go misc.ShowInternalError(fmt.Sprintf("alf.JSON utility (ffjson): %s", err.Error()))
	}

	return ctx.Write(json)

}

func Unmarshall(data []byte, v interface{}) error {
	return ffjson.Unmarshal(data, v)
}