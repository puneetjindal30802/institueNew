package models

import (
	"institute/config"

	"gopkg.in/mgo.v2/bson"
)

type UserSignup struct {
	ID           bson.ObjectId `json:"_id" bson:"_id"`
	FirstName    string        `json:"first_name" bson:"first_name"`
	LastName     string        `json:"last_name" bson:"last_name"`
	FullName     string        `json:"full_name,omitempty," bson:"full_name"`
	Email        string        `json:"email" bson:"email"`
	OTPCode      string        `json:"otp_code" bson:"otp_code"`
	JoinedStatus int           `json:"joined_status" bson:"joined_status"`
	Password     string        `json:"password" bson:"password"`
	TimeDate     int64         `json:"time_date" bson:"time_date"`
}

type AuthorizeUser struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

/*
 *	Function to save user data into database
 *
 *	Return err
 */
func SaveUserData(query interface{}) (err error) {
	err = DbInsert(config.UsersCollection, query)

	return err
}

/*
 *	Function to get the user from the database
 *
 *	Return user data
 */
func GetUserLoginInfo(email, password string) (userData UserSignup, err error) {
	mongoSession := config.ConnectDb(config.Database)
	defer mongoSession.Close()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	getCollection := sessionCopy.DB(config.Database).C(config.UsersCollection)
	err = getCollection.Find(bson.M{"email": email, "password": password, "joined_status": 1}).One(&userData)
	return userData, err
}

/*
 *	Function to get the user from the database
 *
 *	Return user data
 */
func GetUserByOTPCode(otp string) (userData UserSignup, err error) {
	mongoSession := config.ConnectDb(config.Database)
	defer mongoSession.Close()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	getCollection := sessionCopy.DB(config.Database).C(config.UsersCollection)
	err = getCollection.Find(bson.M{"otp_code": otp}).Select(bson.M{"otp_code": 1}).One(&userData)
	return userData, err
}

/*
 *	Function to change the joined status
 *
 *	Return err type error
 */
func UpdateUserJoinedStatus(otp string) (err error) {
	mongoSession := config.ConnectDb(config.Database)
	defer mongoSession.Close()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	getCollection := sessionCopy.DB(config.Database).C(config.UsersCollection)
	err = getCollection.Update(bson.M{"otp_code": otp}, bson.M{"$set": bson.M{"joined_status": 1}})
	return err
}

/*
 *	Function which return the registered users
 *
 *	Returns user and error
 */
func GetRegisteredUser(email string) (user UserSignup, err error) {
	mongoSession := config.ConnectDb(config.Database)
	defer mongoSession.Close()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	getCollection := sessionCopy.DB(config.Database).C(config.UsersCollection)
	err = getCollection.Find(bson.M{"email": email, "joined_status": 1}).One(&user)
	return user, err
}
