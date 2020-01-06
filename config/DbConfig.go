package config

import (
	"crypto/tls"
	"fmt"
	"net"

	"gopkg.in/mgo.v2"
)

/*
 * Function to connect to the database
 *
 * Used by all the models
 *
 * Params mongoSession type *mgo.Session
 *
 * Returns mongoSession type *mgo.Session
 */
func ConnectDb(Db string) (mongoSession *mgo.Session) {
	// mongoDBDialInfo := &mgo.DialInfo{
	// 	Addrs: []string{"mongodb+srv://puneet:1234@cluster0-a8zxr.mongodb.net/test?retryWrites=true&w=majority"},
	// 	// Addrs:    []string{"172.20.1.100:27017"},
	// 	Timeout:  60 * time.Second,
	// 	Database: Db,
	// }

	// mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	// if err != nil {
	// 	log.Fatalf("CreateSession: %s\n", err)
	// }
	// mongoSession.SetMode(mgo.Monotonic, true)

	// return mongoSession

	info, err := mgo.ParseURL("mongodb://puneet:1234@cluster0-shard-00-00-a8zxr.mongodb.net:27017,cluster0-shard-00-01-a8zxr.mongodb.net:27017,cluster0-shard-00-02-a8zxr.mongodb.net:27017/institute?authSource=admin")
	if err != nil {
		panic(fmt.Errorf("%s", err))
	}
	tlsConfig := &tls.Config{}
	info.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	mongoSession, err = mgo.DialWithInfo(info)
	if err != nil {
		panic(fmt.Errorf("%s", err))
	}

	mongoSession.SetMode(mgo.Monotonic, true)

	return mongoSession
}
