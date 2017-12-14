package commands

import (
	"fmt"
	"golang/models"
	"golang/db"
	"io/ioutil"
	"encoding/json"
	"bytes"
)


func run_migrations(){
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

const chunk int = 10000;

func Load_data(){

	path := "/root/"

	_db := db.Database()

	for fi := 1; ; fi += 1 {
		file := read_file(fmt.Sprintf("%vusers_%v.json", path, fi))

		if file == nil {
			break
		}

		var ulist models.Users_list
		json.Unmarshal(file, &ulist)

		go func(file []byte) {
			var buffer bytes.Buffer
			sql := "INSERT INTO users (id, email, first_name, gender, last_name, birth_date) VALUES "
			buffer.WriteString(sql)
			end := len(ulist.Users)
			for j := 0; j < end; j++ {
				u := ulist.Users[j]
				add_sql := fmt.Sprintf("(%v, '%v', '%v', '%v', '%v', %v)",
					u.Id, u.Email, u.First_name, u.Gender, u.Last_name, u.Birth_date)

				if j != end-1 {
					add_sql += ", "
				}
				buffer.WriteString(add_sql)
			}
			sql = buffer.String()
			_db.Exec(sql)
		}(file)


	}

	for fi := 1; ; fi += 1 {
		file := read_file(fmt.Sprintf("%vlocations_%v.json", path, fi))

		if file == nil {
			break
		}

		go func(file []byte) {
			var llist models.Locations_list
			json.Unmarshal(file, &llist)

			sql := "INSERT INTO locations (id, place, country, city, distance) VALUES "
			end := len(llist.Locations)
			for j := 0; j < end; j++ {
				l := llist.Locations[j]
				add_sql := fmt.Sprintf("(%v, '%v', '%v', '%v', %v)", l.Id, l.Place, l.Country, l.City, l.Distance)
				sql += add_sql
				if j != end-1 {
					sql += ", "
				}
			}
			_db := db.Database()
			_db.Exec(sql)
		}(file)
	}

	for fi := 1; ; fi += 1 {
		file := read_file(fmt.Sprintf("%vvisits_%v.json", path, fi))

		if file == nil {
			break
		}
		go func(file []byte) {
			var vlist models.Visits_list
			json.Unmarshal(file, &vlist)

			sql := "INSERT INTO visits (id, location, user, visited_at, mark) VALUES "
			var buffer bytes.Buffer
			buffer.WriteString(sql)
			end := len(vlist.Visits)
			for j := 0; j < end; j++ {
				v := vlist.Visits[j]
				add_sql := fmt.Sprintf("(%v, %v, %v, %v, %v)", v.Id, v.Location, v.User, v.Visited_at, v.Mark)
				if j != end-1 {
					add_sql += ", "
				}
				buffer.WriteString(add_sql)
			}
			_db := db.Database()
			sql = buffer.String()
			_db.Exec(sql)
		}(file)
	}
}

/*
func main() {
	//run_migrations()
	load_data()
}
*/
