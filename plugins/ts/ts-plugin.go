package ts

import (
	"os"
	"strings"

	"github.com/gzuidhof/tygo/tygo"
	"golang.org/x/mod/modfile"
)

type TS_plugin struct {
	Packages     []string
	OutputFolder string
}

func (plugin TS_plugin) Init_plugin() error {

	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	modFile, err := os.ReadFile(cwd + "/go.mod")

	if err != nil {
		return err
	}

	mainPkg := modfile.ModulePath(modFile)

	config := &tygo.Config{
		Packages: []*tygo.PackageConfig{
			{Path: mainPkg, OutputPath: plugin.OutputFolder + "/main.ts"},
		},
	}

	for _, pkg := range plugin.Packages {

		if strings.Contains(pkg, mainPkg) {

			config.Packages = append(config.Packages, &tygo.PackageConfig{
				Path:       pkg,
				OutputPath: plugin.OutputFolder + "/" + strings.Split(pkg, mainPkg)[1] + ".ts",
			})

		} else {

			config.Packages = append(config.Packages, &tygo.PackageConfig{
				Path:       pkg,
				OutputPath: plugin.OutputFolder + "/" + pkg + ".ts",
			})
		}

	}

	gen := tygo.New(config)
	return gen.Generate()

}
