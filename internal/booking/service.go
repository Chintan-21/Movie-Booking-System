package booking

import "context"

type Service struct {
	store BookingStore
}

func NewService(store BookingStore) *Service {
	return &Service{store}
}

func (bs *Service) Book(b Booking) (Booking, error) {
	return bs.store.Book(b)
}

func (bs *Service) ListBookings(movieID string) []Booking {
	return bs.store.ListBookings(movieID)
}
func (s *Service) ConfirmSeat(ctx context.Context, sessionID string, userID string) (Booking, error) {
	return s.store.Confirm(ctx, sessionID, userID)
}

func (s *Service) ReleaseSeat(ctx context.Context, sessionID string, userID string) error {
	return s.store.Release(ctx, sessionID, userID)
}
