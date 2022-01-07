package np

import (
	"errors"
	"fmt"
	"strings"
)

// CreateHolders
// Create preparing holder string
// like ? or $1, $2.. from column names
func CreateHolders(dbtype, colNames string) (string, error) {
	var err error
	var binding string

	if strings.TrimSpace(colNames) == "" {
		return "", errors.New("colNames is empty")
	}
	count := strings.Count(colNames, ",") + 1

	switch strings.ToLower(dbtype) {
	case "postgres":
		for i := 0; i < count; i++ {
			binding += "$" + fmt.Sprint(i+1)
			if i < count-1 {
				binding += ","
			}
		}
	default:
		for i := 0; i < count; i++ {
			binding += "?"
			if i < count-1 {
				binding += ","
			}
		}
	}

	return binding, err
}

// CreateUpdateHolders
// Create preparing holder string for update
// like FIELDA=?, FIELDB=? or FIELDA=$1, FIELDB=$2.. from column names
func CreateUpdateHolders(dbtype string, columnNames interface{}, offset int) (string, int, error) {
	var err error
	var binding string

	colNames := []string{}
	count := 0

	switch cnames := columnNames.(type) {
	case string:
		// for _, n := range strings.Split(colNameSTR, ",") {
		// 	n = strings.TrimSpace(n)
		// 	n = strings.Trim(n, "`")
		// 	n = strings.Trim(n, `"`)
		// 	n = strings.Trim(n, "'")
		// 	colNames = append(colNames, strings.TrimSpace(n))
		// 	colNames = append(colNames, n)
		// }
		colNameSTR := cnames
		colNames = strings.Split(colNameSTR, ",")
		count = len(colNames)

	case []string:
		colNames = cnames
		count = len(colNames)

	case map[string]string:
		for k := range cnames {
			colNames = append(colNames, k)
		}
	}

	for i := 0; i < count; i++ {
		holder := "?"
		if dbtype == "postgres" {
			holder = "$" + fmt.Sprint(i+1+offset)
		}
		binding += colNames[i] + "=" + holder
		if i < count-1 {
			binding += ","
		}
	}

	return binding, count, err
}
