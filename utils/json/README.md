# Util 🧩 JSON

## Send JSON Responses

### Usage :

Import it to your code 🔠
```go
    import (
        alfJSON "github.com/PiterWeb/Alf-Router/utils/json"
    )
```

Create your custom Struct or a Map
```go
    type MyCustomResponse struct {
        Message string
    }
```

Start sending JSON across your API Endpoints
```go
    err := alf.App(&alf.AppConfig{
    	Port: "3000",
    	Routes: alf.CreateRouter([]alf.Route{
    		{
				Path: "/api",
				Handle: func(ctx *alf.Ctx) error {

					_, err := alfJSON.JSON(ctx, MyCustomResponse{
                        Message: "Hello World",
                    })

					return err
				},
				Method: "get",
			},
    	}),
    })
```
