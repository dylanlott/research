package boltdb

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"github.com/google/uuid"
	"log"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}

// CreateUser calls bolt database instance to create user
func (bdb *Client) CreateUser(user User) error {
	return bdb.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("users"))

		usernameKey := []byte(user.Username)
		userBytes, err := json.Marshal(user)
		if err != nil {
			log.Println(err)
		}

		return b.Put(usernameKey, userBytes)
	})
}

func (bdb *Client) GetUser(key []byte) (User, error) {
	var userInfo User
	err := bdb.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("users"))
		v := b.Get(key)
		if v == nil {
			log.Println("user not found")
			return nil
		} else {
			err1 := json.Unmarshal(v, &userInfo)
			return err1
		}
	})

	return userInfo, err
}

func (bdb *Client) UpdateUser(user User) error {
	return bdb.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("users"))

		usernameKey := []byte(user.Username)
		userBytes, err := json.Marshal(user)
		if err != nil {
			log.Println(err)
		}

		return b.Put(usernameKey, userBytes)
	})
}

func (bdb *Client) DeleteUser(key []byte) {
}
