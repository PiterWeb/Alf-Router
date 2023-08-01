package misc

import (
	"github.com/pterm/pterm"
)

func ShowWarning(err string) {

	pterm.Warning.Println(err)

}

func ShowError(err string) {

	pterm.Error.Println(err)

}

func ShowInternalError(err string) {

	pterm.Fatal.WithFatal(true).Println(err)

}
