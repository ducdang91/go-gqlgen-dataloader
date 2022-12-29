package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ducdang91/go-gqlgen-dataloader/graph/generated"
	"github.com/ducdang91/go-gqlgen-dataloader/graph/model"
	"github.com/ducdang91/go-gqlgen-dataloader/service"
)

func (r *mutationResolver) User(ctx context.Context) (*model.UserOps, error) {
	return &model.UserOps{}, nil
}

func (r *mutationResolver) Transaction(ctx context.Context) (*model.TransactionOps, error) {
	return &model.TransactionOps{}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return service.UserGetAll(ctx)
}

func (r *queryResolver) Transactions(ctx context.Context) ([]*model.Transaction, error) {
	return service.TransactionGetAll(ctx)
}

func (r *queryResolver) TransactionDetails(ctx context.Context) ([]*model.TransactionDetail, error) {
	return service.TransactionDetailGetAll(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
