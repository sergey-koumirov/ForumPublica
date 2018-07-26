package service

import (
	"ForumPublica/sde/models"
	"fmt"

	"archive/zip"

	yaml "gopkg.in/yaml.v2"
)

func ImportTypes(f *zip.File) {
	if f.FileHeader.Name == "sde/fsd/typeIDs.yaml" {
		fmt.Printf("%+v\n", f.FileHeader.Name)

		zipFile, zfErr := f.Open()
		if zfErr != nil {
			fmt.Println("f.Open():", zfErr)
			return
		}

		decoder := yaml.NewDecoder(zipFile)
		jsonData := make(models.RawTypes)
		unmErr := decoder.Decode(&jsonData)
		if unmErr != nil {
			fmt.Println("Unmarshal:", unmErr)
			return
		}
		//todo insert to db
		for key, value := range jsonData {
			fmt.Println(key, value)
		}

	}
}
