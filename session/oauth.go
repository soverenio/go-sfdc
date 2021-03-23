package session

import (
	"fmt"
	"net/http"
)

type OAuthSession struct {
	URL string
	AccessToken string
	HTTPClient      *http.Client
	Version     int
}

func (s *OAuthSession) InstanceURL() string {
	return s.URL
}

func (s *OAuthSession) AuthorizationHeader(request *http.Request) {
	auth := fmt.Sprintf("Bearer %s", s.AccessToken)
	request.Header.Add("Authorization", auth)
}

func (s *OAuthSession) Client() *http.Client {
	return s.HTTPClient
}

func (s *OAuthSession) ServiceURL() string {
	return fmt.Sprintf("%s/services/data/v%d.0", s.URL, s.Version)
}
