package datastore

// install dgraph client in $GOROOT:
// go get -u -v github.com/dgraph-io/dgo

// import (
// 	"github.com/jackc/pgx"
// 	"gitlab.com/norzion/temp0/command-handler/errors"
// )

// type psqlStore struct {
// 	connPool *pgx.ConnPool
// }

// // NewPSQLStore creates a postgres database based datastore using "URI" and "maxConn" to connect
// func NewPSQLStore(URI string, maxConn int) (Datastore, *errors.Error) {
// 	ps := &psqlStore{}
// 	con, err := pgx.ParseURI(URI)
// 	if err != nil {
// 		return nil, errors.NewError(err.Error())
// 	}

// 	p, err := pgx.NewConnPool(
// 		pgx.ConnPoolConfig{
// 			ConnConfig:     con,
// 			MaxConnections: maxConn,
// 		})

// 	if err != nil {
// 		return nil, errors.NewError(err.Error())
// 	}

// 	ps.connPool = p
// 	return ps, nil
// }

// func (ds *psqlStore) StoreSensorMeasurements(device *Temp0device) *errors.Error {
// 	return nil
// }
