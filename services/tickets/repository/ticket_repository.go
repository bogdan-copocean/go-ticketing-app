package repository

import (
	"context"

	"github.com/bogdan-user/go-ticketing-app/pkg/errors"
	"github.com/bogdan-user/go-ticketing-app/services/tickets/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TicketsRepository interface {
	CreateTicket(*domain.Ticket) (*domain.Ticket, *errors.CustomErr)
}

type ticketsRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewTicketsRepository(client *mongo.Client, collection *mongo.Collection) TicketsRepository {
	return &ticketsRepository{client: client, collection: collection}
}

func (ar *ticketsRepository) CreateTicket(ticket *domain.Ticket) (*domain.Ticket, *errors.CustomErr) {
	ctx := context.Background()

	res, err := ar.collection.InsertOne(ctx, bson.M{"title": ticket.Title, "price": ticket.Price, "user_id": ticket.UserId})
	if err != nil {
		return nil, errors.NewInternalServerErr("could not insert document")
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.NewInternalServerErr("could not convert to object id")
	}
	ticket.Id = oid.Hex()

	return ticket, nil
}
