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
