package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/smallcase/go-be-template/pkg/binotto"
	"github.com/smallcase/go-be-template/pkg/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoBinottoStore struct {
	col *mongo.Collection
}

const binottoCollectionName = "testBinotto"

type mongoBinotto struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Venue     string             `bson:"venue,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty"`
}

func New(mongoDB *mongo.Database) store.BinottoStore {
	col := mongoDB.Collection(binottoCollectionName)
	return &mongoBinottoStore{col: col}
}

func (b *mongoBinottoStore) GetById(ctx context.Context, id string) (*binotto.Binotto, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Could not convert '%s' to ObjectID", id))
	}
	var doc binotto.Binotto
	if err := b.col.FindOne(ctx, bson.M{"_id": oid}).Decode(&doc); err != nil {
		return nil, errors.Wrap(err, "Did not find a binotto doc in the DB")
	}
	return &doc, nil
}

func (b *mongoBinottoStore) Create(ctx context.Context, venue string) (*binotto.Binotto, error) {
	now := time.Now()
	doc := mongoBinotto{
		Venue:     venue,
		CreatedAt: now,
		UpdatedAt: now,
	}
	res, err := b.col.InsertOne(ctx, doc)
	if err != nil {
		return nil, errors.Wrap(err, "Could not create a new binotto doc in the DB")
	}
	return &binotto.Binotto{
		Id:    res.InsertedID.(string),
		Venue: venue,
	}, nil
}

func (b *mongoBinottoStore) UpdateVenue(ctx context.Context, id, venue string) error {
	update := bson.M{
		"$set": bson.M{
			"venue":     venue,
			"updatedAt": time.Now(),
		},
	}
	if _, err := b.col.UpdateByID(ctx, id, update); err != nil {
		return errors.Wrap(err, "Could not update the binotto MongoDB doc")
	}
	return nil
}
