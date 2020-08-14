package util

import "gopkg.in/mgo.v2/bson"

func GenerateDataId() string {
	id := bson.NewObjectId().Hex()
	return id
}
