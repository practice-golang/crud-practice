# NP

NP - `N`ull types and `P`ointer to comma separated string

Create comma separated names and values from `struct` or `map` for raw SQL query.

This is an experimental and for my own usage.

## Which types of
* `database/sql` Nullnnnn
* `github.com/guregu/null`
* Map - Convert kebab case keys to upper case with underscore
* Pointer

## Example

See `np_test.go`

```go
package main

import (
	"github.com/practice-golang/np"
)

// Human
type Human struct {
	Name         null.String `json:"name" db:"NAME"`
	Age          null.Int    `json:"age"  db:"AGE"`
	EmailAddress null.String `json:"email-address,omitempty" db:"EMAIL_ADDRESS"`
}

func main() {
	john := Human{
		null.NewString("John", true),
		null.NewInt(777, true),
		null.NewString("john@human.io", true),
	}

	np.TagName = "db"
	np.Separator = ","

	colString := np.MakeString(john)

	selectQuery := "SELECT " + colString.Names + " FROM table_name;"
	insertQuery := "INSERT INTO table_name (" + colString.Names + ") VALUES(" + colString.Values + ");"

	fmt.Println(selectQuery)
	fmt.Println(insertQuery)
}

// Result:
// SELECT NAME,AGE,EMAIL_ADDRESS FROM table_name;
// INSERT INTO table_name (NAME,AGE,EMAIL_ADDRESS) VALUES(John,777,john@human.io)
```
