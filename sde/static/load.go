package static

import (
	"ForumPublica/sde/models"
	"ForumPublica/sde/service"
	"archive/zip"
	"fmt"
)

var (
	Types      models.ZipTypes
	Blueprints models.ZipBlueprints
)

func LoadJSONs(fileName string) {
	r, zipErr := zip.OpenReader(fileName)
	if zipErr != nil {
		fmt.Println("zip.OpenReader:", zipErr)
		return
	}
	defer r.Close()

	for _, f := range r.File {
		// fmt.Printf("%+v\n", f.FileHeader.Name)
		if f.FileHeader.Name == "sde/fsd/typeIDs.yaml" {
			Types = *service.ImportTypes(f)
		}
		if f.FileHeader.Name == "sde/fsd/blueprints.yaml" {
			Blueprints = *service.ImportBlueprints(f)
		}
	}
}
