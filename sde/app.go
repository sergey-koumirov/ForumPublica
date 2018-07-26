package main

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/config"
	"ForumPublica/server/db"
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

	static.LoadJSONs(*fileName)

	fmt.Printf("%+v\n", static.Types[17325])
	fmt.Printf("%+v\n", static.Blueprints[17325])

}
