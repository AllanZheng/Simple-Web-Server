package entity

import (

)

type Info struct {
	ID string `json:"id" bson:"id"` //
	Name string `json:"name" bson:"name"`//
	
}

type User struct {
	UserID string `json:"userid" bson:"userid"`
	Age int `json:"age" bson:"age"`
	Dep string `json:"dep" bson:"dep"`
	UserName string `json:"username" bson:"username"`
}

