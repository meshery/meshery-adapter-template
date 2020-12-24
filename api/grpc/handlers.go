package grpc

import (
	"context"

	"github.com/layer5io/meshery-adapter-template/proto"
)

// CreateMeshInstance is the handler function for the method CreateMeshInstance.
func (s *Service) CreateMeshInstance(ctx context.Context, req *proto.CreateMeshInstanceRequest) (*proto.CreateMeshInstanceResponse, error) {
	//TODO Need to modify to change the err check for different type logic
	err := s.Handler.CreateInstance(req.K8SConfig, req.ContextName)
	if err != nil {
		return nil, err
	}
	return &proto.CreateMeshInstanceResponse{}, nil
}

// MeshName is the handler function for the method MeshName.
func (s *Service) MeshName(ctx context.Context, req *proto.MeshNameRequest) (*proto.MeshNameResponse, error) {
	return &proto.MeshNameResponse{
		Name: s.Handler.GetName(),
	}, nil
}

// ApplyOperation is the handler function for the method ApplyOperation.
func (s *Service) ApplyOperation(ctx context.Context, req *proto.ApplyRuleRequest) (*proto.ApplyRuleResponse, error) {
	return &proto.ApplyRuleResponse{
		Error:       " ",
		OperationId: " ",
	}, nil
}

// SupportedOperations is the handler function for the method SupportedOperations.
func (s *Service) SupportedOperations(ctx context.Context, req *proto.SupportedOperationsRequest) (*proto.SupportedOperationsResponse, error) {
	result, err := s.Handler.ListOperations()
	if err != nil {
		return nil, err
	}

	operations := make([]*proto.SupportedOperation, 0)
	for key, val := range result {
		operations = append(operations, &proto.SupportedOperation{
			Key:      key,
			Value:    val.Properties["description"],
			Category: proto.OpCategory(proto.OpCategory_value[val.Type]),
		})
	}

	return &proto.SupportedOperationsResponse{
		Ops:   operations,
		Error: "none",
	}, nil
}

// StreamEvents is the handler function for the method StreamEvents.
func (s *Service) StreamEvents(ctx *proto.EventsRequest, srv proto.MeshService_StreamEventsServer) error {
	return nil
}
