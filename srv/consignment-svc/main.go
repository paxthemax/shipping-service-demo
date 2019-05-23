package main

import (
	"context"
	"log"

	micro "github.com/micro/go-micro"
	pb "github.com/paxthemax/consignment-demo-svc/proto"
)

type consignmentRepository struct {
	consignments []*consignment
}

func (repo *consignmentRepository) create(consignment *consignment) (*consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *consignmentRepository) getAll() []*consignment {
	return repo.consignments
}

type service struct {
	repo consignmentRepository
}

func newService() *service {
	return &service{consignmentRepository{}}
}

// CreateConsignment will create a new consignment in the service repository.
func (svc *service) CreateConsignment(ctx context.Context, req *consignment, res *createResponse) error {
	consignment, err := svc.repo.create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}

// GetConsignments will return all consignments in the service repository.
func (svc *service) GetConsignments(ctx context.Context, req *getRequest, res *getResponse) error {
	consignments := svc.repo.getAll()

	res.Consignments = consignments
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	srv.Init()
	serviceHandler := newService()
	pb.RegisterShippingServiceHandler(srv.Server(), serviceHandler)

	if err := srv.Run(); err != nil {
		log.Printf("Internal error, err = %v", err)
	}

}
