# Plugin 🧩 TS

## Export your golang types to TS for your frontend

#### Credits to [Tygo PKG](https://github.com/gzuidhof/tygo)

### Usage :

Import it to your code 🔠

```go
    import (
        tspl "github.com/PiterWeb/Alf-Router/plugins/ts"
    )
```

Use it on your start point and change Packages with the names of the modules where are the types you want. For default all public types on the default module will be already included

```go
    err := alf.App(&alf.AppConfig{
    	Port: "3000",
    	Routes: alf.CreateRouter([]alf.Route{
    		{
    			...
    		},
    	}),
    	BeforeInit: func(ac *alf.AppConfig) {
            tspl.Init_ts(tspl.TS_config{
    			Packages: []string{
                    "github.com/exampleUser/myProject/submodule"
                },
    			OutputFolder: "./ts-types",
    		})
    	},
    })
```
