package utils

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

func MongoRepository(url ,username, password string)  (*mgo.Session, error){
	fmt.Print(url)
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{url},
		Timeout:  60 * time.Second,
		Username: username,
		Password: password,
	}
	db, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}