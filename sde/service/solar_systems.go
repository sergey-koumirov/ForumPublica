package service

import (
	"ForumPublica/sde/models"
	"fmt"

	"archive/zip"

	yaml "gopkg.in/yaml.v2"
)

//AddSolarSystem add solar system from file
func AddSolarSystem(f *zip.File, regionKey string, result *models.ZipSolarSystemsList) {
	// fmt.Printf("%+v\n", f.FileHeader.Name)

	zipFile, zfErr := f.Open()
	if zfErr != nil {
		fmt.Println("f.Open():", zfErr)
		return
	}

	decoder := yaml.NewDecoder(zipFile)
	yamlData := models.RawSolarSystem{}
	unmErr := decoder.Decode(&yamlData)
	if unmErr != nil {
		fmt.Println("Unmarshal:", unmErr)
	}

	system := models.ZipSolarSystem{
		ID:        yamlData.SolarSystemID,
		Security:  yamlData.Security,
		RegionKey: regionKey,
	}

	*result = append(*result, system)

}
