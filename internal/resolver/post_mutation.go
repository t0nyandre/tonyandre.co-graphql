package resolver

import (
	"context"

	"github.com/gosimple/slug"
	"github.com/t0nyandre/go-graphql/internal/model"
	"github.com/t0nyandre/go-graphql/internal/service"
)

type PostInput struct {
	Title   string
	Excerpt *string
	Text    string
	Image   *string
	Slug    *string
}

func (*Resolver) CreatePost(ctx context.Context, args *struct{ Input PostInput }) (*postResolver, error) {
	post := &model.Post{
		Title:   args.Input.Title,
		Excerpt: args.Input.Excerpt,
		Text:    args.Input.Text,
		Image:   args.Input.Image,
	}

	if args.Input.Slug != nil {
		slug := *args.Input.Slug
		post.Slug = slug
	} else {
		post.Slug = slug.Make(args.Input.Title)
	}

	post, err := ctx.Value("postService").(*service.PostService).CreatePost(post)
	if err != nil {
		return nil, err
	}

	return &postResolver{post}, nil
}
