package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/ducdang91/go-gqlgen-dataloader/graph/model"

	"github.com/ducdang91/go-gqlgen-dataloader/service"
)

// Transactions is the resolver for the transactions field.
func (r *userResolver) Transactions(ctx context.Context, obj *model.User) ([]*model.Transaction, error) {
	return service.TransactionGetByUserID(ctx, obj.ID)
}

// Create is the resolver for the create field.
func (r *userOpsResolver) Create(ctx context.Context, obj *model.UserOps, input model.NewUser) (*model.User, error) {
	return service.UserCreate(ctx, input)
}

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

// UserOps returns UserOpsResolver implementation.
func (r *Resolver) UserOps() UserOpsResolver { return &userOpsResolver{r} }

type userResolver struct{ *Resolver }
type userOpsResolver struct{ *Resolver }
