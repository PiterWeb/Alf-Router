package ts

import (
	"os"
	"strings"

	"github.com/gzuidhof/tygo/tygo"
	"golang.org/x/mod/modfile"
)

type TS_config struct {
	Packages     []string
	OutputFolder string
}

func Init_ts(ts_config TS_config) error {

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
			{Path: mainPkg, OutputPath: ts_config.OutputFolder + "/main.ts"},
		},
	}

	for _, pkg := range ts_config.Packages {

		if strings.Contains(pkg, mainPkg) {

			config.Packages = append(config.Packages, &tygo.PackageConfig{
				Path:       pkg,
				OutputPath: ts_config.OutputFolder + "/" + strings.Split(pkg, mainPkg)[1] + ".ts",
			})

		} else {

			config.Packages = append(config.Packages, &tygo.PackageConfig{
				Path:       pkg,
				OutputPath: ts_config.OutputFolder + "/" + pkg + ".ts",
			})
		}

	}

	gen := tygo.New(config)
	return gen.Generate()

}
