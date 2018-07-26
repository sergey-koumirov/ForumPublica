package service

import (
	"ForumPublica/sde/models"
	"fmt"

	"archive/zip"

	yaml "gopkg.in/yaml.v2"
)

func ImportBlueprints(f *zip.File) *models.ZipBlueprints {
	fmt.Printf("%+v\n", f.FileHeader.Name)

	zipFile, zfErr := f.Open()
	if zfErr != nil {
		fmt.Println("f.Open():", zfErr)
		return nil
	}

	decoder := yaml.NewDecoder(zipFile)
	jsonData := make(models.RawBlueprints)
	unmErr := decoder.Decode(&jsonData)
	if unmErr != nil {
		fmt.Println("Unmarshal:", unmErr)
		return nil
	}

	result := make(models.ZipBlueprints)

	for key, value := range jsonData {
		// fmt.Println(key, value)

		temp := models.ZipBlueprint{
			BlueprintTypeId:    key,
			MaxProductionLimit: value.MaxProductionLimit,
		}

		temp.Copying = extractActivity("copying", value.Activities)
		temp.Manufacturing = extractActivity("manufacturing", value.Activities)
		temp.Invention = extractActivity("invention", value.Activities)
		temp.ResearchMaterial = extractActivity("research_material", value.Activities)
		temp.ResearchTime = extractActivity("research_time", value.Activities)

		result[key] = temp
	}

	return &result

}

func extractActivity(activity string, value map[string]models.RawActivity) *models.RawActivity {
	raw, ex := value[activity]
	if ex {
		return &raw
	}
	return nil
}
