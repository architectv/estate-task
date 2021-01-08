package repository

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/architectv/estate-task/pkg/model"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestBookingPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewBookingPostgres(db)

	type args struct {
		booking *model.Booking
	}
	type mockBehavior func(args args)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				booking: &model.Booking{
					RoomId:    1,
					DateStart: time.Date(2021, time.January, 5, 0, 0, 0, 0, time.UTC),
					DateEnd:   time.Date(2021, time.January, 8, 0, 0, 0, 0, time.UTC),
				},
			},
			mock: func(args args) {
				booking := args.booking
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery(fmt.Sprintf("INSERT INTO %s", bookingsTable)).
					WithArgs(booking.RoomId, booking.DateStart, booking.DateEnd).
					WillReturnRows(rows)
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Wrong Room Id",
			input: args{
				booking: &model.Booking{
					RoomId:    1,
					DateStart: time.Date(2021, time.January, 5, 0, 0, 0, 0, time.UTC),
					DateEnd:   time.Date(2021, time.January, 8, 0, 0, 0, 0, time.UTC),
				},
			},
			mock: func(args args) {
				booking := args.booking
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery(fmt.Sprintf("INSERT INTO %s", bookingsTable)).
					WithArgs(booking.RoomId, booking.DateStart, booking.DateEnd).
					WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock(test.input)

			got, err := r.Create(test.input.booking)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}

func TestBookingPostgres_Delete(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewBookingPostgres(db)

	type args struct {
		id int
	}
	type mockBehavior func(args args)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				id: 1,
			},
			mock: func(args args) {
				mock.ExpectExec(fmt.Sprintf("DELETE FROM %s WHERE (.+)", bookingsTable)).
					WithArgs(args.id).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: false,
		},
		{
			name: "Not Found",
			input: args{
				id: 1,
			},
			mock: func(args args) {
				mock.ExpectExec(fmt.Sprintf("DELETE FROM %s WHERE (.+)", bookingsTable)).
					WithArgs(args.id).WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock(test.input)

			err := r.Delete(test.input.id)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestBookingPostgres_GetByRoomId(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewBookingPostgres(db)

	type args struct {
		roomId int
	}
	type mockBehavior func(args args)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    []*model.Booking
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				roomId: 1,
			},
			mock: func(args args) {
				dateStart1 := time.Date(2021, time.January, 5, 0, 0, 0, 0, time.UTC)
				dateEnd1 := time.Date(2021, time.January, 8, 0, 0, 0, 0, time.UTC)

				dateStart2 := time.Date(2021, time.January, 25, 0, 0, 0, 0, time.UTC)
				dateEnd2 := time.Date(2021, time.January, 28, 0, 0, 0, 0, time.UTC)

				rows := sqlmock.NewRows([]string{"id", "room_id", "date_start", "date_end"}).
					AddRow(1, 1, dateStart1, dateEnd1).
					AddRow(2, 1, dateStart2, dateEnd2)

				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s WHERE (.+)", bookingsTable)).
					WithArgs(args.roomId).WillReturnRows(rows)
			},
			want: []*model.Booking{
				{
					Id:        1,
					RoomId:    1,
					DateStart: time.Date(2021, time.January, 5, 0, 0, 0, 0, time.UTC),
					DateEnd:   time.Date(2021, time.January, 8, 0, 0, 0, 0, time.UTC),
				},
				{
					Id:        2,
					RoomId:    1,
					DateStart: time.Date(2021, time.January, 25, 0, 0, 0, 0, time.UTC),
					DateEnd:   time.Date(2021, time.January, 28, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: false,
		},
		{
			name: "Ok Empty List",
			input: args{
				roomId: 1,
			},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"id", "room_id", "date_start", "date_end"})

				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s WHERE (.+)", bookingsTable)).
					WithArgs(args.roomId).WillReturnRows(rows)
			},
			want:    nil,
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock(test.input)

			got, err := r.GetByRoomId(test.input.roomId)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}

func TestBookingPostgres_GetById(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewBookingPostgres(db)

	type args struct {
		id int
	}
	type mockBehavior func(args args)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    *model.Booking
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				id: 1,
			},
			mock: func(args args) {
				dateStart := time.Date(2021, time.January, 5, 0, 0, 0, 0, time.UTC)
				dateEnd := time.Date(2021, time.January, 8, 0, 0, 0, 0, time.UTC)
				rows := sqlmock.NewRows([]string{"id", "room_id", "date_start", "date_end"}).
					AddRow(1, 1, dateStart, dateEnd)

				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s WHERE (.+)", bookingsTable)).
					WithArgs(args.id).WillReturnRows(rows)
			},
			want: &model.Booking{
				Id:        1,
				RoomId:    1,
				DateStart: time.Date(2021, time.January, 5, 0, 0, 0, 0, time.UTC),
				DateEnd:   time.Date(2021, time.January, 8, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "Not Found",
			input: args{
				id: 1,
			},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"id", "room_id", "date_start", "date_end"})

				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s WHERE (.+)", bookingsTable)).
					WithArgs(args.id).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock(test.input)

			got, err := r.GetById(test.input.id)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}
