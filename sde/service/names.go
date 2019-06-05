package service

import (
	"ForumPublica/sde/models"
	"fmt"

	"archive/zip"

	yaml "gopkg.in/yaml.v2"
)

//ImportNames load Types from file
func ImportNames(f *zip.File) *map[int64]string {
	fmt.Printf("%+v\n", f.FileHeader.Name)

	zipFile, zfErr := f.Open()
	if zfErr != nil {
		fmt.Println("f.Open():", zfErr)
		return nil
	}

	decoder := yaml.NewDecoder(zipFile)
	jsonData := make([]models.RawName, 0)
	unmErr := decoder.Decode(&jsonData)
	if unmErr != nil {
		fmt.Println("Unmarshal:", unmErr)
		return nil
	}

	result := make(map[int64]string)
	for _, value := range jsonData {
		result[value.ItemID] = value.ItemName
	}

	return &result

}
