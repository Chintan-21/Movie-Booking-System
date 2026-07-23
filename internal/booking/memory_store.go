package booking

type MemoryStore struct {
	bookings map[string]Booking
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

func (ms *MemoryStore) Book(b Booking) (Booking, error) {
	if _, exists := ms.bookings[b.SeatID]; exists {
		return Booking{}, ErrSeatAlreadyBooked
	}
	ms.bookings[b.SeatID] = b
	return Booking{}, nil
}

func (ms *MemoryStore) ListBookings(movieID string) []Booking {
	var results []Booking
	for _, b := range ms.bookings {
		if b.MovieID == movieID {
			results = append(results, b)
		}
	}
	return results
}
