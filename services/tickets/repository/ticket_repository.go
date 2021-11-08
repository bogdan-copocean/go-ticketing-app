package repository

import (
	"context"
	"fmt"

	"github.com/bogdan-user/go-ticketing-app/pkg/errors"
	"github.com/bogdan-user/go-ticketing-app/services/tickets/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TicketsRepository interface {
	CreateTicket(*domain.Ticket) (*domain.Ticket, *errors.CustomErr)
	GetTicketById(string) (*domain.Ticket, *errors.CustomErr)
	GetAllTickets() ([]*domain.Ticket, *errors.CustomErr)
	UpdateTicket(*domain.Ticket, string) (*domain.Ticket, *errors.CustomErr)
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

func (ar *ticketsRepository) GetTicketById(id string) (*domain.Ticket, *errors.CustomErr) {
	ctx := context.Background()
	ticket := domain.Ticket{}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewBadRequestErr(err.Error())
	}

	findErr := ar.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&ticket)
	if findErr != nil {
		if findErr.Error() == "mongo: no documents in result" {
			return nil, errors.NewNotFoundErr(fmt.Sprintf("document with id %s not found", id))

		}
		return nil, errors.NewInternalServerErr(findErr.Error())
	}

	return &ticket, nil
}

func (ar *ticketsRepository) GetAllTickets() ([]*domain.Ticket, *errors.CustomErr) {
	ctx := context.Background()
	tickets := []*domain.Ticket{}

	cursor, findErr := ar.collection.Find(ctx, bson.M{})
	if findErr != nil {
		if findErr.Error() == "mongo: no documents in result" {
			return nil, errors.NewNotFoundErr("documents not found")

		}
		return nil, errors.NewInternalServerErr(findErr.Error())
	}
	defer cursor.Close(ctx)

	if cursorErr := cursor.All(context.Background(), &tickets); cursorErr != nil {
		return nil, errors.NewInternalServerErr(cursorErr.Error())
	}

	return tickets, nil
}

func (ar *ticketsRepository) UpdateTicket(ticket *domain.Ticket, foundId string) (*domain.Ticket, *errors.CustomErr) {
	ctx := context.Background()
	updatedTicket := domain.Ticket{}
	opts := options.FindOneAndUpdate()
	opts.SetReturnDocument(options.After)

	objectId, err := primitive.ObjectIDFromHex(foundId)
	if err != nil {
		return nil, errors.NewInternalServerErr(err.Error())
	}

	updateErr := ar.collection.FindOneAndUpdate(ctx, bson.M{"_id": objectId}, bson.M{"$set": bson.M{"title": ticket.Title, "price": ticket.Price}}, opts).Decode(&updatedTicket)
	if updateErr != nil {
		return nil, errors.NewInternalServerErr(updateErr.Error())
	}

	return &updatedTicket, nil
}
