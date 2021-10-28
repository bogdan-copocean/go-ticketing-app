package repository

import (
	"context"

	"github.com/bogdan-user/go-ticketing-app/pkg/crypt"
	"github.com/bogdan-user/go-ticketing-app/pkg/errors"
	"github.com/bogdan-user/go-ticketing-app/services/auth/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository interface {
	CreateUser(*domain.User) (*domain.User, *errors.CustomErr)
	GetUserByEmail(string) (*domain.User, *errors.CustomErr)
}

type authRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewAuthRepository(client *mongo.Client, collection *mongo.Collection) AuthRepository {
	return &authRepository{client: client, collection: collection}
}

func (ar *authRepository) CreateUser(user *domain.User) (*domain.User, *errors.CustomErr) {
	ctx := context.Background()

	hashedPassword, hashErr := crypt.CreatePasswordHash(user.Password)
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

func (ar *authRepository) GetUserByEmail(email string) (*domain.User, *errors.CustomErr) {
	user := &domain.User{}
	ctx := context.Background()

	err := ar.collection.FindOne(ctx, bson.M{"email": email}).Decode(user)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, errors.NewNotFoundErr("email not found")
		}

		return nil, errors.NewInternalServerErr(err.Error())
	}

	return user, nil
}
