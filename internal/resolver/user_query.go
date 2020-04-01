package resolver

import (
	"context"
	"fmt"

	"github.com/t0nyandre/go-graphql/internal/model"
	"github.com/t0nyandre/go-graphql/internal/service"
)

type FindUserInput struct {
	Username *string
	Email    *string
}

func (*Resolver) User(ctx context.Context, args *struct{ Input FindUserInput }) (*userResolver, error) {
	user := &model.User{}
	if args.Input.Email == nil && args.Input.Username == nil {
		return nil, fmt.Errorf("You need to provide either a username or password to search for a user.")
	}
	if args.Input.Username == nil {
		user.Username = ""
	} else {
		user.Username = *args.Input.Username
	}
	if args.Input.Email == nil {
		user.Email = ""
	} else {
		user.Email = *args.Input.Email
	}
	user, err := ctx.Value("userService").(*service.UserService).FindUser(user)
	if err != nil {
		return nil, err
	}

	auth := ctx.Value("is_authorized").(bool)

	fmt.Printf("\nAm I logged in? %t\n", auth)
	return &userResolver{user}, nil
}
