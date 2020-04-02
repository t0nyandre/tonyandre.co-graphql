package resolver

import (
	"context"
	"fmt"

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

	profile, err := ctx.Value("profileService").(*service.ProfileService).CreateAuthor()
	if err != nil {
		return nil, err
	}

	user.Profile = *profile

	user, err = ctx.Value("userService").(*service.UserService).CreateUser(user)
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

func (*Resolver) Login(ctx context.Context, args *struct {
	Email    string
	Password string
}) (*tokenResolver, error) {
	ok := ctx.Value("is_authorized").(bool)
	if !ok {
		return nil, fmt.Errorf("Not authorized")
	}
	user := &model.User{
		Email: args.Email,
	}
	user, err := ctx.Value("userService").(*service.UserService).FindUser(user)
	if err != nil {
		return nil, err
	}

	if !user.Confirmed {
		return nil, fmt.Errorf("We need you to confirm this user. Please check your email for instructions.")
	}

	if user.Disabled {
		return nil, fmt.Errorf("This user is disabled and cannot use our service no more")
	}

	valid := user.VerifyPassword(args.Password)
	if !valid {
		return nil, fmt.Errorf("Wrong username and/or password")
	}

	tokens, err := ctx.Value("authService").(*service.AuthService).CreateTokens(user)
	if err != nil {
		return nil, err
	}

	err = ctx.Value("authService").(*service.AuthService).CreateAuth(user.ID, tokens)
	if err != nil {
		return nil, err
	}

	return &tokenResolver{Access: &tokens.AccessToken, Refresh: &tokens.RefreshToken}, nil
}

func (*Resolver) Logout(ctx context.Context) (bool, error) {
	accessid := ctx.Value("access_id").(*string)

	ok, err := ctx.Value("authService").(*service.AuthService).InvalidateAuth("AccessToken", *accessid)
	if !ok {
		return ok, err
	}
	return true, nil
}
