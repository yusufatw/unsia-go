package controllers

import (
	"context"
	"unsia/pb/cities"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// City struct
type City struct{}

// GetCity function
func (s *City) GetCity(ctx context.Context, in *cities.City) (*cities.City, error) {
	return &cities.City{Id: 1, Name: "Solo"}, nil
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
