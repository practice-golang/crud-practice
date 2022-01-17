package np

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// ColumnStrings - strings for names and values
type ColumnStrings struct {
	Names  string
	Values string
}

var (
	TagName       = "db"
	TagNameNPSKIP = "npskip"
	Separator     = ","
)

func createString(o interface{}, dbtype, skipValue, separatorNames, separatorValues string) ColumnStrings {
	names := ""
	values := ""

	ot := reflect.TypeOf(o)
	ov := reflect.ValueOf(o)

	switch ot.Kind() {
	case reflect.Struct:
		for i := 0; i < ov.NumField(); i++ {
			skipTag := ov.Type().Field(i).Tag.Get(TagNameNPSKIP)
			if skipTag != "" && skipValue != "" && strings.Contains(skipTag, skipValue) {
				continue
			}

			switch true {
			case ov.Field(i).Type().PkgPath() == "database/sql":
				// sql.NullString, swl.NullInt..
				valid := ov.Field(i).Field(1).Bool()
				if valid {
					names += ov.Type().Field(i).Tag.Get(TagName)
					values += fmt.Sprint(ov.Field(i).Field(0))
				}

			case ov.Field(i).Kind() == reflect.Struct:
				// Maybe null.String, null.Int..
				valueStruct := createString(ov.Field(i).Interface(), dbtype, skipValue, separatorNames, separatorValues)
				if valueStruct.Names != "" {
					names += valueStruct.Names
				} else {
					names += ov.Type().Field(i).Tag.Get(TagName)
				}
				if valueStruct.Values != "" {
					values += valueStruct.Values
				}

			case ov.Field(i).Kind() == reflect.Ptr:
				// Maybe pointer
				if !ov.Field(i).IsNil() {
					names += ov.Type().Field(i).Tag.Get(TagName)
					values += fmt.Sprint(ov.Field(i).Elem())
				}

			case ov.Field(i).Kind() == reflect.Interface:
				// Maybe interface
				if ov.Type().Field(i).Tag.Get(TagName) != "" {
					names += ov.Type().Field(i).Tag.Get(TagName)
					values += fmt.Sprint(ov.Field(i).Elem())
				}

			default:
				// Maybe string, int.. How to do?
				unknownValue := fmt.Sprint(ov.Field(i).Field(0))
				log.Println("Maybe string, int..", unknownValue)
			}

			if i < ov.NumField()-1 {
				names += separatorNames
				values += separatorValues
			}
		}

	case reflect.Map:
		for i, k := range ov.MapKeys() {
			key := strings.ReplaceAll(k.String(), "-", "_")
			key = strings.ToUpper(key)
			names += key
			values += fmt.Sprint(ov.MapIndex(k))

			if i < ov.Len()-1 {
				names += separatorNames
				values += separatorValues
			}
		}

	}

	result := ColumnStrings{
		Names:  names,
		Values: values,
	}

	return result
}

// CreateString - create string from struct, map
func CreateString(o interface{}, dbtype, skipValue string) ColumnStrings {
	quote := ""
	separatorNames := Separator
	separatorValues := Separator
	switch dbtype {
	case "mysql":
		quote = "`"
		separatorNames = "`,`"
		separatorValues = `','`
	case "postgres":
		quote = `"`
		separatorNames = `","`
		separatorValues = `','`
	case "sqlserver":
		quote = `"`
		separatorNames = `","`
		separatorValues = `','`
	}

	result := createString(o, dbtype, skipValue, separatorNames, separatorValues)
	result.Names = quote + result.Names + quote
	result.Values = quote + result.Values + quote

	return result
}

