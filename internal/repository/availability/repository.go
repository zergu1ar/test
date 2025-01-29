package availability

import (
	"applicationDesignTest/internal/entity"
	"applicationDesignTest/internal/helpers"
	"sync"
)

type Repository interface {
	GetAvailable() (entity.RoomAvailabilities, error)
	Save(available *entity.RoomAvailability) error
}

func NewRepository(cfg *Config) Repository {
	return &availability{
		cfg: cfg,
		data: map[uint]entity.RoomAvailability{
			1: {ID: 1, HotelID: "reddison", RoomID: "lux", Date: helpers.Date(2024, 1, 1), Quota: 1},
			2: {ID: 2, HotelID: "reddison", RoomID: "lux", Date: helpers.Date(2024, 1, 2), Quota: 1},
			3: {ID: 3, HotelID: "reddison", RoomID: "lux", Date: helpers.Date(2024, 1, 3), Quota: 1},
			4: {ID: 4, HotelID: "reddison", RoomID: "lux", Date: helpers.Date(2024, 1, 4), Quota: 1},
			5: {ID: 5, HotelID: "reddison", RoomID: "lux", Date: helpers.Date(2024, 1, 5)},
		},
		mtx: &sync.RWMutex{},
	}
}

type availability struct {
	cfg  *Config
	data map[uint]entity.RoomAvailability
	mtx  *sync.RWMutex
}

func (r *availability) GetAvailable() (entity.RoomAvailabilities, error) {
	r.mtx.RLock()
	res := make([]entity.RoomAvailability, 0, len(r.data))
	for _, v := range r.data {
		res = append(res, v)
	}
	defer r.mtx.RUnlock()
	return res, nil
}

func (r *availability) Save(available *entity.RoomAvailability) error {
	r.mtx.Lock()
	r.data[available.ID] = *available
	r.mtx.Unlock()
	return nil
}
