package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/t0nyandre/go-graphql/internal/model"
)

type userResolver struct {
	u *model.User
}

func (r *userResolver) ID() graphql.ID {
	return graphql.ID(r.u.ID)
}

func (r *userResolver) Username() string {
	return r.u.Username
}

func (r *userResolver) Email() string {
	return r.u.Email
}

func (r *userResolver) IsConfirmed() bool {
	return r.u.Confirmed
}

func (r *userResolver) IsDisabled() bool {
	return r.u.Disabled
}

func (r *userResolver) IsUpdated() bool {
	return r.u.IsUpdated()
}

func (r *userResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.u.CreatedAt}
}

func (r *userResolver) UpdatedAt() *graphql.Time {
	if r.u.IsUpdated() {
		uat := *r.u.UpdatedAt
		return &graphql.Time{Time: uat}
	}
	return nil
}
