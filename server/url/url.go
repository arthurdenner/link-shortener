package url

import (
	"math/rand"
	"net/url"
	"time"
)

const (
	size    = 5
	simbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-+"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Url struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Destiny   string    `json:"destiny"`
}

type Repository interface {
	GetClicks(id string) int
	IdExists(id string) bool
	RegisterClick(id string)
	SearchById(id string) *Url
	SearchByUrl(url string) *Url
	Save(url Url) error
}

type Stats struct {
	Url    *Url `json:"url"`
	Clicks int  `json:"clicks"`
}

var repo Repository

func ConfigRepository(r Repository) {
	repo = r
}

func SearchOrCreateNewUrl(destiny string) (u *Url, isNew bool, err error) {
	if u = repo.SearchByUrl(destiny); u != nil {
		return u, false, nil
	}

	if _, err = url.ParseRequestURI(destiny); err != nil {
		return nil, false, err
	}

	uri := Url{generateId(), time.Now(), destiny}
	repo.Save(uri)

	return &uri, true, nil
}

func generateId() string {
	newId := func() string {
		id := make([]byte, size, size)
		for i := range id {
			id[i] = simbols[rand.Intn(len(simbols))]
		}

		return string(id)
	}

	for {
		if id := newId(); !repo.IdExists(id) {
			return id
		}
	}
}

func Search(id string) *Url {
	return repo.SearchById(id)
}

func RegisterClick(id string) {
	repo.RegisterClick(id)
}

func (u *Url) Stats() *Stats {
	clicks := repo.GetClicks(u.Id)

	return &Stats{u, clicks}
}
