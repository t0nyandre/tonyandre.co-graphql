package resolver

import (
	"context"

	"github.com/t0nyandre/go-graphql/internal/service"
)

func (*Resolver) Posts(ctx context.Context) ([]*postResolver, error) {
	allPosts := make([]*postResolver, 0)
	posts, err := ctx.Value("postService").(*service.PostService).AllPosts()
	if err != nil {
		return nil, err
	}
	for _, post := range posts {
		allPosts = append(allPosts, &postResolver{post})
	}
	return allPosts, nil
}

func (*Resolver) Post(ctx context.Context, args *struct{ PostId string }) (*postResolver, error) {
	post, err := ctx.Value("postService").(*service.PostService).GetPost(args.PostId)
	if err != nil {
		return nil, err
	}
	return &postResolver{post}, nil
}
