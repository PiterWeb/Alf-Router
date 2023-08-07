# Util 🧩 Templates

## Use html/template Out of the Box

### Usage :

Import it to your code 🔠

```go
    import (
        alfTemplates "github.com/PiterWeb/Alf-Router/utils/templates"
    )
```

Use it on your custom Routes

```go
// main.go

    useTemplate := alfTemplates.Templates("./templatesFolder", ".go.html")

    alf.App(&alf.AppConfig{
    	Port: "3000",
    	Routes: alf.CreateRouter([]alf.Route{
    		{
				Path: "/",
				Handle: func(ctx *alf.Ctx) error {
					return useTemplate(ctx, "index", "I am a Text")
				},
				Method: "get",
			},
    	}),
    })
```

```html
<!-- /templatesFolder/index.go.html -->

<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>Document</title>
    </head>
    <body>
        <h1>Hello World</h1>

        <h2>Text from go: {{.}}</h2>
    </body>
</html>
```
