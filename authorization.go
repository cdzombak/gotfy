package gotfy

import "fmt"

type Authorization interface {
	Header() string
}

func AccessToken(token string) Authorization {
	return accessTokenAuth{token: token}
}

type accessTokenAuth struct {
	token string
}

func (a accessTokenAuth) Header() string {
	return fmt.Sprintf("Bearer %s", a.token)
}
