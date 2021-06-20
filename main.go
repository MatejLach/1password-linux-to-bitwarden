package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/MatejLach/1password-linux-to-bitwarden/decoder"
	"github.com/MatejLach/1password-linux-to-bitwarden/encoder"
)

var (
	exportFilePath string
	outputFilePath string
)

func init() {
	defaultOutputFile, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	defaultOutputFile = fmt.Sprintf("%s/bitwarden-import.json", defaultOutputFile)

	flag.StringVar(&exportFilePath, "1pass-backup-file-path", "", "The absolute path (location) of a 1Password .1pux backup file")
	flag.StringVar(&outputFilePath, "bitwarden-export-file-path", defaultOutputFile, "The absolute path to where you want to export the JSON to import into Bitwarden")
}

func main() {
	flag.Parse()

	if exportFilePath == "" || !filepath.IsAbs(exportFilePath) || !strings.HasSuffix(exportFilePath, ".1pux") {
		log.Fatalln("invalid file path; pass '--1pass-backup-file-path=/absolute/path/to/file.1pux'")
	}

	var onepass decoder.OnePassword
	err := onepass.Unmarshal1pux(exportFilePath)
	if err != nil {
		log.Fatal(err)
	}

	bitwarden, err := encoder.Encode1PasswordExportAsBitWardenImport(onepass)
	if err != nil {
		log.Fatal(err)
	}

	importB, err := bitwarden.EncodeBitwardenImport()
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(outputFilePath, importB, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
