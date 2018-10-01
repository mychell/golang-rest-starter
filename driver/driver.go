package driver

import (
	"github.com/gobuffalo/pop"
)

// DB ...A struct of holding various db connections
type DB struct {
	PQ *pop.Connection
}

// DBConn ...
var dbConn = &DB{}

// ConnectPQ ...
func ConnectPQ(env string) (*DB, error) {

	d, err := pop.Connect(env)
	if err != nil {
		panic(err)
	}
	dbConn.PQ = d
	return dbConn, err
}
