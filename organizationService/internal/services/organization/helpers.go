package organization

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func (org *Organization) RefreshOrganizationCache(ctx context.Context, organizationId int64) {
	organization, err := org.organizationProvider.GetById(ctx, organizationId)
	if err != nil {
		org.logger.Warn().Err(err).Int64("organization_id", organizationId).Msg("failed to refresh organization cache: get from provider")
		return
	}

	data, err := json.Marshal(organization)
	if err != nil {
		org.logger.Warn().Err(err).Int64("organization_id", organizationId).Msg("failed to marshal organization for cache")
		return
	}

	key := org.OrgKey(organizationId)
	if err := org.redisStorage.Set(ctx, key, data, 120*time.Minute); err != nil {
		org.logger.Warn().Err(err).Int64("organization_id", organizationId).Msg("failed to update organization cache")
	}
}

func (org *Organization) OrgKey(id int64) string { return fmt.Sprintf("organization:%d", id) }
func (org *Organization) OrgListKey(user int64) string {
	return fmt.Sprintf("organizer:%d:organizations", user)
}
