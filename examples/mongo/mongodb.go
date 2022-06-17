package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"suzaku/pkg/common/config"
	"sync"
	"time"
)

var MongoDB *mongoDB

type mongoDB struct {
	sync.RWMutex
	dbMap map[string]*mongo.Database
}

func init() {
	MongoDB = &mongoDB{dbMap: map[string]*mongo.Database{}}
	MongoDB.open(config.Config.Mongo.Address[0], config.Config.Mongo.Db)
}

func (mg *mongoDB) open(address string, dbName string) (err error) {
	var (
		uri           string
		clientOptions *options.ClientOptions
		client        *mongo.Client
	)

	uri = fmt.Sprintf("mongodb://%s/?maxPoolSize=%d", address, config.Config.Mongo.MaxPoolSize)
	// Set client options
	clientOptions = options.Client().ApplyURI(uri).SetAuth(
		options.Credential{
			AuthMechanism: "SCRAM-SHA-256",
			Username:      config.Config.Mongo.Username,
			Password:      config.Config.Mongo.Password,
		}).SetConnectTimeout(time.Duration(config.Config.Mongo.Timeout) * time.Millisecond)

	// Connect to MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Connection to MongoDB Error:", err)
		return
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Check the connection error:", err)
		return
	}
	fmt.Println("Connected to MongoDB!")

	mg.dbMap[dbKey(address, dbName)] = client.Database(dbName)
	return
}

func (mg *mongoDB) getDB(address string, dbName string) (db *mongo.Database, err error) {
	var (
		key string
		ok  bool
	)
	key = dbKey(address, dbName)
	if _, ok = mg.dbMap[key]; ok == false {
		if err = mg.open(address, dbName); err != nil {
			return
		}
	}
	mg.Lock()
	defer mg.Unlock()
	db = mg.dbMap[key]
	return
}

func dbKey(address string, dbName string) string {
	return address + "_" + dbName
}

func GetDB(address string, dbName string) (db *mongo.Database, err error) {
	db, err = MongoDB.getDB(address, dbName)
	return
}

func MgDB() (db *mongo.Database, err error) {
	return GetDB(config.Config.Mongo.Address[0], config.Config.Mongo.Db)
}
