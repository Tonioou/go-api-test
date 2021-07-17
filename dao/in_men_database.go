package dao

import (
	"fmt"

	"github.com/Tonioou/go-api-test/model"
	"github.com/hashicorp/go-memdb"
	"github.com/joomcode/errorx"
)

type InMemDatabase struct {
	db *memdb.MemDB
}

func NewInMemDatabase() *InMemDatabase {
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
	return &InMemDatabase{
		db: db,
	}
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
