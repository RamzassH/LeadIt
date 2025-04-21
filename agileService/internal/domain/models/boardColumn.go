package models

type BoardColumnDTO struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	BoardID int64  `json:"board_id"`
	Order   int32  `json:"order"`
}

type CreateBoardColumnDTO struct {
	BoardID int64  `json:"board_id"`
	Name    string `json:"name"`
	Order   int32  `json:"order"`
}

type UpdateBoardColumnDTO struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Order int32  `json:"order"`
}
