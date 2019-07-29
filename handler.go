package main

import (
	"context"

	pb "github.com/dariosantos86/shippy-service-consignment/proto/consignment"
)

// Our grpc service handler
type handler struct {
	repository
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	// Find the next available vessel
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}

// Create a new vessel
func (s *handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err := s.repository.Create(req); err != nil {
		return err
	}
	res.Vessel = req
	return nil
}
