package models

import (
	"instituteNew/config"
	"time"
)

func DbInsert(collectionName string, query interface{}) error {
	mongoSession := config.ConnectDb(config.Database)
	defer mongoSession.Close()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	getCollection := sessionCopy.DB(config.Database).C(collectionName)
	err := getCollection.Insert(query)
	return err

}

/*
 * Function to get current timestamp in UTC
 *
 * Returns type int64 (It returns timestamp for date + time)
 */
func GetCurrentDateTimestamp() int64 {
	currDateTime := time.Unix(time.Now().Unix(), 0)
	currDateTimeStamp := currDateTime.Unix()
	return currDateTimeStamp
}
