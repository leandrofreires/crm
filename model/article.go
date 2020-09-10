package model

import (
	"context"
	"time"

	"github.com/leandrofreires/crm/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const articleCollection = "article"

//Article is a model
type Article struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string               `bson:"title" json:"title" binding:"required"`
	Description string               `bson:"description" json:"description" binding:"required"`
	Content     string               `bson:"content" json:"content" binding:"required"`
	Thumbnail   string               `bson:"thumbnail" json:"thumbnail"`
	Keywords    []string             `bson:"keywords" json:"keywords" binding:"required"`
	Categories  []primitive.ObjectID `bson:"categories" json:"categories"`
	Slug        string               `bson:"slug" json:"slug"`
	DateTime
}

//Create save new article on database
func (a *Article) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	a.init()
	insertResult, err := database.Db.Collection(articleCollection).InsertOne(ctx, a)
	if err != nil {
		return err
	}
	a.ID = insertResult.InsertedID.(primitive.ObjectID)
	return nil
}

//GetArticles return a list of article
func (a *Article) GetArticles(page int64, limit int64) ([]Article, error) {
	var skip int64
	articles := make([]Article, 0)
	//pular na paginacao
	if page > 0 {
		skip = (page * limit) - limit
	}
	findWithPaginate := options.FindOptions{Limit: &limit, Skip: &skip}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := database.Db.Collection(articleCollection).Find(ctx, bson.D{}, &findWithPaginate)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var article Article
		err := cur.Decode(&article)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return articles, nil
}

//GetArticle find on model of article id
func (a *Article) GetArticle() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := database.Db.Collection(articleCollection).FindOne(ctx, bson.M{"_id": a.ID})
	if err := result.Decode(&a); err != nil {
		return err
	}
	return nil
}

//UpdateArticle on database
func (a *Article) UpdateArticle() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	a.init()
	result := database.Db.Collection(articleCollection).FindOneAndUpdate(ctx, bson.M{"_id": a.ID}, bson.M{"$set": a, "$currentDate": bson.M{"modified_at": true}})
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

//Delete from database
func (a *Article) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := database.Db.Collection(articleCollection).FindOneAndDelete(ctx, bson.M{"_id": a.ID})
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
