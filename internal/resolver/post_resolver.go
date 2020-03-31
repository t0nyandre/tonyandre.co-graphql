package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/t0nyandre/go-graphql/internal/model"
)

type postResolver struct {
	p *model.Post
}

func (r *postResolver) ID() graphql.ID {
	return graphql.ID(r.p.ID)
}

func (r *postResolver) Title() string {
	return r.p.Title
}

func (r *postResolver) Excerpt() *string {
	return r.p.Excerpt
}

func (r *postResolver) Text() string {
	return r.p.Text
}

func (r *postResolver) Image() *string {
	return r.p.Image
}

func (r *postResolver) Slug() string {
	return r.p.Slug
}

func (r *postResolver) IsArchived() bool {
	return r.p.IsArchived()
}

func (r *postResolver) IsUpdated() bool {
	return r.p.IsUpdated()
}

func (r *postResolver) IsPublished() bool {
	return r.p.IsPublished()
}

func (r *postResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.p.CreatedAt}
}

func (r *postResolver) PublishedAt() *graphql.Time {
	if r.p.IsPublished() {
		pat := *r.p.PublishedAt
		return &graphql.Time{Time: pat}
	}
	return nil
}

func (r *postResolver) UpdatedAt() *graphql.Time {
	if r.p.IsUpdated() {
		uat := *r.p.UpdatedAt
		return &graphql.Time{Time: uat}
	}
	return nil
}
