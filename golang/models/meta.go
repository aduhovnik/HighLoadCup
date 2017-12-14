package models

type metadata struct {
	fields []string
	json_fields []string
	data_type []string
	table_name string
}

var fields_meta map[string]metadata = map[string]metadata{
	"Users": {
		[]string{"Id", "Email", "First_name", "Last_name", "Birth_date"},
		[]string{"id", "email", "first_name", "last_name", "birth_date"},
		[]string{"int", "string", "string", "string", "int"},
		"Users"},
}
