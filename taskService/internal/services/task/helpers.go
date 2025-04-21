package task

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func (t *Task) RefreshTaskCache(ctx context.Context, taskId int64) {
	task, err := t.taskProvider.GetById(ctx, taskId)
	if err != nil {
		t.logger.Warn().Err(err).Int64("task_id", taskId).Msg("failed to refresh task cache: get")
		return
	}
	data, err := json.Marshal(task)
	if err != nil {
		t.logger.Warn().Err(err).Int64("task_id", taskId).Msg("failed to refresh task cache: marshal")
		return
	}

	key := fmt.Sprintf("task:%d", taskId)
	if err := t.redisStorage.Set(ctx, key, data, 10*time.Minute); err != nil {
		t.logger.Warn().Err(err).Int64("task_id", taskId).Msg("failed to refresh task cache: set")
	}
}

func (t *Task) TaskKey(taskId int64) string { return fmt.Sprintf("task:%d", taskId) }
func (t *Task) TaskListKey(projectId int64) string {
	return fmt.Sprintf("project:%d:tasks", projectId)
}
