package static

import (
	"ForumPublica/sde/models"
	"ForumPublica/sde/service"
	"archive/zip"
	"fmt"
	"strings"
)

var (
	//Types map with all types description
	Types models.ZipTypes

	//Blueprints map with all blueprints description
	Blueprints models.ZipBlueprints

	//T2toT1 get T1 bpo id by T2 bpo id
	T2toT1 map[int32]int32

	//BpoIDByTypeID get bpo id by product id
	BpoIDByTypeID map[int32]int32

	//SolarSystemsList array with K-space solar systems
	SolarSystemsList models.ZipSolarSystems

	//RegionsList array with K-space regions
	RegionsList models.ZipRegions

	//Regions hash with K-space regions
	Regions map[int64]models.ZipRegion
)

//LoadJSONs load data from zipped jsons
func LoadJSONs(fileName string) {
	r, zipErr := zip.OpenReader(fileName)
	if zipErr != nil {
		fmt.Println("zip.OpenReader:", zipErr)
		return
	}
	defer r.Close()

	var names map[int64]string
	SolarSystemsList = make(models.ZipSolarSystems, 0)

	for _, f := range r.File {
		// fmt.Printf("%+v\n", f.FileHeader.Name)
		if f.FileHeader.Name == "sde/fsd/typeIDs.yaml" {
			Types = *service.ImportTypes(f)
		}

		if f.FileHeader.Name == "sde/fsd/blueprints.yaml" {
			Blueprints = *service.ImportBlueprints(f)
		}

		if f.FileHeader.Name == "sde/bsd/invNames.yaml" {
			names = *service.ImportNames(f)
		}

		if strings.HasPrefix(f.FileHeader.Name, "sde/fsd/universe/eve") && strings.HasSuffix(f.FileHeader.Name, "region.staticdata") {
			regionKey := strings.Split(f.FileHeader.Name, "/")[4]
			service.AddRegion(f, regionKey, &RegionsList)
		}

		if strings.HasPrefix(f.FileHeader.Name, "sde/fsd/universe/eve") && strings.HasSuffix(f.FileHeader.Name, "solarsystem.staticdata") {
			regionKey := strings.Split(f.FileHeader.Name, "/")[4]
			service.AddSolarSystem(f, regionKey, &SolarSystemsList)
		}

	}

	Regions := make(map[int64]models.ZipRegion)
	regionsByKey := make(map[string]models.ZipRegion)

	for i, r := range RegionsList {
		RegionsList[i].Name = names[r.ID]
		Regions[r.ID] = RegionsList[i]
		regionsByKey[r.RegionKey] = RegionsList[i]
	}
	fmt.Println("Loaded regions: ", len(RegionsList))

	for i, s := range SolarSystemsList {
		SolarSystemsList[i].Name = names[s.ID]
		r, _ := regionsByKey[s.RegionKey]
		SolarSystemsList[i].Region = &r
	}
	fmt.Println("Loaded solar systems: ", len(SolarSystemsList))

	T2toT1 = make(map[int32]int32)
	BpoIDByTypeID = make(map[int32]int32)
	for bpoID, bpo := range Blueprints {
		if bpo.Invention != nil {
			for _, product := range bpo.Invention.Products {
				T2toT1[product.TypeID] = bpo.BlueprintTypeID
			}
		}
		if bpo.Manufacturing != nil && len(bpo.Manufacturing.Products) > 0 {
			BpoIDByTypeID[bpo.Manufacturing.Products[0].TypeID] = bpoID
		}
	}

}
