package resolver

import (
	"context"

	"github.com/t0nyandre/go-graphql/internal/model"
	"github.com/t0nyandre/go-graphql/internal/service"
)

func (*Resolver) User(ctx context.Context, args *struct{ Username string }) (*userResolver, error) {
	user := &model.User{
		Username: args.Username,
	}
	user, err := ctx.Value("userService").(*service.UserService).GetUser(user)
	if err != nil {
		return nil, err
	}
	return &userResolver{user}, nil
}
