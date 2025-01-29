package entity

import "time"

type RoomAvailabilities []RoomAvailability

type RoomAvailability struct {
	ID      uint      `json:"id"`
	HotelID string    `json:"hotel_id"`
	RoomID  string    `json:"room_id"`
	Date    time.Time `json:"date"`
	Quota   int       `json:"quota"`
}

func (ra RoomAvailabilities) ToDaysMap() map[string]RoomAvailability {
	res := make(map[string]RoomAvailability)
	for _, available := range ra {
		res[available.Date.Format(time.DateOnly)] = available
	}
	return res
}
