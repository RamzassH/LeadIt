package task

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RamzassH/LeadIt/libs/kafka"
	redisStorage "github.com/RamzassH/LeadIt/libs/redis"
	"github.com/RamzassH/LeadIt/taskService/internal/domain/models"
	"github.com/rs/zerolog"
	"time"
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
	GetManyByProjectId(ctx context.Context, projectId int64) (tasList []*models.TaskDTO, err error)
	Update(ctx context.Context, payload models.UpdateTaskDTO) (updatedTask *models.TaskDTO, err error)
	ChangeStatus(ctx context.Context, payload models.ChangeStatusDTO) (int64, error)
	AddTag(ctx context.Context, payload models.AddTagDTO) (int64, error)
	RemoveTag(ctx context.Context, payload models.RemoveTagDTO) (int64, error)
	SetSolver(ctx context.Context, payload models.SetSolverDTO) (int64, error)
	SetIsActive(ctx context.Context, taskId int64) error
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

	_ = t.redisStorage.Del(ctx, t.TaskListKey(payload.ProjectID))

	return taskId, nil
}

func (t *Task) GetTask(ctx context.Context, taskId int64) (*models.TaskDTO, error) {
	const op = "Task.GetTask"
	logger := t.logger.With().Str("operation", op).Logger()

	logger.Info().
		Msg("getting task")

	key := t.TaskKey(taskId)

	cachedTask, err := t.redisStorage.Get(ctx, key)
	if err == nil && cachedTask != "" {
		var task models.TaskDTO
		if err := json.Unmarshal([]byte(cachedTask), &task); err == nil {
			logger.Debug().Int64("task_id", taskId).Msg("task loaded from cache")
			return &task, nil
		}
		logger.Warn().Err(err).Msg("failed to unmarshal cached task")
	}

	task, err := t.taskProvider.GetById(ctx, taskId)
	if err != nil {
		logger.Error().
			Err(err).
			Int64("task_id", taskId).
			Msg("failed to get task")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	data, err := json.Marshal(task)
	err = t.redisStorage.Set(ctx, key, data, 120*time.Minute)
	if err != nil {
		logger.Warn().Msg("failed to save task in cache")
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

	key := t.TaskListKey(projectId)

	cached, err := t.redisStorage.Get(ctx, key)
	if err == nil && cached != "" {
		var tasks []*models.TaskDTO
		if err := json.Unmarshal([]byte(cached), &tasks); err == nil {
			logger.Debug().Int64("project_id", projectId).Msg("tasks loaded from cache")
			return tasks, nil
		}
		logger.Warn().Err(err).Msg("failed to unmarshal cached task list")
	}

	tasks, err := t.taskProvider.GetManyByProjectId(ctx, projectId)
	if err != nil {
		logger.Error().Err(err).Int64("project_id", projectId).Msg("failed to get tasks for project")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if data, err := json.Marshal(tasks); err == nil {
		if err := t.redisStorage.Set(ctx, key, data, 120*time.Minute); err != nil {
			logger.Warn().Err(err).Msg("failed to cache task list")
		}
	} else {
		logger.Warn().Err(err).Msg("failed to marshal task")
	}

	logger.Info().Int64("project_id", projectId).Msg("Successfully fetched tasks")
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

	key := t.TaskKey(updatedTask.ID)

	data, err := json.Marshal(updatedTask)
	if err != nil {
		logger.Warn().Err(err).Msg("failed to marshal task")
	} else {
		err = t.redisStorage.Set(ctx, key, data, 120*time.Minute)
		if err != nil {
			logger.Warn().Err(err).Msg("failed to cache task list")
		}
	}

	logger.Info().
		Int64("task_id", updatedTask.ID).
		Str("task_name", updatedTask.Name).
		Msg("Successfully updated task")
	return updatedTask, nil
}

func (t *Task) ChangeStatus(ctx context.Context, payload models.ChangeStatusDTO) (int64, error) {
	const op = "Task.ChangeStatus"
	logger := t.logger.With().Str("operation", op).Logger()

	logger.Info().Msg("changing task status")

	statusId, err := t.taskProvider.ChangeStatus(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Int64("task_id", payload.TaskID).Msg("failed to change task status")
		return 0, err
	}

	t.RefreshTaskCache(ctx, payload.TaskID)

	logger.Info().Int64("task_id", statusId).Msg("Successfully changed task status")
	return statusId, nil
}

func (t *Task) AddTag(ctx context.Context, payload models.AddTagDTO) (int64, error) {
	const op = "Task.AddTag"
	logger := t.logger.With().Str("operation", op).Logger()

	logger.Info().Msg("adding tag")

	tagId, err := t.taskProvider.AddTag(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Int64("task_id", payload.TaskID).Msg("failed to add tag")
		return 0, err
	}

	t.RefreshTaskCache(ctx, payload.TaskID)

	logger.Info().Int64("task_id", tagId).Msg("Successfully added tag")
	return tagId, nil
}

func (t *Task) RemoveTag(ctx context.Context, payload models.RemoveTagDTO) (int64, error) {
	const op = "Task.RemoveTag"
	logger := t.logger.With().Str("operation", op).Logger()

	logger.Info().Msg("removing tag")

	rowsAffected, err := t.taskProvider.RemoveTag(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Int64("task_id", payload.TaskID).Msg("failed to remove tag")
		return 0, err
	}

	t.RefreshTaskCache(ctx, payload.TaskID)

	logger.Info().Int64("task_id", payload.TaskID).Msg("Successfully removed tag")
	return rowsAffected, nil
}

func (t *Task) SetSolver(ctx context.Context, payload models.SetSolverDTO) (int64, error) {
	const op = "Task.SetSolver"
	logger := t.logger.With().Str("operation", op).Logger()

	logger.Info().Msg("setting solver")

	solverId, err := t.taskProvider.SetSolver(ctx, payload)
	if err != nil {
		logger.Error().Err(err).Int64("task_id", payload.TaskID).Msg("failed to set solver")
		return 0, err
	}

	t.RefreshTaskCache(ctx, payload.TaskID)

	logger.Info().Int64("task_id", payload.TaskID).Msg("Successfully set solver")
	return solverId, nil
}

func (t *Task) SetIsActive(ctx context.Context, taskId int64) error {
	const op = "Task.SetTaskState"
	logger := t.logger.With().Str("operation", op).Logger()

	logger.Info().Msg("setting task state")

	err := t.taskProvider.SetIsActive(ctx, taskId)
	if err != nil {
		logger.Error().Err(err).Int64("task_id", taskId).Msg("failed to set task state")
		return err
	}

	t.RefreshTaskCache(ctx, taskId)

	logger.Info().Int64("task_id", taskId).Msg("Successfully set task state")
	return nil
}

func (t *Task) DeleteTask(ctx context.Context, taskId int64) (int64, error) {
	const op = "Task.DeleteTask"
	logger := t.logger.With().Str("operation", op).Int64("task_id", taskId).Logger()

	logger.Info().Msg("Starting deletion")

	task, err := t.taskProvider.GetById(ctx, taskId)
	if err != nil {
		logger.Error().Err(err).Int64("task_id", taskId).Msg("failed to delete task")
		return 0, err
	}
	projectId := task.ProjectID

	rowsAffected, err := t.taskProvider.Delete(ctx, taskId)
	if err != nil {
		logger.Error().Err(err).Msg("Deletion failed")
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	_ = t.redisStorage.Del(ctx, t.TaskKey(taskId))
	_ = t.redisStorage.Del(ctx, t.TaskListKey(projectId))

	logger.Info().Int64("rows_affected", rowsAffected).Msg("Task deleted")
	return rowsAffected, nil
}
