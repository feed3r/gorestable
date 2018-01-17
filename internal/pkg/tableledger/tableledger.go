package tableledger

import (
	"errors"
	//	"github.com/google/uuid"
	"time"
)

const DefaultReservationTime time.Duration = 3 * time.Hour

// all tables are meant equal
type Table int

type TimeSlot struct {
	Begin time.Time
	End   time.Time
}

type Reservation struct {
	User     string
	Slot     TimeSlot
	Location Table
}

// must be unique. We use uuids, but this is one implementation detail.
// The only safe operation is compare for equality
type ReservationCode string

type reservationRecord struct {
	Reservation
	code ReservationCode
}

type TableLedger struct {
	// yep, we are going to use a plain slice. We expect this kind of load:
	// Order of magnitude:
	// - tables: 10s
	// - users: 10s
	// - active reservation records: 100s
	// With this expected load, a fancier data structure is out of
	// scope for the initial release(s)
	// Furthermore, we are going to purge elapsed record. A Record is elapsed
	// if its Slot.End refers to a past time (Slot.End < Now)
	records      []reservationRecord
	now          func() time.Time
	slotDuration time.Duration
	tables       uint
}

func New(tables uint, slotDuration time.Duration) (*TableLedger, error) {
	return NewWithTime(tables, slotDuration, time.Now)
}

func NewWithTime(tables uint, slotDuration time.Duration, now func() time.Time) (*TableLedger, error) {
	return &TableLedger{
		now:          now,
		slotDuration: slotDuration,
		tables:       tables,
	}, nil
}

func (tl *TableLedger) Size() int {
	return 0
}

func (tl *TableLedger) Purge() (int, error) {
	return 0, nil
}

func (tl *TableLedger) Add(r Reservation) (ReservationCode, error) {
	return "", errors.New("Not yet implemeted")
}

func (tl *TableLedger) Remove(user string, code ReservationCode) error {
	return errors.New("Not yet implemented")
}

func (tl *TableLedger) GetAll() ([]Reservation, error) {
	return nil, errors.New("Not yet implemented")
}

func (tl *TableLedger) GetByTimeSpan(begin time.Time, end time.Time) ([]Reservation, error) {
	return nil, errors.New("Not yet implemented")
}

func (tl *TableLedger) GetByUser(user string) ([]Reservation, error) {
	return nil, errors.New("Not yet implemented")
}
