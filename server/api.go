package hlserver

import (
	"log"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	"golang.org/x/net/context"
)

// GetSecretRequest ...
type GetSecretRequest struct {
	Key string `json:"key"`
	Min int    `json:"min" endpoints:"d=1"`
	Max int    `json:"max" endpoints:"d=1001"`
}

// NewSecretRequest ...
type NewSecretRequest struct {
	Min int `json:"min" endpoints:"d=1"`
	Max int `json:"max" endpoints:"d=1001"`
}

// SecretService is the main service for Secrets
type SecretService struct{}

// Get ...
func (svc *SecretService) Get(c context.Context, r *GetSecretRequest) (*Secret, error) {
	if r.Key != "" {
		return GetSecret(c, r.Key)
	}
	return GetAnySecret(c, r.Min, r.Max)
}

// New ...
func (svc *SecretService) New(c context.Context, r *NewSecretRequest) (*Secret, error) {
	s := NewSecret(r.Min, r.Max)
	s, err := s.Save(c)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func init() {
	api, err := endpoints.RegisterService(&SecretService{}, "highlow", "v1", "High-Low API", true)
	if err != nil {
		log.Fatalf("RegisterService: %v", err)
	}

	register := func(method, httpMethod, path, name, desc string) {
		m := api.MethodByName(method).Info()
		m.HTTPMethod = httpMethod
		m.Path = path
		m.Name = name
		m.Desc = desc
	}

	register("Get", "GET", "secrets", "get", "Get a secret")
	register("New", "GET", "secrets:new", "new", "Get a new secret")

	endpoints.HandleHTTP()
}
