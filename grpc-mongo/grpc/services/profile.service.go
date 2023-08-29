package rpcService

import (
	"context"

	pro "grpc-mongo/grpc/profile"
	"grpc-mongo/interfaces"
	"grpc-mongo/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type RPCServer struct {
	pro.UnimplementedProfileServiceServer
}

var (
	ProfileCollection *mongo.Collection
	ProfileService    interfaces.IProfile
)

func (s *RPCServer) CreateProfile(ctx context.Context, req *pro.Profile) (*pro.ProfileResponse, error) {
	dbProfile := &models.Profile{Name: req.Name}
	result, err := ProfileService.CreateProfile(dbProfile)
	if err != nil {
		return nil, err
	} else {
		responseProfile := &pro.ProfileResponse{
			Name: result.Name,
		}
		return responseProfile, nil
	}
}
