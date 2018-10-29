package static

import (
	"ForumPublica/sde/models"
	"ForumPublica/sde/service"
	"archive/zip"
	"fmt"
)

var (
	Types         models.ZipTypes
	Blueprints    models.ZipBlueprints
	T2toT1        map[int32]int32
	BpoIdByTypeId map[int32]int32
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

	T2toT1 = make(map[int32]int32)
	BpoIdByTypeId = make(map[int32]int32)
	for bpoId, bpo := range Blueprints {
		if bpo.Invention != nil {
			for _, product := range bpo.Invention.Products {
				T2toT1[product.TypeId] = bpo.BlueprintTypeId
			}
		}
		if bpo.Manufacturing != nil && len(bpo.Manufacturing.Products) > 0 {
			// if len(bpo.Manufacturing.Products) == 0 {
			// 	fmt.Printf("%+v\n", Types[bpoId])
			// 	fmt.Printf("%+v\n", bpo.Manufacturing)
			// }
			BpoIdByTypeId[bpo.Manufacturing.Products[0].TypeId] = bpoId
		}
	}

}
