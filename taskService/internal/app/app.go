package app

import (
	"fmt"
	"github.com/RamzassH/LeadIt/libs/kafka"
	"github.com/RamzassH/LeadIt/libs/redis"
	grpcapp "github.com/RamzassH/LeadIt/taskService/internal/app/grpc"
	"github.com/RamzassH/LeadIt/taskService/internal/config"
	"github.com/RamzassH/LeadIt/taskService/internal/services/action"
	"github.com/RamzassH/LeadIt/taskService/internal/services/comment"
	"github.com/RamzassH/LeadIt/taskService/internal/services/progressBar"
	"github.com/RamzassH/LeadIt/taskService/internal/services/status"
	"github.com/RamzassH/LeadIt/taskService/internal/services/tag"
	"github.com/RamzassH/LeadIt/taskService/internal/services/task"
	"github.com/RamzassH/LeadIt/taskService/internal/storage/postgreSQL"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type App struct {
	GRPCServer *grpcapp.App
	validator  *validator.Validate
}

type Services struct {
	task.Task
	action.Action
	tag.Tag
	comment.Comment
	status.Status
	progressBar.ProgressBar
}

func New(logger zerolog.Logger,
	config *config.Config,
	validate *validator.Validate,
	db *sqlx.DB,
	redisClient *redis.Client,
	kafka *kafka.Producer) (*App, error) {

	taskStorage, err := postgreSQL.NewTaskStorage(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize task storage: %w", err)
	}
	actionStorage, err := postgreSQL.NewActionStorage(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize action storage: %w", err)
	}
	tagStorage, err := postgreSQL.NewTagStorage(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize tag storage: %w", err)
	}
	commentStorage, err := postgreSQL.NewCommentStorage(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize comment storage: %w", err)
	}
	statusStorage, err := postgreSQL.NewStatusStorage(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize status storage: %w", err)
	}

	rStorage, redisErr := redisStorage.New(redisClient, logger)

	if redisErr != nil {
		return nil, fmt.Errorf("failed to initialize redis storage: %w", redisErr)
	}

	taskService := task.New(
		logger,
		taskStorage,
		taskStorage,
		rStorage,
		kafka,
	)
	actionService := action.New(
		logger,
		actionStorage,
		actionStorage,
		rStorage,
		kafka)
	tagService := tag.New(
		logger,
		tagStorage,
		tagStorage,
		rStorage,
		kafka)
	commentService := comment.New(
		logger,
		commentStorage,
		commentStorage,
		rStorage,
		kafka)

	statusService := status.New(
		logger,
		statusStorage,
		statusStorage,
		rStorage,
		kafka)

	services := &Services{
		Task:    *taskService,
		Action:  *actionService,
		Tag:     *tagService,
		Comment: *commentService,
		Status:  *statusService,
	}

	grpcApp := grpcapp.New(logger, config, validate, services)

	return &App{
		GRPCServer: grpcApp,
		validator:  validate,
	}, nil
}
