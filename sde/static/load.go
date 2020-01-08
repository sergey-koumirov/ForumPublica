package static

import (
	"ForumPublica/sde/models"
	"ForumPublica/sde/service"
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	//Types map with all types description
	Types models.ZipTypes

	//Groups map with all types description
	Groups models.ZipGroups

	//Blueprints map with all blueprints description
	Blueprints models.ZipBlueprints

	//T2toT1 get T1 bpo id by T2 bpo id
	T2toT1 map[int32]int32

	//BpoIDByTypeID get bpo id by product id
	BpoIDByTypeID map[int32]int32

	//SolarSystemsList array with K-space solar systems
	SolarSystemsList models.ZipSolarSystemsList

	//SolarSystems hash with K-space solar systems
	SolarSystems models.ZipSolarSystems

	//RegionsList array with K-space regions
	RegionsList models.ZipRegions

	//Regions hash with K-space regions
	Regions map[int64]models.ZipRegion
)

//LoadJSONs load data from zipped jsons
func LoadJSONs(fileName string, resetCache bool) {

	if resetCache {
		loadZip(fileName)
		saveYAML(Types, "./sde/cache/_types.yaml")
		saveYAML(Blueprints, "./sde/cache/_blueprints.yaml")
		saveYAML(T2toT1, "./sde/cache/_t2_to_t1.yaml")
		saveYAML(BpoIDByTypeID, "./sde/cache/_bpo_id_by_type_id.yaml")
		saveYAML(SolarSystemsList, "./sde/cache/_solar_systems_list.yaml")
		saveYAML(SolarSystems, "./sde/cache/_solar_systems.yaml")
		saveYAML(RegionsList, "./sde/cache/_regions_list.yaml")
		saveYAML(Regions, "./sde/cache/_regions.yaml")
		saveYAML(Groups, "./sde/cache/_groups.yaml")
	} else {
		loadYAML(&Types, "./sde/cache/_types.yaml")
		loadYAML(&Blueprints, "./sde/cache/_blueprints.yaml")
		loadYAML(&T2toT1, "./sde/cache/_t2_to_t1.yaml")
		loadYAML(&BpoIDByTypeID, "./sde/cache/_bpo_id_by_type_id.yaml")
		loadYAML(&SolarSystemsList, "./sde/cache/_solar_systems_list.yaml")
		loadYAML(&SolarSystems, "./sde/cache/_solar_systems.yaml")
		loadYAML(&RegionsList, "./sde/cache/_regions_list.yaml")
		loadYAML(&Regions, "./sde/cache/_regions.yaml")
		loadYAML(&Groups, "./sde/cache/_groups.yaml")
	}

}

func loadYAML(obj interface{}, filename string) {
	dir, errWd := os.Getwd()
	if errWd != nil {
		fmt.Println("loadYAML: ", errWd, filename)
		return
	}

	freshData, err := ioutil.ReadFile(path.Join(dir, filename))
	if err != nil {
		fmt.Println("loadYAML", err, filename)
	} else {
		yaml.Unmarshal([]byte(freshData), obj)
	}
}

func saveYAML(obj interface{}, filename string) {
	dir, errWd := os.Getwd()
	if errWd != nil {
		fmt.Println("loadYAML: ", errWd, filename)
		return
	}

	bytes, err := yaml.Marshal(obj)
	if err != nil {
		fmt.Println("saveYAML: ", err, filename)
		return
	}

	fmt.Println("save: ", path.Join(dir, filename))
	errWr := ioutil.WriteFile(path.Join(dir, filename), bytes, 0644)
	if err != nil {
		fmt.Println("saveYAML: ", errWr, filename)
	}

}

func loadZip(fileName string) {
	r, zipErr := zip.OpenReader(fileName)
	if zipErr != nil {
		fmt.Println("zip.OpenReader:", zipErr)
		return
	}
	defer r.Close()

	var names map[int64]string
	SolarSystemsList = make(models.ZipSolarSystemsList, 0)

	for _, f := range r.File {
		// fmt.Printf("%+v\n", f.FileHeader.Name)
		if f.FileHeader.Name == "sde/fsd/typeIDs.yaml" {
			Types = *service.ImportTypes(f)
		}

		if f.FileHeader.Name == "sde/fsd/groupIDs.yaml" {
			Groups = *service.ImportGroups(f)
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

	SolarSystems = make(models.ZipSolarSystems)
	for i, s := range SolarSystemsList {
		SolarSystemsList[i].Name = names[s.ID]
		r, _ := regionsByKey[s.RegionKey]
		SolarSystemsList[i].RegionID = r.ID
		SolarSystems[s.ID] = SolarSystemsList[i]
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
