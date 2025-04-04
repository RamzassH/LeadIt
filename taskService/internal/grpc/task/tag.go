package task

import (
	"context"
	tagv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/tag"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s ServerAPI) CreateTag(ctx context.Context, req *tagv1.CreateTagRequest) (*tagv1.CreateTagResponse, error) {
	payload := models.CreateTagDTO{
		Name:      req.GetName(),
		TagColor:  req.GetColor(),
		ProjectID: req.GetProjectId(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	id, err := s.service.CreateTag(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create tag: %s", err.Error())
	}

	return &tagv1.CreateTagResponse{Id: id}, nil
}

func (s ServerAPI) GetTags(ctx context.Context, req *tagv1.GetTagsRequest) (*tagv1.GetTagsResponse, error) {
	projectID := req.GetProjectId()
	if projectID == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid project id")
	}

	tags, err := s.service.GetTagsForProject(ctx, projectID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get tags: %s", err.Error())
	}

	resp := make([]*tagv1.TagType, 0, len(tags))
	for _, tag := range tags {
		resp = append(resp, &tagv1.TagType{
			Id:        tag.ID,
			Name:      tag.Name,
			Color:     tag.TagColor,
			ProjectId: tag.ProjectID,
		})
	}

	return &tagv1.GetTagsResponse{Tags: resp}, nil
}

func (s ServerAPI) UpdateTag(ctx context.Context, req *tagv1.UpdateTagRequest) (*tagv1.UpdateTagResponse, error) {
	payload := models.UpdateTagDTO{
		ID:       req.GetId(),
		Name:     req.GetName(),
		TagColor: req.GetColor(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	updatedTag, err := s.service.UpdateTag(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update tag: %s", err.Error())
	}

	return &tagv1.UpdateTagResponse{
		Tag: &tagv1.TagType{
			Id:        updatedTag.ID,
			Name:      updatedTag.Name,
			Color:     updatedTag.TagColor,
			ProjectId: updatedTag.ProjectID,
		},
	}, nil
}

func (s ServerAPI) DeleteTag(ctx context.Context, req *tagv1.DeleteTagRequest) (*tagv1.DeleteTagResponse, error) {
	id := req.GetId()
	if id == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid tag id")
	}

	deletedID, err := s.service.DeleteTag(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete tag: %s", err.Error())
	}

	return &tagv1.DeleteTagResponse{Id: deletedID}, nil
}
