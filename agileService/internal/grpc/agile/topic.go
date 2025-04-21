package agile

import (
	"context"
	"github.com/RamzassH/LeadIt/agileService/internal/domain/models"
	topicV1 "github.com/RamzassH/LeadIt/libs/contracts/gen/agile/topic"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerApi) CreateTopic(ctx context.Context, req *topicV1.CreateTopicRequest) (*topicV1.CreateTopicResponse, error) {
	payload := models.CreateTopicDTO{
		ProjectID:   req.GetProjectId(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
	}
	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	id, err := s.service.CreateTopic(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create topic: %v", err)
	}
	return &topicV1.CreateTopicResponse{Id: id}, nil
}

func (s *ServerApi) GetTopic(ctx context.Context, req *topicV1.GetTopicRequest) (*topicV1.TopicType, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid topic id")
	}
	topic, err := s.service.GetTopic(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get topic: %v", err)
	}
	return &topicV1.TopicType{
		Id: topic.ID, ProjectId: topic.ProjectID, Name: topic.Name, Description: topic.Description,
	}, nil
}

func (s *ServerApi) GetTopicsByProject(ctx context.Context, req *topicV1.GetTopicsByProjectRequest) (*topicV1.GetTopicsByProjectResponse, error) {
	topics, err := s.service.GetTopicsByProject(ctx, req.GetProjectId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch topics: %v", err)
	}
	resp := make([]*topicV1.TopicType, 0, len(topics))
	for _, t := range topics {
		resp = append(resp, &topicV1.TopicType{
			Id: t.ID, ProjectId: t.ProjectID, Name: t.Name, Description: t.Description,
		})
	}
	return &topicV1.GetTopicsByProjectResponse{Topics: resp}, nil
}

func (s *ServerApi) UpdateTopic(ctx context.Context, req *topicV1.UpdateTopicRequest) (*topicV1.TopicType, error) {
	payload := models.UpdateTopicDTO{
		ID:          req.GetId(),
		Name:        &req.Name,
		Description: &req.Description,
	}
	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	topic, err := s.service.UpdateTopic(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update topic: %v", err)
	}
	return &topicV1.TopicType{
		Id: topic.ID, ProjectId: topic.ProjectID, Name: topic.Name, Description: topic.Description,
	}, nil
}

func (s *ServerApi) DeleteTopic(ctx context.Context, req *topicV1.DeleteTopicRequest) (*topicV1.DeleteTopicResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid topic id")
	}
	id, err := s.service.DeleteTopic(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete topic: %v", err)
	}
	return &topicV1.DeleteTopicResponse{Id: id}, nil
}
