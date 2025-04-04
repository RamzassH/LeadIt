package task

import (
	"context"
	commentv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/comment"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s ServerAPI) AddComment(ctx context.Context, req *commentv1.CreateCommentRequest) (*commentv1.CreateCommentResponse, error) {
	payload := models.CreateCommentDTO{
		TaskID: req.GetTaskId(),
		UserID: req.GetUserId(),
		Body:   req.GetBody(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	commentID, err := s.service.CreateComment(ctx, payload)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add comment: %s", err.Error())
	}

	return &commentv1.CreateCommentResponse{
		CommentId: commentID,
	}, nil
}

func (s ServerAPI) GetCommentsByTask(ctx context.Context, req *commentv1.GetCommentsForTaskRequest) (*commentv1.GetCommentsForTaskResponse, error) {
	taskID := req.GetTaskId()

	if taskID == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid task id")
	}

	comments, err := s.service.GetCommentsForTask(ctx, taskID)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get comments: %s", err.Error())
	}

	var resp []*commentv1.CommentType
	for _, comment := range comments {
		resp = append(resp, &commentv1.CommentType{
			Id:     comment.ID,
			TaskId: comment.TaskID,
			Date:   timestamppb.New(comment.Date),
			UserId: comment.UserID,
			Body:   comment.Body,
		})
	}

	return &commentv1.GetCommentsForTaskResponse{Comments: resp}, nil
}

func (s ServerAPI) UpdateComment(ctx context.Context, req *commentv1.UpdateCommentRequest) (*commentv1.UpdateCommentResponse, error) {
	payload := models.UpdateCommentDTO{
		ID:   req.GetId(),
		Body: req.GetBody(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to validate update comment: %s", err.Error())
	}

	updatedComment, err := s.service.UpdateComment(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update comment: %s", err.Error())
	}

	return &commentv1.UpdateCommentResponse{
		UpdatedComment: &commentv1.CommentType{
			Id:     updatedComment.ID,
			TaskId: updatedComment.TaskID,
			Date:   timestamppb.New(updatedComment.Date),
			UserId: updatedComment.UserID,
			Body:   updatedComment.Body,
		},
	}, nil
}

func (s ServerAPI) DeleteComment(ctx context.Context, req *commentv1.DeleteCommentRequest) (*commentv1.DeleteCommentResponse, error) {
	commentId := req.GetId()

	if commentId == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid comment id")
	}

	deletedCommentId, err := s.service.DeleteComment(ctx, commentId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete comment: %s", err.Error())
	}

	return &commentv1.DeleteCommentResponse{
		Id: deletedCommentId,
	}, nil
}
