package actions

import (
	"errors"
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var (
	DB_HOST   = ""
	DB_USER   = ""
	DB_PASSWD = ""
	DB_NAME   = ""
)

func GetCollection(collection string) (*mgo.Collection, error) {
	session, err := CreateDatabaseSession()
	if err != nil {
		return nil, err
	}

	c := session.DB("tourtique").C(collection)

	return c, nil
}

func CreateDatabaseSession() (*mgo.Session, error) {
	if ok := validateDBSettings(); !ok {
		return nil, errors.New("bad environment")
	}

	session, err := mgo.Dial(DB_HOST)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Eventual, true)

	creds := &mgo.Credential{
		Username: DB_USER,
		Password: DB_PASSWD,
		Source:   DB_NAME,
	}

	err = session.Login(creds)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func validateDBSettings() bool {
	// Validate DB_HOST
	if os.Getenv("DB_HOST") == "" {
		if DB_HOST == "" {
			log.Fatal("CAN'T detect a DB_HOST!")
			return false
		}
	} else {
		DB_HOST = os.Getenv("DB_HOST")
	}

	// Validate DB_USER
	if os.Getenv("DB_USER") == "" {
		if DB_USER == "" {
			log.Fatal("CAN'T detect a DB_USER!")
			return false
		}
	} else {
		DB_USER = os.Getenv("DB_USER")
	}

	// Validate DB_PASSWD
	if os.Getenv("DB_PASSWD") == "" {
		if DB_PASSWD == "" {
			log.Fatal("CAN'T detect a DB_PASSWD!")
			return false
		}
	} else {
		DB_PASSWD = os.Getenv("DB_PASSWD")
	}

	// Validate DB_NAME
	if os.Getenv("DB_NAME") == "" {
		if DB_NAME == "" {
			log.Fatal("CAN'T detect a DB_NAME!")
			return false
		}
	} else {
		DB_NAME = os.Getenv("DB_NAME")
	}

	// All required config strings found, proceed
	return true
}
