package agile

import (
	"context"
	"github.com/RamzassH/LeadIt/agileService/internal/domain/models"
	boardColumnV1 "github.com/RamzassH/LeadIt/libs/contracts/gen/agile/board_column"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerApi) CreateColumn(ctx context.Context, req *boardColumnV1.CreateColumnRequest) (*boardColumnV1.CreateColumnResponse, error) {
	payload := models.CreateBoardColumnDTO{
		BoardID: req.GetBoardId(),
		Name:    req.GetName(),
		Order:   req.GetOrder(),
	}
	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	id, err := s.service.CreateColumn(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create column: %v", err)
	}
	return &boardColumnV1.CreateColumnResponse{Id: id}, nil
}

func (s *ServerApi) GetColumnsByBoard(ctx context.Context, req *boardColumnV1.GetColumnsByBoardRequest) (*boardColumnV1.GetColumnsByBoardResponse, error) {
	cols, err := s.service.GetColumnsByBoard(ctx, req.GetBoardId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch columns: %v", err)
	}
	resp := make([]*boardColumnV1.BoardColumnType, 0, len(cols))
	for _, c := range cols {
		resp = append(resp, &boardColumnV1.BoardColumnType{
			Id: c.ID, Name: c.Name, BoardId: c.BoardID, Order: c.Order,
		})
	}
	return &boardColumnV1.GetColumnsByBoardResponse{Columns: resp}, nil
}

func (s *ServerApi) UpdateColumn(ctx context.Context, req *boardColumnV1.UpdateColumnRequest) (*boardColumnV1.BoardColumnType, error) {
	payload := models.UpdateBoardColumnDTO{
		ID:    req.Id,
		Name:  req.Name,
		Order: req.Order,
	}

	if err := s.ValidateStruct(payload); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	updated, err := s.service.UpdateColumn(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update column: %v", err)
	}

	return &boardColumnV1.BoardColumnType{
		Id:      updated.ID,
		Name:    updated.Name,
		BoardId: updated.BoardID,
		Order:   updated.Order,
	}, nil

}

func (s *ServerApi) DeleteColumn(ctx context.Context, req *boardColumnV1.DeleteColumnRequest) (*boardColumnV1.DeleteColumnResponse, error) {
	columnId := req.GetId()

	if columnId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "column id can not be empty")
	}

	rowsAffected, err := s.service.DeleteColumn(ctx, columnId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete column: %v", err)
	}

	return &boardColumnV1.DeleteColumnResponse{Id: rowsAffected}, nil
}
