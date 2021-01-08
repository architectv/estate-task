package service

import (
	"testing"
	"time"

	. "github.com/architectv/estate-task/pkg/error"
	"github.com/architectv/estate-task/pkg/model"
	mock_repository "github.com/architectv/estate-task/pkg/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBookingService_Create(t *testing.T) {
	type args struct {
		booking *model.Booking
	}
	type mockBehavior func(repo *mock_repository.MockBooking, roomRepo *mock_repository.MockRoom, args args)

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
			mock: func(repo *mock_repository.MockBooking, roomRepo *mock_repository.MockRoom, args args) {
				roomRepo.EXPECT().GetById(args.booking.RoomId).Return(&model.Room{}, nil)
				repo.EXPECT().Create(args.booking).Return(1, nil)
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
			mock: func(repo *mock_repository.MockBooking, roomRepo *mock_repository.MockRoom, args args) {
				roomRepo.EXPECT().GetById(args.booking.RoomId).Return(nil, ErrWrongRoomId)
			},
			wantErr: true,
		},
		{
			name: "Wrong Dates",
			input: args{
				booking: &model.Booking{
					RoomId:    1,
					DateStart: time.Date(2021, time.January, 9, 0, 0, 0, 0, time.UTC),
					DateEnd:   time.Date(2021, time.January, 8, 0, 0, 0, 0, time.UTC),
				},
			},
			mock: func(repo *mock_repository.MockBooking, roomRepo *mock_repository.MockRoom, args args) {
				roomRepo.EXPECT().GetById(args.booking.RoomId).Return(&model.Room{}, nil)
			},
			wantErr: true,
		},
		{
			name: "DB Error",
			input: args{
				booking: &model.Booking{
					RoomId:    1,
					DateStart: time.Date(2021, time.January, 5, 0, 0, 0, 0, time.UTC),
					DateEnd:   time.Date(2021, time.January, 8, 0, 0, 0, 0, time.UTC),
				},
			},
			mock: func(repo *mock_repository.MockBooking, roomRepo *mock_repository.MockRoom, args args) {
				roomRepo.EXPECT().GetById(args.booking.RoomId).Return(&model.Room{}, nil)
				repo.EXPECT().Create(args.booking).Return(0, ErrInternalService)
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repository.NewMockBooking(c)
			roomRepo := mock_repository.NewMockRoom(c)
			test.mock(repo, roomRepo, test.input)
			s := &BookingService{repo: repo, roomRepo: roomRepo}

			got, err := s.Create(test.input.booking)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}

func TestBookingService_Delete(t *testing.T) {
	type args struct {
		id int
	}
	type mockBehavior func(r *mock_repository.MockBooking, args args)

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
			mock: func(r *mock_repository.MockBooking, args args) {
				r.EXPECT().GetById(args.id).Return(&model.Booking{}, nil)
				r.EXPECT().Delete(args.id).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Wrong Booking Id",
			input: args{
				id: 1,
			},
			mock: func(r *mock_repository.MockBooking, args args) {
				r.EXPECT().GetById(args.id).Return(nil, ErrWrongBookingId)
			},
			wantErr: true,
		},
		{
			name: "DB Error",
			input: args{
				id: 1,
			},
			mock: func(r *mock_repository.MockBooking, args args) {
				r.EXPECT().GetById(args.id).Return(&model.Booking{}, nil)
				r.EXPECT().Delete(args.id).Return(ErrInternalService)
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repository.NewMockBooking(c)
			roomRepo := mock_repository.NewMockRoom(c)
			test.mock(repo, test.input)
			s := &BookingService{repo: repo, roomRepo: roomRepo}

			err := s.Delete(test.input.id)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRoomService_GetByRoomId(t *testing.T) {
	type args struct {
		roomId int
	}
	type mockBehavior func(repo *mock_repository.MockBooking, roomRepo *mock_repository.MockRoom, args args)

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
			mock: func(repo *mock_repository.MockBooking, roomRepo *mock_repository.MockRoom, args args) {
				roomRepo.EXPECT().GetById(args.roomId).Return(&model.Room{}, nil)
				bookings := []*model.Booking{
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
				}
				repo.EXPECT().GetByRoomId(args.roomId).Return(bookings, nil)
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
			name: "Wrong Room Id",
			input: args{
				roomId: 1,
			},
			mock: func(repo *mock_repository.MockBooking, roomRepo *mock_repository.MockRoom, args args) {
				roomRepo.EXPECT().GetById(args.roomId).Return(nil, ErrWrongRoomId)
			},
			wantErr: true,
		},
		{
			name: "DB Error",
			input: args{
				roomId: 1,
			},
			mock: func(repo *mock_repository.MockBooking, roomRepo *mock_repository.MockRoom, args args) {
				roomRepo.EXPECT().GetById(args.roomId).Return(&model.Room{}, nil)
				repo.EXPECT().GetByRoomId(args.roomId).Return(nil, ErrInternalService)
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repository.NewMockBooking(c)
			roomRepo := mock_repository.NewMockRoom(c)
			test.mock(repo, roomRepo, test.input)
			s := &BookingService{repo: repo, roomRepo: roomRepo}

			got, err := s.GetByRoomId(test.input.roomId)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}
