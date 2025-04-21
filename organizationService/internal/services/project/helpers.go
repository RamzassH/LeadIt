package project

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func (p *Project) RefreshProjectCache(ctx context.Context, projectId int64) {
	key := fmt.Sprintf("project:%d", projectId)

	project, err := p.projectProvider.GetById(ctx, projectId)
	if err != nil {
		p.logger.Warn().Err(err).Int64("project_id", projectId).Msg("failed to refresh project cache: get")
		return
	}

	data, err := json.Marshal(project)
	if err != nil {
		p.logger.Warn().Err(err).Int64("project_id", projectId).Msg("failed to refresh project cache: marshal")
		return
	}
	if err := p.redisStorage.Set(ctx, key, data, 10*time.Minute); err != nil {
		p.logger.Warn().Err(err).Int64("project_id", projectId).Msg("failed to refresh project cache: set")
	}
}
