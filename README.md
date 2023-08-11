# 🦌 ALF (API Like Flutter) Router

#### 🔴 This library is actually well tested but the API may change with the time introducing breaking changes

### Description 

👉 This is a micro-framework / router built on ⬆ top of the fasthttp package. Alf relies on scalability and his simple structure wich is  similar to Flutter projects 📴

✨ Inspired by Flutter & [Fiber](https://github.com/gofiber/fiber)

## Purpouse 

👷‍♂️ Make a router to start faster and simpler my backend projects

📖 Learn 

 - Explore the fasthttp package ⏭
 - Publish my own package 📦
 - Learn more deeply how a web server works 🌐

## Technologies used 📘

 - Go (Golang)

#### Core External Packages  📦:

 1. [fasthttp](github.com/valyala/fasthttp) (http ☁)
 2. [ffjson](github.com/pquerna/ffjson/ffjson) (parse interfaces to json fast)
 3. [pterm](github.com/pterm/pterm) (show info, errors & warnings)

## Features :

- [x] Fast Router 💨
- [x] Concurrent Route Setup ⌚
- [x] Send JSON Responses [(docs)](/utils/json/README.md)
- [x] HTML/Go Templates Out of the BOX ✨ [(docs)](/utils/templates/README.md)
- [x] Plugins 🧩 : 
	- [x] Generate Types for TS (tygo pkg) [(docs)](/plugins/ts/README.md)

## Docs

### Prerequisites 📌

 - [Go 1.18](https://go.dev/) 

### Set up your first project 💻

Download the package ⬇
```shell
go get github.com/PiterWeb/Alf-Router
```

Import it into your code 🔠

```go
    import (
	    alf "github.com/PiterWeb/Alf-Router"
    )
```

Write this simple structure

```go
    err := alf.App(&alf.AppConfig{
		Port: "3000",
		Routes: alf.CreateRouter([]alf.Route{
			{
				Path: "/",
				Handle: func(ctx *alf.Ctx) error {
					_, err := ctx.WriteString("Hello World")
					return err
				},
				Method: "get",
			},
		}),
	})

	if err != nil {
		panic(err)
	}
```
	
Now you have setup 🔨 an index route
