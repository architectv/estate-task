package handler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	. "github.com/architectv/estate-task/pkg/error"
	"github.com/architectv/estate-task/pkg/model"
	"github.com/architectv/estate-task/pkg/service"
	mock_service "github.com/architectv/estate-task/pkg/service/mock"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_createBooking(t *testing.T) {
	type mockBehavior func(r *mock_service.MockBooking, booking *model.Booking)

	tests := []struct {
		name                 string
		inputBody            string
		inputBooking         *model.Booking
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"room_id": 1, "date_start": "2021-01-05", "date_end": "2021-01-08"}`,
			inputBooking: &model.Booking{
				RoomId:    1,
				DateStart: time.Date(2021, time.January, 5, 0, 0, 0, 0, time.UTC),
				DateEnd:   time.Date(2021, time.January, 8, 0, 0, 0, 0, time.UTC),
			},
			mockBehavior: func(r *mock_service.MockBooking, booking *model.Booking) {
				r.EXPECT().Create(booking).Return(1, nil)
			},
			expectedStatusCode:   fiber.StatusOK,
			expectedResponseBody: `{"booking_id":1}`,
		},
		{
			name:                 "Empty Request Body",
			inputBody:            ``,
			inputBooking:         &model.Booking{},
			mockBehavior:         func(r *mock_service.MockBooking, booking *model.Booking) {},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: `{"error":"json: unexpected end of JSON input: "}`,
		},
		{
			name:      "Wrong Room Id",
			inputBody: `{"room_id": 1, "date_start": "2021-01-05", "date_end": "2021-01-08"}`,
			inputBooking: &model.Booking{
				RoomId:    1,
				DateStart: time.Date(2021, time.January, 5, 0, 0, 0, 0, time.UTC),
				DateEnd:   time.Date(2021, time.January, 8, 0, 0, 0, 0, time.UTC),
			},
			mockBehavior: func(r *mock_service.MockBooking, booking *model.Booking) {
				r.EXPECT().Create(booking).Return(0, ErrWrongRoomId)
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: fmt.Sprintf(`{"error":"%s"}`, ErrWrongRoomId),
		},
		{
			name:      "Service Error",
			inputBody: `{"room_id": 1, "date_start": "2021-01-05", "date_end": "2021-01-08"}`,
			inputBooking: &model.Booking{
				RoomId:    1,
				DateStart: time.Date(2021, time.January, 5, 0, 0, 0, 0, time.UTC),
				DateEnd:   time.Date(2021, time.January, 8, 0, 0, 0, 0, time.UTC),
			},
			mockBehavior: func(r *mock_service.MockBooking, booking *model.Booking) {
				r.EXPECT().Create(booking).Return(0, ErrInternalService)
			},
			expectedStatusCode:   fiber.StatusInternalServerError,
			expectedResponseBody: fmt.Sprintf(`{"error":"%s"}`, ErrInternalService),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockBooking(c)
			test.mockBehavior(repo, test.inputBooking)

			services := &service.Service{Booking: repo}
			handler := Handler{services}

			r := fiber.New()
			handler.InitRoutes(r)

			req := httptest.NewRequest(
				"POST",
				"/bookings/",
				bytes.NewBufferString(test.inputBody),
			)
			req.Header.Set("Content-type", "application/json")

			w, err := r.Test(req, -1)
			assert.Nil(t, err)

			bytesBody, err := ioutil.ReadAll(w.Body)
			assert.Nil(t, err)

			body := string(bytesBody)

			assert.Equal(t, test.expectedStatusCode, w.StatusCode)
			assert.Equal(t, test.expectedResponseBody, body)
		})
	}
}

func TestHandler_deleteBooking(t *testing.T) {
	type mockBehavior func(r *mock_service.MockBooking, bookingId int)

	tests := []struct {
		name                 string
		inputBookingId       int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:           "Ok",
			inputBookingId: 1,
			mockBehavior: func(r *mock_service.MockBooking, bookingId int) {
				r.EXPECT().Delete(bookingId).Return(nil)
			},
			expectedStatusCode:   fiber.StatusOK,
			expectedResponseBody: `"OK"`,
		},
		{
			name:           "Wrong Booking Id",
			inputBookingId: 1,
			mockBehavior: func(r *mock_service.MockBooking, bookingId int) {
				r.EXPECT().Delete(bookingId).Return(ErrWrongBookingId)
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: fmt.Sprintf(`{"error":"%s"}`, ErrWrongBookingId),
		},
		{
			name:           "Service Error",
			inputBookingId: 1,
			mockBehavior: func(r *mock_service.MockBooking, bookingId int) {
				r.EXPECT().Delete(bookingId).Return(ErrInternalService)
			},
			expectedStatusCode:   fiber.StatusInternalServerError,
			expectedResponseBody: fmt.Sprintf(`{"error":"%s"}`, ErrInternalService),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockBooking(c)
			test.mockBehavior(repo, test.inputBookingId)

			services := &service.Service{Booking: repo}
			handler := Handler{services}

			r := fiber.New()
			handler.InitRoutes(r)

			req := httptest.NewRequest(
				"DELETE",
				"/bookings/"+strconv.Itoa(test.inputBookingId),
				nil,
			)

			w, err := r.Test(req, -1)
			assert.Nil(t, err)

			bytesBody, err := ioutil.ReadAll(w.Body)
			assert.Nil(t, err)

			body := string(bytesBody)

			assert.Equal(t, test.expectedStatusCode, w.StatusCode)
			assert.Equal(t, test.expectedResponseBody, body)
		})
	}
}

func TestHandler_getBookingsByRoomId(t *testing.T) {
	type mockBehavior func(r *mock_service.MockBooking, roomId int)

	tests := []struct {
		name                 string
		inputRoomId          int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Ok",
			inputRoomId: 1,
			mockBehavior: func(r *mock_service.MockBooking, roomId int) {
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
				r.EXPECT().GetByRoomId(roomId).Return(bookings, nil)
			},
			expectedStatusCode: fiber.StatusOK,
			expectedResponseBody: `[{"booking_id":1,"date_start":"2021-01-05","date_end":"2021-01-08"},` +
				`{"booking_id":2,"date_start":"2021-01-25","date_end":"2021-01-28"}]`,
		},
		{
			name:        "Wrong Room Id",
			inputRoomId: 1,
			mockBehavior: func(r *mock_service.MockBooking, roomId int) {
				r.EXPECT().GetByRoomId(roomId).Return(nil, ErrWrongRoomId)
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: fmt.Sprintf(`{"error":"%s"}`, ErrWrongRoomId),
		},
		{
			name:        "Service Error",
			inputRoomId: 1,
			mockBehavior: func(r *mock_service.MockBooking, roomId int) {
				r.EXPECT().GetByRoomId(roomId).Return(nil, ErrInternalService)
			},
			expectedStatusCode:   fiber.StatusInternalServerError,
			expectedResponseBody: fmt.Sprintf(`{"error":"%s"}`, ErrInternalService),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockBooking(c)
			test.mockBehavior(repo, test.inputRoomId)

			services := &service.Service{Booking: repo}
			handler := Handler{services}

			r := fiber.New()
			handler.InitRoutes(r)

			req := httptest.NewRequest(
				"GET",
				"/bookings/?room_id="+strconv.Itoa(test.inputRoomId),
				nil,
			)

			w, err := r.Test(req, -1)
			assert.Nil(t, err)

			bytesBody, err := ioutil.ReadAll(w.Body)
			assert.Nil(t, err)

			body := string(bytesBody)

			assert.Equal(t, test.expectedStatusCode, w.StatusCode)
			assert.Equal(t, test.expectedResponseBody, body)
		})
	}
}
