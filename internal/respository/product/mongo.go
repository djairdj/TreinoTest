package product

//Implementação do Repository

import (
	"TreinoTest/entity"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongoRepository(db *mongo.Database) Repository {
	return &mongoRepository{
		db: db,
	}
}

func (r mongoRepository) Create(ctx context.Context, name string) (*entity.Product, error) {
	coll := r.db.Collection("product")
	data := bson.D{{"name", name}, {"votes", 0}}
	res, err := coll.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	id, ok := res.InsertedID.(primitive.ObjectID) //feitura de cast
	if !ok {
		return nil, fmt.Errorf("nao foi possivel converter o id")
	}
	p := entity.Product{
		ID:    id.Hex(),
		Name:  name,
		Votes: 0,
	}

	return &p, nil
}
