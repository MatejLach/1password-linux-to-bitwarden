package decoder

import (
	"archive/zip"
	"encoding/json"
	"strings"
)

func (onp *OnePassword) Unmarshal1pux(srcFilePath string) error {
	zipReader, err := zip.OpenReader(srcFilePath)
	if err != nil {
		return err
	}

	defer zipReader.Close()

	for _, file := range zipReader.File {
		if strings.HasSuffix(file.Name, ".data") {
			br, err := file.Open()
			if err != nil {
				return err
			}

			jsonDecoder := json.NewDecoder(br)
			err = jsonDecoder.Decode(onp)

			br.Close()

			if err != nil {
				return err
			}
		}
	}

	return nil
}
