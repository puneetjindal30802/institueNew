package models

import "gopkg.in/mgo.v2/bson"

/*
 * Session  is used to hold the session data
 */
type Session struct {
	Token              string        `json:"token", bson:"token,omitempty"`
	UserId             bson.ObjectId `json:"user_id" bson:"user_id"`
	CreatedOn          int64         `json:"created_on,omitempty" bson:"created_on,omitempty"`
	LastActivityOn     int64         `json:"last_activity_on,omitempty" bson:"last_activity_on,omitempty"`
	DeviceTokenId      string        `json:"device_token_id,omitempty" bson:"device_token_id,omitempty"`
	DeviceType         string        `json:"device_type,omitempty" bson:"device_type"`
	OnetimeAccessToken string        `json:"onetime_access_token,omitempty" bson:"onetime_access_token,omitempty"`
}
