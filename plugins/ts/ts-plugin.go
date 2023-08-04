package ts

import (
	"github.com/gzuidhof/tygo/tygo"
)

type TS_config struct {
	Packages     []string
	OutputFolder string
}

func Init_ts(ts_config TS_config) error {

	config := &tygo.Config{
		Packages: []*tygo.PackageConfig{},
	}

	for _, pkg := range ts_config.Packages {
		config.Packages = append(config.Packages, &tygo.PackageConfig{
			Path:       pkg,
			OutputPath: ts_config.OutputFolder + "/" + pkg + ".ts",
		})

	}

	gen := tygo.New(config)
	return gen.Generate()

}
