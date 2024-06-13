package repo

import (
	"context"
	"log"
	"time"

	"github.com/IraIvanishak/stack_app/internal/data"
	"github.com/IraIvanishak/stack_app/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type JobRepo struct {
	client *mongo.Client
}

func NewJobRepo(client *mongo.Client) *JobRepo {
	return &JobRepo{client: client}
}

func (r *JobRepo) SaveTest() error {
	_, err := r.client.Database("stack").Collection("job").InsertMany(context.Background(), []interface{}{
		model.Vacancy{
			Text:     data.TestData[0],
			Category: "Golang",
		},
		model.Vacancy{
			Text:     data.TestData[1],
			Category: "Golang",
		},
		model.Vacancy{
			Text:     data.TestData[2],
			Category: "Golang",
		},
	})
	return err
}
func (r *JobRepo) Save(job model.Vacancy) error {
	_, err := r.client.Database("stack").Collection("job").InsertOne(context.Background(), job)
	if err != nil {
		return err
	}
	return nil
}

func (r *JobRepo) List() ([]model.Vacancy, error) {
	// implement me
	return nil, nil
}

func (r *JobRepo) Delete(id string) error {
	// implement me
	return nil
}

func (r *JobRepo) GetRaw() ([]model.Vacancy, error) {
	filter := bson.D{{"requiredstack", primitive.Null{}}}
	cursor, err := r.client.Database("stack").Collection("job").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var results []model.Vacancy
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *JobRepo) SaveRaw(jobs []model.Vacancy) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, job := range jobs {
		filter := bson.M{"text": job.Text}
		update := bson.M{
			"$set": job,
		}

		opts := options.Update().SetUpsert(true)

		result, err := r.client.Database("stack").Collection("job").UpdateOne(ctx, filter, update, opts)
		if err != nil {
			return err
		}
		log.Printf("Modified count: %v, Upserted count: %v\n", result.ModifiedCount, result.UpsertedCount)
	}

	return nil
}

func (r *JobRepo) GetByCategory(category string) ([]model.Vacancy, error) {
	filter := bson.D{{"category", category}}
	cursor, err := r.client.Database("stack").Collection("job").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var results []model.Vacancy
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
