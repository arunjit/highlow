package hlserver

import (
	"math/rand"
	"time"

	"github.com/arunjit/highlow"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const (
	secretKind = "Secret"
)

// Secret is the datastore type to store secret numbers and their range.
type Secret struct {
	Key     string `json:"key" datastore:"-"`
	Min     int    `json:"min" datastore:"min"`
	Max     int    `json:"max" datastore:"max"`
	Secret  int    `json:"secret" datastore:"secret"`
	Guesses int    `json:"guesses" datastore:"guesses"`
}

// NewSecret ...
func NewSecret(min, max int) *Secret {
	secret := highlow.GenerateSecret(min, max)
	guesses := highlow.CountGuesses(secret, min, max)
	return &Secret{
		Min:     min,
		Max:     max,
		Secret:  secret,
		Guesses: guesses,
	}
}

// Save ...
func (s *Secret) Save(c context.Context) (*Secret, error) {
	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, secretKind, nil), s)
	if err != nil {
		return nil, err
	}
	s.Key = key.Encode()
	return s, nil
}

// GetSecret ...
func GetSecret(c context.Context, keystr string) (*Secret, error) {
	key, err := datastore.DecodeKey(keystr)
	if err != nil {
		return nil, err
	}
	var s Secret
	err = datastore.Get(c, key, &s)
	if err != nil {
		return nil, err
	}
	s.Key = keystr
	return &s, nil
}

// GetAnySecret ...
func GetAnySecret(c context.Context, min, max int) (*Secret, error) {
	// TODO(arunjit): reduce 2 DS calls to one? This method will be called a lot.
	q := datastore.NewQuery(secretKind).
		Filter("min =", min).
		Filter("max =", max).
		KeysOnly()
	keys, err := q.GetAll(c, nil)
	if err != nil {
		return nil, err
	}
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(keys))
	key := keys[i]
	var s Secret
	err = datastore.Get(c, key, &s)
	if err != nil {
		return nil, err
	}
	s.Key = key.Encode()
	return &s, nil
}
