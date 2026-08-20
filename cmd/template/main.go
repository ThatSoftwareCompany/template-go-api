package main

import (
	"flag"
	"fmt"
	"os"

	platformtemplate "github.com/ThatSoftwareCompany/template-go-api/internal/platform/template"
)

func main() {
	command := flag.String("command", "validate", "template command: validate")
	manifestPath := flag.String("manifest", ".template/manifest.json", "path to the template manifest")
	flag.Parse()

	switch *command {
	case "validate":
		manifest, err := platformtemplate.Load(*manifestPath)
		if err == nil {
			err = manifest.Validate()
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "template manifest validation failed")
			os.Exit(1)
		}
		fmt.Println("template manifest is valid")
	default:
		fmt.Fprintln(os.Stderr, "-command must be validate")
		os.Exit(2)
	}
}
