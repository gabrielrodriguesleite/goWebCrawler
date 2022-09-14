package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindAllLinks() (entities []VisitedLink, err error) {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database("crawler").Collection("links")

	// opção para realizar a busca e retornar com ordenação decrescente (último primeiro)
	opts := options.Find().SetSort(bson.D{{Key: "visited_date", Value: -1}})
	cursor, err := c.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return
	}

	err = cursor.All(context.TODO(), &entities)
	return
}
