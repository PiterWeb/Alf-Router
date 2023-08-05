# 🦌 ALF (API Like Flutter) Router  [Still Experimental]

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

- [x] Router 💨
- [x] Concurrent Route Setup ⌚
- [x] Send JSON Responses [(docs)](https://github.com/PiterWeb/Alf-Router/blob/master/utils/JSON.md)
- [x] Plugins 🧩 : 
	- [x] Generate Types for TS (tygo pkg) [(docs)](https://github.com/PiterWeb/Alf-Router/blob/master/plugins/ts/README.md)
- [ ] Zero config HTML/Go Templates

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

	
Now you have setup 🔨 an index route
