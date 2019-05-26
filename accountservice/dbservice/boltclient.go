package dbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/adityaladwa/go-microservices/accountservice/model"
	"github.com/boltdb/bolt"
)

type IBoltClient interface {
	OpenBoltDb()
	QueryAccount(accountId string) (model.Account, error)
	Seed()
}

type BoltClient struct {
	boltDb *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	boltDb, err := bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer boltDb.Close()
}

func (bc *BoltClient) initializeBucket() {
	bc.boltDb.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("AccountBucket"))
		if err != nil {
			return fmt.Errorf("Create bucket failed: %s", err)
		}
		return nil
	})
}

func (bc *BoltClient) Seed() {
	initializeBucket()
	seedAccounts()
}

func (bc *BoltClient) seedAccounts() {
	total := 100
	for i := 0; i < total; i++ {
		key := strconv.Itoa(10000 + i)
		acc := model.Account{
			Id:   key,
			Name: "Person_" + strconv.Itoa(i),
		}

		jsonBytes, _ := json.Marshal(acc)
		bc.boltDb.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("AccountBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Seeded %v fake accounts... \n", total)
}