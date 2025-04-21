package agile

import (
	"context"
	"github.com/RamzassH/LeadIt/agileService/internal/domain/models"
	taskTopicV1 "github.com/RamzassH/LeadIt/libs/contracts/gen/agile/task_topic"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerApi) AttachTopicToTask(ctx context.Context, req *taskTopicV1.TaskTopicRequest) (*taskTopicV1.TaskTopicResponse, error) {
	payload := models.TaskTopicDTO{TaskID: req.GetTaskId(), TopicID: req.GetTopicId()}
	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	success, err := s.service.AttachTopicToTask(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "attach topic failed: %v", err)
	}
	return &taskTopicV1.TaskTopicResponse{Success: success}, nil
}

func (s *ServerApi) DetachTopicFromTask(ctx context.Context, req *taskTopicV1.TaskTopicRequest) (*taskTopicV1.TaskTopicResponse, error) {
	payload := models.TaskTopicDTO{TaskID: req.GetTaskId(), TopicID: req.GetTopicId()}
	success, err := s.service.DetachTopicFromTask(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "detach topic failed: %v", err)
	}
	return &taskTopicV1.TaskTopicResponse{Success: success}, nil
}
