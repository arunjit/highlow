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

// GameRequest ...
type GameRequest struct {
	Min int `json:"min" endpoints:"d=1"`
	Max int `json:"max" endpoints:"d=1001"`
}

// GameResponse ...
type GameResponse struct {
	Key string `json:"key"`
	Min int    `json:"min"`
	Max int    `json:"max"`
}

// GuessRequest ...
type GuessRequest struct {
	Key   string `json:"key" endpoints:"req"`
	Guess int    `json:"guess" endpoints:"req"`
}

// GuessResponse ...
type GuessResponse struct {
	Result          string `json:"result"`
	ComputerGuesses int    `json:"computer_guesses"`
}

var (
	guessTooLow  = "LOW"
	guessTooHigh = "HIGH"
	guessExact   = "EQUAL"
)

// HighLowService is the main service for Secrets
type HighLowService struct{}

// Get ...
func (svc *HighLowService) Get(c context.Context, r *GetSecretRequest) (*Secret, error) {
	if r.Key != "" {
		return GetSecret(c, r.Key)
	}
	return GetAnySecret(c, r.Min, r.Max)
}

// New ...
func (svc *HighLowService) New(c context.Context, r *NewSecretRequest) (*Secret, error) {
	s := NewSecret(r.Min, r.Max)
	s, err := s.Save(c)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Game ...
func (svc *HighLowService) Game(c context.Context, r *GameRequest) (*GameResponse, error) {
	s, err := GetAnySecret(c, r.Min, r.Max)
	if err == ErrNoData {
		s = NewSecret(r.Min, r.Max)
		s, err = s.Save(c)
	}
	if err != nil {
		return nil, err
	}
	return &GameResponse{
		Key: s.Key,
		Min: s.Min,
		Max: s.Max,
	}, nil
}

// Guess ...
func (svc *HighLowService) Guess(c context.Context, r *GuessRequest) (*GuessResponse, error) {
	s, err := GetSecret(c, r.Key)
	if err != nil {
		return nil, err
	}
	res := &GuessResponse{"", s.Guesses}
	if r.Guess < s.Secret {
		res.Result = guessTooLow
	} else if r.Guess > s.Secret {
		res.Result = guessTooHigh
	} else {
		res.Result = guessExact
	}
	return res, nil
}

func init() {
	api, err := endpoints.RegisterService(&HighLowService{}, "highlow", "v1", "High-Low API", true)
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
	register("New", "POST", "secrets", "new", "Create a new secret")
	register("Game", "GET", "game", "game", "Get a number to guess")
	register("Guess", "GET", "guess", "guess", "Make a guess")

	endpoints.HandleHTTP()
}
