package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userDB struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Username string             `bson:"username" json:"username"`
}

type UserStorage struct {
	db *mongo.Database
}

func NewUserStorage(db *mongo.Database) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (s *UserStorage) createUser(email, password, username string, ctx context.Context) (string, error) {
	collection := s.db.Collection("users")
	result, err := collection.InsertOne(ctx, bson.M{"email": email, "password": password, "username": username})
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (s *UserStorage) getAllUsers(ctx context.Context) ([]userDB, error) {
	collection := s.db.Collection("users")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	users := make([]userDB, 0)
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
