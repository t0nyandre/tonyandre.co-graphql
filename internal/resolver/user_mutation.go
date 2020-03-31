package resolver

import (
	"context"

	"github.com/t0nyandre/go-graphql/internal/model"
	"github.com/t0nyandre/go-graphql/internal/service"
)

type UserInput struct {
	Username string
	Email    string
	Password string
}

func (*Resolver) Register(ctx context.Context, args *struct{ Input UserInput }) (*userResolver, error) {
	user := &model.User{
		Username: args.Input.Username,
		Email:    args.Input.Email,
		Password: args.Input.Password,
	}

	user, err := ctx.Value("userService").(*service.UserService).CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &userResolver{user}, err
}

func (*Resolver) ConfirmUser(ctx context.Context, args *struct{ Token string }) (bool, error) {
	token := args.Token

	success, err := ctx.Value("userService").(*service.UserService).ConfirmUser(token)
	if err != nil {
		return false, err
	}

	return success, err
}
