package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	//"reflect"
	//"strings"
	//"strconv"
	"fmt"
)



// grid structure
type grid struct {
	ID       string
	Power    float64
}

// cfg is the struct to unmarshal all of twins.json into
type cfg struct {
	ID             string
	UpdateRate     uint
	PublishRate    uint
	TimeMultiplier uint
	ManualTick     bool
	TestNum        float64
	Grids          []grid
}

// readConfig() looks for args to the program, follows the first arg
// as a path, and looks for twins.json in it
func readConfig(config *cfg) {
	// Looking for twins.json
	cpath :=  "twins.json"
	if len(os.Args) > 2 {
		cpath = os.Args[2]
	}		
	if len(os.Args) < 2 {
		log.Print("Config path argument not found. Usage 'twins /path/to/config'. Trying current working directory")
		cpath = "twins.json"
	} else {
		info, err := os.Stat(os.Args[1])
		if os.IsNotExist(err) {
			log.Fatal("TWINS configuration file not found at: ", os.Args[1])
		} else if info.IsDir() {
			cpath = path.Join(os.Args[1], cpath)
		} else {
			cpath = os.Args[1]
		}
	}
	
	configjson, err := ioutil.ReadFile(cpath)
	if err != nil {
		log.Fatalf("Couldn't read the file %s: %s", cpath, err)
	}
	err = json.Unmarshal(configjson, config)
	if err != nil {
		log.Fatal("Failed to Unmarshal config file")
	}
}

func main () {
	mycfg := new(cfg)
	readConfig(mycfg)
	fmt.Println(" Got config ID [" + mycfg.ID + "]" )
	fmt.Print(" Got mycfg.TestNum ["  )
	fmt.Print( mycfg.TestNum  )
	fmt.Print("]\n" )
	fmt.Println(" Got Grid ID [" + mycfg.Grids[0].ID + "]" )
}