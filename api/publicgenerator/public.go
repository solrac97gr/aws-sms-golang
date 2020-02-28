package public

import (
	"os"

	. "github.com/solrac97gr/sendsms/api/database"
)

var err error

func GetPublic() []string {
	var publics []string
	result, err := DBCon.Query(os.Getenv("QUERY_PUBLIC"))
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var public string
		err := result.Scan(&public)
		if err != nil {
			panic(err.Error())
		}
		publics = append(publics, public)
	}
	return publics
}
