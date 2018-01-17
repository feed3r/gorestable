package tableledger

import "testing"

func TestCreatedEmpty(t *testing.T) {
	tl, err := New(3, DefaultReservationTime)
	if err != nil {
		t.Errorf("Unexpected error in creation: %s", err)
	}

	size := tl.Size()
	if size != 0 {
		t.Errorf("Unexpected size after creation: %d", size)
	}
}
