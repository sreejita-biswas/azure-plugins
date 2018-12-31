package main

import (
	"flag"

	"github.com/sreejita-biswas/azure-plugins/authorization"
)

func main() {
	authorization.GetParameterValues()
	flag.Parse()
	_, success := authorization.Authorize()
	if !success {
		return
	}
}
