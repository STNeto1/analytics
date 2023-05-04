package auth

import (
	"_schemas/ent"
	"_schemas/ent/user"
	"context"

	"github.com/google/uuid"
)

func (r *AuthService) GetUserFromId(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return r.client.User.Query().Where(user.ID(id)).First(ctx)
}
