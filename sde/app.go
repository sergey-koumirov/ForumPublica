package main

import (
	"ForumPublica/sde/service"
	"ForumPublica/server/config"
	"ForumPublica/server/db"
	"archive/zip"
	"flag"
	"fmt"
)

func main() {

	config.LoadVars()
	if config.Vars == nil {
		fmt.Println("Vars load problem")
		return
	}

	db.Connect()
	if db.DB == nil {
		return
	}

	fileName := flag.String("file", "NONE", "a string")
	flag.Parse()

	fmt.Println("Unzip...", *fileName)
	if *fileName == "NONE" {
		return
	}

	r, zipErr := zip.OpenReader(*fileName)
	if zipErr != nil {
		fmt.Println("zip.OpenReader:", zipErr)
		return
	}
	defer r.Close()

	for _, f := range r.File {
		// fmt.Printf("%+v\n", f.FileHeader.Name)
		service.ImportTypes(f)
		service.ImportBlueprints(f)
	}

}
