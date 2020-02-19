package service

import (
	"ForumPublica/sde/models"
	"fmt"

	"archive/zip"

	yaml "gopkg.in/yaml.v2"
)

//ImportGroups load Groups from file
func ImportGroups(f *zip.File) *models.ZipGroups {
	fmt.Printf("%+v\n", f.FileHeader.Name)

	zipFile, zfErr := f.Open()
	if zfErr != nil {
		fmt.Println("f.Open():", zfErr)
		return nil
	}

	decoder := yaml.NewDecoder(zipFile)
	jsonData := make(models.RawGroups)
	unmErr := decoder.Decode(&jsonData)
	if unmErr != nil {
		fmt.Println("Unmarshal:", unmErr)
		return nil
	}

	result := make(models.ZipGroups)

	for key, value := range jsonData {
		result[key] = models.ZipGroup{
			ID:         key,
			CategoryID: value.CategoryID,
			Name:       value.Names["en"],
			Published:  value.Published,
		}
	}

	return &result

}
