package np

import (
	"database/sql"
	"encoding/json"
	"log"
	"strings"
	"testing"

	"gopkg.in/guregu/null.v4"
)

// Human - guregu/null
type Human struct {
	Name         null.String `json:"name" db:"NAME"`
	Age          null.Int    `json:"age" db:"AGE"`
	EmailAddress null.String `json:"email-address,omitempty" db:"EMAIL_ADDRESS"`
}

// Agent - sql.Null..
type Agent struct {
	Name         sql.NullString `json:"name" db:"NAME"`
	Age          sql.NullInt64  `json:"age" db:"AGE"`
	EmailAddress sql.NullString `json:"email-address,omitempty" db:"EMAIL_ADDRESS"`
}

// Ander - pointer
type Ander struct {
	Name         *string `json:"name" db:"NAME"`
	Age          *int    `json:"age" db:"AGE"`
	EmailAddress *string `json:"email-address,omitempty" db:"EMAIL_ADDRESS"`
}

func Test_createColString_null_struct(t *testing.T) {
	john := Human{
		null.NewString("John", true),
		null.NewInt(777, true),
		null.NewString("john@human.io", true),
	}

	colString := CreateString(john, "sqlite", "")

	names := strings.Split(colString.Names, ",")
	values := strings.Split(colString.Values, ",")

	nameSample := []string{"NAME", "AGE", "EMAIL_ADDRESS"}
	valueSample := []string{"John", "777", "john@human.io"}

	for i, name := range names {
		isExist := false
		for j, n := range nameSample {
			if n == name {
				if values[i] != valueSample[j] {
					t.Fatal("\nExpect:", strings.Join(valueSample, ","), "\nResult:", colString.Values)
				}

				isExist = true
				break
			}
		}
		if !isExist {
			t.Fatal("\nExpect:", strings.Join(nameSample, ","), "\nResult:", colString.Names)
		}
	}

	log.Println()
	log.Println("INSERT INTO table_name (" + colString.Names + ") VALUES(" + colString.Values + ")")
}

func Test_createColString_json_struct(t *testing.T) {
	jane := Human{}
	janeJSON := []byte("{ \"name\": \"Jane\", \"age\": 999, \"email-address\": \"jane@human.io\" }")
	json.Unmarshal(janeJSON, &jane)

	colString := CreateString(jane, "sqlite", "")

	names := strings.Split(colString.Names, ",")
	values := strings.Split(colString.Values, ",")

	nameSample := []string{"NAME", "AGE", "EMAIL_ADDRESS"}
	valueSample := []string{"Jane", "999", "jane@human.io"}

	for i, name := range names {
		isExist := false
		for j, n := range nameSample {
			if n == name {
				if values[i] != valueSample[j] {
					t.Fatal("\nExpect:", strings.Join(valueSample, ","), "\nResult:", colString.Values)
				}
				isExist = true

				break
			}
		}
		if !isExist {
			t.Fatal("\nExpect:", strings.Join(nameSample, ","), "\nResult:", colString.Names)
		}
	}
}

func Test_createColString_json_map(t *testing.T) {
	james := map[string]interface{}{}
	jamesJSON := []byte("{ \"name\": \"James\", \"age\": 888, \"email-address\": \"james@human.io\" }")
	json.Unmarshal(jamesJSON, &james)

	colString := CreateString(james, "sqlite", "")

	names := strings.Split(colString.Names, ",")
	values := strings.Split(colString.Values, ",")

	nameSample := []string{"NAME", "AGE", "EMAIL_ADDRESS"}
	valueSample := []string{"James", "888", "james@human.io"}

	for i, name := range names {
		isExist := false
		for j, n := range nameSample {
			if n == name {
				if values[i] != valueSample[j] {
					t.Fatal("\nExpect:", strings.Join(valueSample, ","), "\nResult:", colString.Values)
				}

				isExist = true
				break
			}
		}
		if !isExist {
			t.Fatal("\nExpect:", strings.Join(nameSample, ","), "\nResult:", colString.Names)
		}
	}
}

func Test_createColString_sql_null(t *testing.T) {
	smith := Agent{
		Name:         sql.NullString{String: "Smith", Valid: true},
		Age:          sql.NullInt64{Int64: int64(222), Valid: true},
		EmailAddress: sql.NullString{String: "smith@machine.io", Valid: true},
	}

	colString := CreateString(smith, "sqlite", "")

	names := strings.Split(colString.Names, ",")
	values := strings.Split(colString.Values, ",")

	nameSample := []string{"NAME", "AGE", "EMAIL_ADDRESS"}
	valueSample := []string{"Smith", "222", "smith@machine.io"}

	for i, name := range names {
		isExist := false
		for j, n := range nameSample {
			if n == name {
				if values[i] != valueSample[j] {
					t.Fatal("\nExpect:", strings.Join(valueSample, ","), "\nResult:", colString.Values)
				}

				isExist = true
				break
			}
		}
		if !isExist {
			t.Fatal("\nExpect:", strings.Join(nameSample, ","), "\nResult:", colString.Names)
		}
	}
}

func Test_createColString_pointer(t *testing.T) {
	thomas := Ander{}
	thomasJSON := []byte("{ \"name\": \"Thomas\", \"age\": 444, \"email-address\": \"thomas@son.io\" }")
	json.Unmarshal(thomasJSON, &thomas)

	colString := CreateString(thomas, "sqlite", "")

	names := strings.Split(colString.Names, ",")
	values := strings.Split(colString.Values, ",")

	nameSample := []string{"NAME", "AGE", "EMAIL_ADDRESS"}
	valueSample := []string{"Thomas", "444", "thomas@son.io"}

	for i, name := range names {
		isExist := false
		for j, n := range nameSample {
			if n == name {
				if values[i] != valueSample[j] {
					t.Fatal("\nExpect:", strings.Join(valueSample, ","), "\nResult:", colString.Values)
				}

				isExist = true
				break
			}
		}
		if !isExist {
			t.Fatal("\nExpect:", strings.Join(nameSample, ","), "\nResult:", colString.Names)
		}
	}
}
