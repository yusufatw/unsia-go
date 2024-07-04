package controllers

import (
	"context"
	"database/sql"
	"unsia/models"
	"unsia/pb/cities"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// City struct
type City struct {
	DB *sql.DB
	cities.UnimplementedCitiesServiceServer
}

// GetCity function
func (s *City) GetCity(ctx context.Context, in *cities.Id) (*cities.City, error) {
	var cityModel models.City
	err := cityModel.Get(ctx, s.DB, in)
	return &cityModel.Pb, err
}

func (s *City) GetCities(in *cities.EmptyMessage, stream cities.CitiesService_GetCitiesServer) error {
	for i := 1; i < 50; i++ {
		res := &cities.CitiesStream{
			City: &cities.City{Id: int32(i), Name: "Jakarta"},
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
