package task

import (
	"context"
	"fmt"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"github.com/rs/zerolog"
)

type Task struct {
	logger       zerolog.Logger
	taskSaver    Saver
	taskProvider Provider
	redisStorage redisStorage.RedisStore
	kafka        *kafka.Producer
}

type Saver interface {
	Save(ctx context.Context, payload models.CreateTaskDTO) (int64, error)
}
type Provider interface {
	GetById(ctx context.Context, taskId int64) (task *models.TaskDTO, err error)
	GetManyByProject(ctx context.Context, projectId int64) (tasks []*models.TaskDTO, err error)
	Update(ctx context.Context, payload models.UpdateTaskDTO) (updatedTask *models.TaskDTO, err error)
	Delete(ctx context.Context, taskId int64) (rowsAffected int64, err error)
}

func New(
	logger zerolog.Logger,
	taskSaver Saver,
	taskProvider Provider,
	redisStorage redisStorage.RedisStore,
	kafka *kafka.Producer,
) *Task {
	return &Task{
		logger:       logger,
		taskSaver:    taskSaver,
		taskProvider: taskProvider,
		redisStorage: redisStorage,
		kafka:        kafka,
	}
}

func (t *Task) CreateTask(ctx context.Context, payload models.CreateTaskDTO) (int64, error) {
	const op = "Task.CreateTask"
	logger := t.logger.With().Str("operation", op).Logger()

	logger.Info().
		Msg("creating task")

	taskId, err := t.taskSaver.Save(ctx, payload)
	if err != nil {
		logger.Error().
			Err(err).
			Int64("project_id", payload.ProjectID).
			Str("task_name", payload.Name).
			Msg("failed to save task")

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Int64("task_id", taskId).
		Str("task_name", payload.Name).
		Msg("Successfully saved task")
	return taskId, nil
}
func (t *Task) GetTask(ctx context.Context, taskId int64) (*models.TaskDTO, error) {
	const op = "Task.GetTask"
	logger := t.logger.With().Str("operation", op).Logger()

	logger.Info().
		Msg("getting task")

	task, err := t.taskProvider.GetById(ctx, taskId)
	if err != nil {
		logger.Error().
			Err(err).
			Int64("task_id", taskId).
			Msg("failed to get task")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Int64("task_id", task.ID).
		Str("task_name", task.Name).
		Msg("Successfully fetched task")
	return task, nil
}

func (t *Task) GetTasksForProject(ctx context.Context, projectId int64) ([]*models.TaskDTO, error) {
	const op = "Task.GetTasksForProject"
	logger := t.logger.With().Str("operation", op).Logger()

	logger.Info().
		Msg("getting tasks for project")

	tasks, err := t.taskProvider.GetManyByProject(ctx, projectId)
	if err != nil {
		logger.Error().Err(err).Int64("project_id", projectId).Msg("failed to get tasks for project")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Int64("project_id", projectId).
		Msg("Successfully fetched tasks")

	return tasks, nil
}

func (t *Task) UpdateTask(ctx context.Context, payload models.UpdateTaskDTO) (*models.TaskDTO, error) {
	const op = "Task.UpdateTask"
	logger := t.logger.With().Str("operation", op).Logger()

	logger.Info().
		Msg("updating task")

	updatedTask, err := t.taskProvider.Update(ctx, payload)
	if err != nil {
		logger.Error().
			Err(err).
			Int64("task_id", payload.ID).
			Str("task_name", payload.Name).
			Msg("failed to update task")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Int64("task_id", updatedTask.ID).
		Str("task_name", updatedTask.Name).
		Msg("Successfully updated task")
	return updatedTask, nil
}

func (t *Task) DeleteTask(ctx context.Context, taskId int64) (int64, error) {
	const op = "Task.DeleteTask"
	logger := t.logger.With().
		Str("operation", op).
		Int64("task_id", taskId).
		Logger()

	logger.Info().Msg("Starting deletion")

	rowsAffected, err := t.taskProvider.Delete(ctx, taskId)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Deletion failed")
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info().
		Int64("rows_affected", rowsAffected).
		Msg("Task deleted")

	return rowsAffected, nil
}
