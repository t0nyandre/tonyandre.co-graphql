package resolver

type tokenResolver struct {
	Access  *string
	Refresh *string
}

func (r *tokenResolver) AccessToken() *string {
	return r.Access
}

func (r *tokenResolver) RefreshToken() *string {
	return r.Refresh
}
