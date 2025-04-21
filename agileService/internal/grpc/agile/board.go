package agile

import (
	"context"
	"github.com/RamzassH/LeadIt/agileService/internal/domain/models"
	boardV1 "github.com/RamzassH/LeadIt/libs/contracts/gen/agile/board"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerApi) CreateBoard(ctx context.Context, req *boardV1.CreateBoardRequest) (*boardV1.CreateBoardResponse, error) {
	payload := models.CreateBoardDTO{
		ProjectID:   req.GetProjectId(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Type:        req.GetType(),
	}
	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	id, err := s.service.CreateBoard(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create board: %v", err)
	}
	return &boardV1.CreateBoardResponse{Id: id}, nil
}

func (s *ServerApi) GetBoard(ctx context.Context, req *boardV1.GetBoardRequest) (*boardV1.BoardType, error) {
	boardId := req.GetId()

	if boardId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid board id")
	}

	board, err := s.service.GetBoard(ctx, boardId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get board: %v", err)
	}
	return &boardV1.BoardType{
		Id:          board.ID,
		ProjectId:   board.ProjectID,
		Name:        board.Name,
		Description: board.Description,
		Type:        board.Type,
	}, nil
}

func (s *ServerApi) UpdateBoard(ctx context.Context, req *boardV1.UpdateBoardRequest) (*boardV1.BoardType, error) {
	payload := models.UpdateBoardDTO{
		ID:          req.GetId(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Type:        req.GetType(),
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	board, err := s.service.UpdateBoard(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update board: %v", err)
	}

	return &boardV1.BoardType{
		Id:          board.ID,
		ProjectId:   board.ProjectID,
		Name:        board.Name,
		Description: board.Description,
		Type:        board.Type,
	}, nil
}

func (s *ServerApi) GetBoardsByProject(ctx context.Context, req *boardV1.GetBoardsByProjectRequest) (*boardV1.GetBoardsByProjectResponse, error) {
	boards, err := s.service.GetBoardsByProject(ctx, req.GetProjectId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch boards: %v", err)
	}
	resp := make([]*boardV1.BoardType, 0, len(boards))
	for _, b := range boards {
		resp = append(resp, &boardV1.BoardType{
			Id: b.ID, ProjectId: b.ProjectID, Name: b.Name, Description: b.Description, Type: b.Type,
		})
	}
	return &boardV1.GetBoardsByProjectResponse{Boards: resp}, nil
}

func (s *ServerApi) DeleteBoard(ctx context.Context, req *boardV1.DeleteBoardRequest) (*boardV1.DeleteBoardResponse, error) {
	boardId := req.GetId()
	if boardId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid board id")
	}

	rowsAffected, err := s.service.DeleteBoard(ctx, boardId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete board: %v", err)
	}

	return &boardV1.DeleteBoardResponse{
		Id: rowsAffected,
	}, nil
}
