package dao

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
	"github.com/joomcode/errorx"
)

var (
	doOnce   sync.Once
	database *InMemDatabase
)

type InMemDatabase struct {
	db *memdb.MemDB
}

func newInMemDatabase() *InMemDatabase {
	doOnce.Do(
		func() {
			schema := &memdb.DBSchema{
				Tables: map[string]*memdb.TableSchema{
					"person": {
						Name: "person",
						Indexes: map[string]*memdb.IndexSchema{
							"id": {
								Name:    "id",
								Unique:  true,
								Indexer: &memdb.StringFieldIndex{Field: "Id"},
							},
							"age": {
								Name:    "age",
								Unique:  false,
								Indexer: &memdb.IntFieldIndex{Field: "Age"},
							},
							"email": {
								Name:    "email",
								Unique:  true,
								Indexer: &memdb.StringFieldIndex{Field: "Email"},
							},
						},
					},
				},
			}
			db, err := memdb.NewMemDB(schema)
			if err != nil {
				panic(err)
			}
			database = &InMemDatabase{
				db: db,
			}
		},
	)
	return database
}

func GetDatabaseInMemoryDatabase() *InMemDatabase {
	return newInMemDatabase()
}

func (id *InMemDatabase) Add(tableName string, obj interface{}) *errorx.Error {
	txn := id.db.Txn(true)

	if err := txn.Insert(tableName, obj); err != nil {
		return errorx.Decorate(err, fmt.Sprintf("An error occurred while inserting data on %s", tableName))
	}

	// Commit the transaction
	txn.Commit()

	// Create read-only transaction
	txn = id.db.Txn(false)
	defer txn.Abort()
	return nil
}

func (id *InMemDatabase) Get(tableName string) ([]interface{}, *errorx.Error) {
	result := make([]interface{}, 0)
	txn := id.db.Txn(false)
	it, err := txn.Get("person", "id")
	if err != nil {
		return nil, errorx.Decorate(err, fmt.Sprintf("Error while retrieving data from %s", tableName))
	}
	for obj := it.Next(); obj != nil; obj = it.Next() {
		result = append(result, obj)
	}

	// Commit the transaction
	txn.Commit()

	defer txn.Abort()
	return result, nil
}

func (im *InMemDatabase) GetById(tableName string, id uuid.UUID) (interface{}, *errorx.Error) {
	return nil, nil
}
