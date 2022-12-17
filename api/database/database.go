package database

import (
	"context"
	"fmt"
	"time"

	"github.com/DTreshy/szukaj-szpitala-api/api/models"
	"github.com/DTreshy/szukaj-szpitala-api/pkg/geocoding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Db struct {
	Client    *mongo.Client
	hospitals *mongo.Collection
}

var commandTimeout time.Duration = 5 * time.Second

func NewDb(user, passwd, uri string) (*Db, error) {
	credential := options.Credential{
		Username: user,
		Password: passwd,
	}

	dbclient, err := mongo.NewClient(options.Client().ApplyURI(uri).SetAuth(credential))
	if err != nil {
		return nil, err
	}

	return &Db{
		Client:    dbclient,
		hospitals: dbclient.Database("szukaj-szpitala").Collection("hospitals"),
	}, nil
}

func (db *Db) Connect(ctx context.Context) error {
	err := db.Client.Connect(ctx)
	if err != nil {
		return err
	}

	if err := db.Client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	return nil
}

func (db *Db) InsertHospital(ctx context.Context, hospital *models.Hospital) error {
	ctx, cancel := context.WithTimeout(ctx, commandTimeout)

	defer cancel()

	_, err := db.hospitals.InsertOne(ctx, hospital)
	if err != nil {
		return err
	}

	return nil
}

func (db *Db) QueryNearestHospitals(ctx context.Context, place *models.Place, number int) ([]models.HospitalInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, commandTimeout)
	matchingHospitals := []models.HospitalInfo{}

	defer cancel()

	cur, err := db.hospitals.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var curr *models.Hospital
	var largestDistanceMatching float64 = 0

	for {
		if ok := cur.Next(ctx); !ok {
			break
		}

		if err := bson.Unmarshal(cur.Current, curr); err != nil {
			return nil, fmt.Errorf("failed to unmarshall hospital: %w", err)
		}

		dist := geocoding.CalculateDistance(curr.Place.Coordinates, place.Coordinates)

		if len(matchingHospitals) < number {
			matchingHospitals = append(matchingHospitals, models.HospitalInfo{
				Hospital: curr,
				Distance: dist,
			})

			if largestDistanceMatching < dist {
				largestDistanceMatching = dist
			}

			continue
		}

		if largestDistanceMatching < dist {
			continue
		}

	}

	return matchingHospitals, nil
}
