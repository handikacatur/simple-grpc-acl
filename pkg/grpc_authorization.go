package pkg

import (
	"context"
	"errors"
)

type GrpcServerAuthorization struct {
	GrpcServerAuthorizationInterface
	acl *GrpcAcl
}

type GrpcServerAuthorizationInterface interface {
	GetModelFromContext(ctx context.Context) RoleAndPermission
}

func (g *GrpcServerAuthorization) Authorize(ctx context.Context, permission *Permission, action ...string) error {
	model := g.GetModelFromContext(ctx)
	allowed := g.acl.CheckPermissionInModel(permission, model, action...)
	if !allowed {
		return errors.New("Current user doesn't has permission in this resource")
	}

	return nil
}

func (g *GrpcServerAuthorization) AuthorizeWithTeam(ctx context.Context, permission *Permission, teamId string, action ...string) {
	//model := g.GetModelFromContext(ctx)

}
