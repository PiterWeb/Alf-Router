# 🦌 ALF Router (API Like Flutter)

### Description 

👉 This is a framework / router built on ⬆ top of the fasthttp package. Alf relies on scalability and his structure wich is  similar to Flutter projects 📴

## Purpouse 

📖 Learn 

 - Explore the fasthttp package ⏭
 - Publish my own package 📦
 - Learn more deeply how a web server works 🌐

## Technologies used 📘

 - Go (Golang)

#### External Packages  📦:

 1. [fasthttp](github.com/valyala/fasthttp) (http ☁)
 2. [ffjson](github.com/pquerna/ffjson/ffjson) (parse interfaces to json fast)

## Docs

### Prerequisites 📌

 - [Go 1.18](https://go.dev/) 

### Set up your first project 💻

Download the package ⬇

    go get github.com/PiterWeb/Alf-Router

Import it into your code 🔠

    import (
	    alf "github.com/PiterWeb/Alf-Router"
    )

Write this simple structure

    err := alf.App(
	    alf.AppConfig(
		    Port: "8080",
		    Routes: alf.CreateRouter([]alf.Route{
			    {
				    Path: "/",
				    Method: "GET",
				    Handle: func(ctx *alf.Ctx) {
					    ctx.WriteString("Hello World!")
					},
				},
			NotFound: func(ctx *alf.Ctx) {
				ctx.SetContentType("application/json")
				ctx.Write(alf.JSON(map[string]string{
					"error":"not found"
				}))
			},
			Middleware: []alf.Middleware{},
			Headers: []alf.Header{},
    )
    
    if err != nil {
	    panic(err)
	} 
	
Now you have setup 🔨 an index path  and the 404 route 📁 
