package repository

import (
	"context"

	"github.com/bogdan-user/go-ticketing-app/pkg/crypto"
	"github.com/bogdan-user/go-ticketing-app/pkg/errors"
	"github.com/bogdan-user/go-ticketing-app/services/auth/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TicketsRepository interface {
	// 	CreateUser(*domain.User) (*domain.User, *errors.CustomErr)
	// 	GetUserByEmail(*domain.User) (*domain.User, *errors.CustomErr)
}

type ticketsRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewTicketsRepository(client *mongo.Client, collection *mongo.Collection) TicketsRepository {
	return &ticketsRepository{client: client, collection: collection}
}

func (ar *ticketsRepository) CreateUser(user *domain.User) (*domain.User, *errors.CustomErr) {
	ctx := context.Background()

	hashedPassword, hashErr := crypto.CreatePasswordHash(user.Password)
	if hashErr != nil {
		return nil, errors.NewInternalServerErr("could not hash the password")
	}

	res, err := ar.collection.InsertOne(ctx, bson.M{"email": user.Email, "password": hashedPassword})
	if err != nil {
		return nil, errors.NewInternalServerErr("could not insert document")
	}

	user.Password = hashedPassword

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.NewInternalServerErr("could not convert to object id")
	}
	user.Id = oid.Hex()

	return user, nil
}

func (ar *ticketsRepository) GetUserByEmail(user *domain.User) (*domain.User, *errors.CustomErr) {
	foundUser := &domain.User{}
	ctx := context.Background()

	err := ar.collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(foundUser)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, errors.NewNotFoundErr("email not found")
		}

		return nil, errors.NewInternalServerErr(err.Error())
	}

	if cryptErr := crypto.CompareHashWithPassword(foundUser.Password, user.Password); cryptErr != nil {
		return nil, errors.NewBadRequestErr("invalid credentials")
	}

	return user, nil
}
