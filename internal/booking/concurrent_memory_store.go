package booking

import "sync"

type ConcurrentMemoryStore struct {
	bookings map[string]Booking
	sync.RWMutex
}

func NewConcurrentMemoryStore() *ConcurrentMemoryStore {
	return &ConcurrentMemoryStore{
		bookings: map[string]Booking{},
	}
}

func (cms *ConcurrentMemoryStore) Book(b Booking) (Booking, error) {
	cms.Lock()
	defer cms.Unlock()

	if _, exists := cms.bookings[b.SeatID]; exists {
		return Booking{}, ErrSeatAlreadyBooked
	}
	cms.bookings[b.SeatID] = b
	return Booking{}, nil
}

func (cms *ConcurrentMemoryStore) ListBookings(movieID string) []Booking {
	cms.RLock()
	defer cms.RUnlock()

	var results []Booking
	for _, b := range cms.bookings {
		if b.MovieID == movieID {
			results = append(results, b)
		}
	}
	return results
}
