package resolver

import (
	"fmt"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/t0nyandre/go-graphql/internal/model"
)

type userResolver struct {
	u *model.User
}

func (r *userResolver) ID() graphql.ID {
	return graphql.ID(r.u.ID)
}

func (r *userResolver) FirstName() string {
	return r.u.FirstName
}

func (r *userResolver) LastName() *string {
	return r.u.LastName
}

func (r *userResolver) FullName() string {
	if r.u.LastName != nil {
		lastName := *r.u.LastName
		return fmt.Sprintf("%s %s", r.u.FirstName, lastName)
	} else {
		return fmt.Sprintf("%s", r.u.FirstName)
	}
}

func (r *userResolver) Email() string {
	return r.u.Email
}

func (r *userResolver) IsConfirmed() bool {
	return r.u.Confirmed
}

func (r *userResolver) IsUpdated() bool {
	return r.u.IsUpdated()
}

func (r *userResolver) CreatedAt() graphql.Time {
	date, err := time.Parse(time.RFC822Z, r.u.CreatedAt)
	if err != nil {
		panic(err)
	}
	return graphql.Time{Time: date}
}

func (r *userResolver) UpdatedAt() *graphql.Time {
	if r.u.IsUpdated() {
		uat := *r.u.UpdatedAt
		date, err := time.Parse(time.RFC822Z, uat)
		if err != nil {
			panic(err)
		}
		return &graphql.Time{Time: date}
	}
	return nil
}