// CreateMapSlice
// create map which contains slice interface of names and values
// from struct or map
func CreateMapSlice(o interface{}, skipValue string) map[string][]interface{} {
	names := []interface{}{}
	values := []interface{}{}

	ot := reflect.TypeOf(o)
	ov := reflect.ValueOf(o)

	switch ot.Kind() {
	case reflect.Struct:
		for i := 0; i < ov.NumField(); i++ {
			skipTag := ov.Type().Field(i).Tag.Get(TagNameNPSKIP)
			if skipTag != "" && strings.Contains(skipTag, skipValue) {
				continue
			}

			switch true {
			case ov.Field(i).Type().PkgPath() == "database/sql":
				// sql.NullString, swl.NullInt..
				valid := ov.Field(i).Field(1).Bool()
				if valid {
					names = append(names, ov.Type().Field(i).Tag.Get(TagName))
					values = append(values, fmt.Sprint(ov.Field(i).Field(0)))
				}

			case ov.Field(i).Kind() == reflect.Struct:
				// Maybe null.String, null.Int..
				valueStruct := CreateMapSlice(ov.Field(i).Interface(), skipValue)
				if len(valueStruct["names"]) > 0 && valueStruct["names"][0] != "" {
					names = append(names, valueStruct["names"]...)
				} else {
					names = append(names, ov.Type().Field(i).Tag.Get(TagName))
				}
				if len(valueStruct["values"]) > 0 {
					values = append(values, valueStruct["values"]...)
				}

			case ov.Field(i).Kind() == reflect.Ptr, ov.Field(i).Kind() == reflect.Interface:
				// Maybe pointer, interface
				if !ov.Field(i).IsNil() {
					names = append(names, ov.Type().Field(i).Tag.Get(TagName))
					values = append(values, fmt.Sprint(ov.Field(i).Elem()))
				}

			default:
				// Maybe string, int.. How to do?
				unknownValue := fmt.Sprint(ov.Field(i).Field(0))
				log.Println("Maybe string, int..", unknownValue)
			}
		}

	case reflect.Map:
		for _, k := range ov.MapKeys() {
			key := strings.ReplaceAll(k.String(), "-", "_")
			key = strings.ToUpper(key)
			names = append(names, key)
			values = append(values, fmt.Sprint(ov.MapIndex(k)))
		}

	}

	result := map[string][]interface{}{
		"names":  names,
		"values": values,
	}

	return result
}

// CreateMap - create map from struct, map
func CreateMap(o interface{}, skipValue string) map[string]string {
	pairs := map[string]string{}
	name := ""
	value := ""

	ot := reflect.TypeOf(o)
	ov := reflect.ValueOf(o)

	switch ot.Kind() {
	case reflect.Struct:
		for i := 0; i < ov.NumField(); i++ {
			skipTag := ov.Type().Field(i).Tag.Get(TagNameNPSKIP)
			if skipTag != "" && strings.Contains(skipTag, skipValue) {
				continue
			}

			switch true {
			case ov.Field(i).Type().PkgPath() == "database/sql":
				// sql.NullString, swl.NullInt..
				valid := ov.Field(i).Field(1).Bool()
				if valid {
					name = ov.Type().Field(i).Tag.Get(TagName)
					value = fmt.Sprint(ov.Field(i).Field(0))

					pairs[name] = value
				}

			case ov.Field(i).Kind() == reflect.Struct:
				// Maybe null.String, null.Int..
				valueStruct := CreateMap(ov.Field(i).Interface(), skipValue)
				if len(valueStruct) > 0 {
					for _, v := range valueStruct {
						name := ov.Type().Field(i).Tag.Get(TagName)
						pairs[name] = v
					}
				}

			case ov.Field(i).Kind() == reflect.Ptr, ov.Field(i).Kind() == reflect.Interface:
				// Maybe pointer, interface
				if !ov.Field(i).IsNil() {
					name = ov.Type().Field(i).Tag.Get(TagName)
					value = fmt.Sprint(ov.Field(i).Elem())

					pairs[name] = value
				}

			default:
				// Maybe string, int.. How to do?
				unknownValue := fmt.Sprint(ov.Field(i).Field(0))
				log.Println("Maybe string, int..", unknownValue)
			}
		}

	case reflect.Map:
		for _, k := range ov.MapKeys() {
			key := strings.ReplaceAll(k.String(), "-", "_")
			key = strings.ToUpper(key)

			pairs[key] = fmt.Sprint(ov.MapIndex(k))
		}

	}

	return pairs
}
