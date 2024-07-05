package controllers

import (
	"context"
	"database/sql"
	"log"
	"unsia/models"
	"unsia/pb/cities"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// City struct
type City struct {
	DB  *sql.DB
	Log *log.Logger
	cities.UnimplementedCitiesServiceServer
}

// GetCity function
func (s *City) GetCity(ctx context.Context, in *cities.Id) (*cities.City, error) {
	var cityModel models.City
	cityModel.Log = s.Log
	err := cityModel.Get(ctx, s.DB, in)
	return &cityModel.Pb, err
}

func (s *City) GetCities(in *cities.EmptyMessage, stream cities.CitiesService_GetCitiesServer) error {

	query := `SELECT id, name FROM cities`
	row, err := s.DB.Query(query)
	if err != nil {
		return err
	}

	defer row.Close()
	for row.Next() {
		var city cities.City
		err = row.Scan(&city.Id, &city.Name)
		if err != nil {
			return err
		}

		res := &cities.CitiesStream{
			City: &city,
		}

		err := stream.Send(res)
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot send stream response: %v", err)
		}
	}

	return nil
}

func (s *City) Create(ctx context.Context, in *cities.CityInput) (*cities.City, error) {
	var cityModel models.City
	err := cityModel.Create(ctx, s.DB, in)
	return &cityModel.Pb, err
}

func (s *City) Delete(ctx context.Context, in *cities.Id) (*cities.MyBoolean, error) {
	var cityModel models.City
	err := cityModel.Delete(ctx, s.DB, in)
	if err != nil {
		return &cities.MyBoolean{Boolean: false}, err
	}
	return &cities.MyBoolean{Boolean: true}, nil
}
