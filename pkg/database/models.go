package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Email    string             `bson:"email"`
}

type ProductItem struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	Name        string             `bson:"name"`
	Price       int32              `bson:"price"`
	Description string             `bson:"description"`
}
