package service

import (
	"ForumPublica/sde/models"
	"fmt"

	"archive/zip"

	yaml "gopkg.in/yaml.v2"
)

func ImportTypes(f *zip.File) *models.ZipTypes {
	fmt.Printf("%+v\n", f.FileHeader.Name)

	zipFile, zfErr := f.Open()
	if zfErr != nil {
		fmt.Println("f.Open():", zfErr)
		return nil
	}

	decoder := yaml.NewDecoder(zipFile)
	jsonData := make(models.RawTypes)
	unmErr := decoder.Decode(&jsonData)
	if unmErr != nil {
		fmt.Println("Unmarshal:", unmErr)
		return nil
	}
	//todo insert to db
	result := make(models.ZipTypes)

	for key, value := range jsonData {
		result[key] = models.ZipType{
			Id:          key,
			GroupId:     value.GroupId,
			Name:        value.Names["en"],
			PortionSize: value.PortionSize,
			Published:   value.Published,
			Volume:      value.Volume,
		}
	}

	return &result

}
