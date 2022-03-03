package repository

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/sigit14ap/go-kumparan/internal/domain"
	"github.com/sigit14ap/go-kumparan/internal/domain/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type ArticlesRepo struct {
	db *mongo.Collection
}

func (repo *ArticlesRepo) Get(ctx context.Context, query dto.SearchArticleDTO) ([]domain.Article, error) {
	searchQuery := bson.D{}

	if query.Author != "" {
		searchQuery = append(searchQuery, bson.E{
			"author", bson.M{"$regex": query.Author, "$options": "im"},
		})
	}

	if query.Search != "" {

		searchQuery = append(searchQuery, bson.E{
			"$or", bson.A{
				bson.M{"title": bson.M{"$regex": query.Search, "$options": "im"}},
				bson.M{"body": bson.M{"$regex": query.Search, "$options": "im"}},
			},
		})

	}

	findOptions := options.Find().SetSort(bson.D{{"created", -1}})

	cursor, err := repo.db.Find(ctx, searchQuery, findOptions)
	if err != nil {
		return nil, err
	}

	data := []domain.Article{}
	err = cursor.All(ctx, &data)
	return data, err
}

func (repo *ArticlesRepo) Find(ctx context.Context, articleID primitive.ObjectID) (domain.Article, error) {
	result := repo.db.FindOne(ctx, bson.M{"_id": articleID})

	var article domain.Article
	err := result.Decode(&article)

	return article, err
}

func (repo *ArticlesRepo) Create(ctx context.Context, article dto.ArticleDTO) (domain.Article, error) {
	result, err := repo.db.InsertOne(ctx, article)

	data := domain.Article{}
	copier.Copy(&data, &article)
	data.ID = result.InsertedID.(primitive.ObjectID)

	return data, err
}

func NewArticlesRepo(db *mongo.Database) *ArticlesRepo {
	collection := db.Collection(articlesCollection)
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"title": 1},
		Options: options.Index().SetUnique(false),
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatalf("unable to create articles collection index, %v", err)
	}

	return &ArticlesRepo{
		db: collection,
	}
}
