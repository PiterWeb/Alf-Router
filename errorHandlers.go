package alf

import (
	"github.com/pterm/pterm"
)

func showWarning(err string) {

	pterm.Warning.Println(err)

}

func showError(err string) {

	pterm.Error.Println(err)

}

func showInternalError(err string) {

	pterm.Fatal.WithFatal(true).Println(err)

}
