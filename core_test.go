package alf

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

const (
	// Host name of the HTTP Server
	Host = "localhost"
	// Port of the HTTP Server
	Port = "8080"
)

const (
	BodyResponse = "Hello World"
)

func TestRawGet(t *testing.T) {

	generateApp(t)

	body := getRequest(t, "/raw")

	t.Log(string(body))

	if string(body) != BodyResponse {
		t.FailNow()
	}

}

func TestRawPost(t *testing.T) {

	generateApp(t)

	body := postRequest(t, "/api/postRaw", []byte(BodyResponse))

	t.Log(string(body))

	if string(body) != BodyResponse {
		t.FailNow()
	}

}

func generateApp(t *testing.T) {

	err := App(&AppConfig{
		Routes: CreateRouter([]Route{
			{
				Path:   "/raw",
				Method: "get",
				Handle: func(ctx *Ctx) error {
					ctx.WriteString(BodyResponse)
					return nil
				},
			},
			{
				Path:   "/api",
				Method: "get",
				Handle: func(ctx *Ctx) error {
					ctx.WriteString("Working 💪")
					return nil
				},
				Children: []Route{
					{
						Path:   "/postRaw",
						Method: "post",
						Handle: func(ctx *Ctx) error {

							ctx.Write(ctx.Request.Body())

							return nil

						},
					},
				},
			},
		}),
	})

	if err != nil {
		t.Error(err)
	}

}

func getRequest(t *testing.T, path string) []byte {

	resp, err := http.Get("http://" + Host + ":" + Port + path)

	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Error(err)
	}

	return body

}

func postRequest(t *testing.T, path string, body []byte) []byte {

	resp, err := http.Post("http://"+Host+":"+Port+path, "text/plain", bytes.NewReader(body))

	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Error(err)
	}

	return respBody

}
