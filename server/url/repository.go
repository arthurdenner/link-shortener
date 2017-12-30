package url

type memoryRepository struct {
	clicks map[string]int
	urls   map[string]*Url
}

func NewMemoryRepository() *memoryRepository {
	return &memoryRepository{
		make(map[string]int),
		make(map[string]*Url),
	}
}

func (r *memoryRepository) GetClicks(id string) int {
	return r.clicks[id]
}

func (r *memoryRepository) IdExists(id string) bool {
	_, exists := r.urls[id]

	return exists
}

func (r *memoryRepository) RegisterClick(id string) {
	r.clicks[id]++
}

func (r *memoryRepository) SearchById(id string) *Url {
	return r.urls[id]
}

func (r *memoryRepository) SearchByUrl(url string) *Url {
	for _, u := range r.urls {
		if u.Destiny == url {
			return u
		}
	}

	return nil
}

func (r *memoryRepository) Save(url Url) error {
	r.urls[url.Id] = &url

	return nil
}
