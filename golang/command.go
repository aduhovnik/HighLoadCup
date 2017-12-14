package main

import (
	"fmt"
	"golang/models"
	"golang/db"
	"io/ioutil"
	"encoding/json"
	"math"
	"time"
)


func run_migrations(){
	//Migrate the schema
	_db := db.Database()
	_db.Debug().AutoMigrate(&models.User{})
	_db.Debug().AutoMigrate(&models.Location{})
	_db.Debug().AutoMigrate(&models.Visit{})
	fmt.Print("Migrations were performed")
}

func read_file(filename string) []byte{
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		return nil
	}
	fmt.Printf("%s\n", string(file))
	return file
}

func load_data(){
	_db := db.Database()

	path := "./data/data/TRAIN/"

	for fi := 1; ; fi += 1 {
		file := read_file(fmt.Sprintf("%vusers_%v.json", path, fi))

		if file == nil {
			break
		}

		var ulist models.Users_list
		json.Unmarshal(file, &ulist)
		fmt.Printf("Results: %v\n", ulist)

		for i := 0; i < len(ulist.Users); i += 100 {
			for j := i; j < int(math.Min(float64(i+100), float64(len(ulist.Users)))); j++ {
				go _db.Create(ulist.Users[j])
			}

			time.Sleep(1 * time.Second)
		}

	}

	for fi := 1; ; fi += 1 {
		file := read_file(fmt.Sprintf("%vlocations_%v.json", path, fi))

		if file == nil {
			break
		}

		var llist models.Locations_list
		json.Unmarshal(file, &llist)
		fmt.Printf("Results: %v\n", llist)

		for i := 0; i < len(llist.Locations); i += 100 {
			for j := i; j < int(math.Min(float64(i+100), float64(len(llist.Locations)))); j++ {
				go _db.Create(llist.Locations[j])
			}

			time.Sleep(1 * time.Second)
		}
	}

	for fi := 1; ; fi += 1 {
		file := read_file(fmt.Sprintf("%vvisits_1.json", path, fi))

		if file == nil {
			break
		}

		var vlist models.Visits_list
		json.Unmarshal(file, &vlist)
		fmt.Printf("Results: %v\n", vlist)

		for i := 0; i < len(vlist.Visits); i += 100 {
			for j := i; j < int(math.Min(float64(i+100), float64(len(vlist.Visits)))); j++ {
				go _db.Create(vlist.Visits[j])
			}

			time.Sleep(1 * time.Second)
		}
	}
}


//func main() {
	//run_migrations()
//	load_data()
//}
