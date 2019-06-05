package service

import (
	"ForumPublica/sde/models"
	"fmt"

	"archive/zip"

	yaml "gopkg.in/yaml.v2"
)

//AddRegion add solar system from file
func AddRegion(f *zip.File, regionKey string, result *models.ZipRegions) {
	// fmt.Printf("%+v\n", f.FileHeader.Name)

	zipFile, zfErr := f.Open()
	if zfErr != nil {
		fmt.Println("f.Open():", zfErr)
		return
	}

	decoder := yaml.NewDecoder(zipFile)
	yamlData := models.RawRegion{}
	unmErr := decoder.Decode(&yamlData)
	if unmErr != nil {
		fmt.Println("Unmarshal:", unmErr)
	}

	region := models.ZipRegion{
		ID:        yamlData.RegionID,
		RegionKey: regionKey,
	}

	*result = append(*result, region)
}
