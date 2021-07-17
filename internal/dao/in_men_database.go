package dao

import (
	"fmt"
	"sync"

	"github.com/Tonioou/go-api-test/internal/model"
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
								Indexer: &memdb.StringFieldIndex{Field: "id"},
							},
							"age": {
								Name:    "age",
								Unique:  false,
								Indexer: &memdb.IntFieldIndex{Field: "age"},
							},
							"email": {
								Name:    "email",
								Unique:  true,
								Indexer: &memdb.StringFieldIndex{Field: "email"},
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
func (id *InMemDatabase) Add(tableName string, obj interface{}) (uuid.UUID, *errorx.Error) {
	txn := id.db.Txn(true)

	if err := txn.Insert(tableName, obj); err != nil {
		return uuid.Nil, errorx.Decorate(err, fmt.Sprintf("An error occurred while inserting data on %s", tableName))
	}

	// Commit the transaction
	txn.Commit()

	// Create read-only transaction
	txn = id.db.Txn(false)
	defer txn.Abort()
	return uuid.Nil, nil
}

func (id *InMemDatabase) Get() (interface{}, *errorx.Error) {
	people := make([]*model.Person, 0)
	txn := id.db.Txn(false)
	it, err := txn.Get("person", "id")
	if err != nil {
		panic(err)
	}
	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*model.Person)
		people = append(people, p)
	}

	// Commit the transaction
	txn.Commit()

	// Create read-only transaction
	txn = id.db.Txn(false)
	defer txn.Abort()
	return interface{}(people), nil
}

func (im *InMemDatabase) GetById(id uuid.UUID) (interface{}, *errorx.Error) {
	return nil, nil
}
