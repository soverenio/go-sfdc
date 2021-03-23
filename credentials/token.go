package credentials

import (
	"io"
	"strings"
)

type accessTokenProvider struct {
	creds AccessTokenCredentials
}

func (provider *accessTokenProvider) Retrieve() (io.Reader, error) {
	return strings.NewReader(provider.creds.Token), nil
}

func (provider *accessTokenProvider) URL() string {
	return provider.creds.URL
}

func (provider *accessTokenProvider) AccessToken() string {
	return provider.creds.Token
}
