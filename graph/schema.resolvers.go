package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/Abdubek/auth-test/graph/service"

	"github.com/Abdubek/auth-test/graph/generated"
	"github.com/Abdubek/auth-test/graph/model"
)

func (r *mutationResolver) Users(ctx context.Context) (*model.UsersMutation, error) {
	return &model.UsersMutation{}, nil
}

func (r *queryResolver) Viewer(ctx context.Context) (*model.Viewer, error) {
	email := "user@example.com"

	return &model.Viewer{
		ID:    "1",
		Email: &email,
	}, nil
}

func (r *usersMutationResolver) Login(ctx context.Context, obj *model.UsersMutation, input *model.LoginInput) (*model.Logged, error) {
	email := "user@example.com"

	if input.Email == email && input.Password == "user123#" {
		token, err := service.CreateToken(1)
		if err != nil {
			return nil, err
		}
		return &model.Logged{
			Token: token,
			Viewer: &model.Viewer{
				ID:    "1",
				Email: &email,
			},
		}, nil
	}
	return nil, errors.New("INVALID_CREDENTIALS")
}

func (r *usersMutationResolver) Logout(ctx context.Context, obj *model.UsersMutation, refreshToken string) (*bool, error) {
	state := true
	return &state, nil
}

func (r *usersMutationResolver) Refresh(ctx context.Context, obj *model.UsersMutation, refreshToken string) (*model.Token, error) {
	token, err := service.CreateToken(1)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (r *viewerResolver) Sites(ctx context.Context, obj *model.Viewer) ([]*model.Sites, error) {
	var sites []*model.Sites

	host1 := "bugbounty.kz"
	host2 := "cybersec.kz"
	host3 := "wtotem.com"

	sites = append(sites, &model.Sites{
		ID:   "1",
		Host: &host1,
	})
	sites = append(sites, &model.Sites{
		ID:   "2",
		Host: &host2,
	})
	sites = append(sites, &model.Sites{
		ID:   "3",
		Host: &host3,
	})

	return sites, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// UsersMutation returns generated.UsersMutationResolver implementation.
func (r *Resolver) UsersMutation() generated.UsersMutationResolver { return &usersMutationResolver{r} }

// Viewer returns generated.ViewerResolver implementation.
func (r *Resolver) Viewer() generated.ViewerResolver { return &viewerResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type usersMutationResolver struct{ *Resolver }
type viewerResolver struct{ *Resolver }
