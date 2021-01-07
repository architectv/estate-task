package handler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strconv"
	"testing"

	. "github.com/architectv/property-task/pkg/error"
	"github.com/architectv/property-task/pkg/model"
	"github.com/architectv/property-task/pkg/service"
	mock_service "github.com/architectv/property-task/pkg/service/mock"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_createRoom(t *testing.T) {
	type mockBehavior func(r *mock_service.MockRoom, room *model.Room)

	tests := []struct {
		name                 string
		inputBody            string
		inputRoom            *model.Room
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"description": "test description", "price": 1000}`,
			inputRoom: &model.Room{
				Description: "test description",
				Price:       1000,
			},
			mockBehavior: func(r *mock_service.MockRoom, room *model.Room) {
				r.EXPECT().Create(room).Return(1, nil)
			},
			expectedStatusCode:   fiber.StatusOK,
			expectedResponseBody: `{"room_id":1}`,
		},
		{
			name:                 "Empty Request Body",
			inputBody:            ``,
			inputRoom:            &model.Room{},
			mockBehavior:         func(r *mock_service.MockRoom, room *model.Room) {},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: `{"error":"json: unexpected end of JSON input: "}`,
		},
		{
			name:      "Empty Description",
			inputBody: `{"description": "", "price": 1000}`,
			inputRoom: &model.Room{
				Description: "",
				Price:       1000,
			},
			mockBehavior: func(r *mock_service.MockRoom, room *model.Room) {
				r.EXPECT().Create(room).Return(0, ErrEmptyDescription)
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: fmt.Sprintf(`{"error":"%s"}`, ErrEmptyDescription),
		},
		{
			name:      "Wrong Price",
			inputBody: `{"description": "test description", "price": -1}`,
			inputRoom: &model.Room{
				Description: "test description",
				Price:       -1,
			},
			mockBehavior: func(r *mock_service.MockRoom, room *model.Room) {
				r.EXPECT().Create(room).Return(0, ErrNotPositivePrice)
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: fmt.Sprintf(`{"error":"%s"}`, ErrNotPositivePrice),
		},
		{
			name:      "Service Error",
			inputBody: `{"description": "test description", "price": 1000}`,
			inputRoom: &model.Room{
				Description: "test description",
				Price:       1000,
			},
			mockBehavior: func(r *mock_service.MockRoom, room *model.Room) {
				r.EXPECT().Create(room).Return(0, ErrInternalService)
			},
			expectedStatusCode:   fiber.StatusInternalServerError,
			expectedResponseBody: fmt.Sprintf(`{"error":"%s"}`, ErrInternalService),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockRoom(c)
			test.mockBehavior(repo, test.inputRoom)

			services := &service.Service{Room: repo}
			handler := Handler{services}

			r := fiber.New()
			handler.InitRoutes(r)

			req := httptest.NewRequest(
				"POST",
				"/rooms/",
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

func TestHandler_deleteRoom(t *testing.T) {
	type mockBehavior func(r *mock_service.MockRoom, roomId int)

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
			mockBehavior: func(r *mock_service.MockRoom, roomId int) {
				r.EXPECT().Delete(roomId).Return(nil)
			},
			expectedStatusCode:   fiber.StatusOK,
			expectedResponseBody: "OK",
		},
		{
			name:        "Wrong Room Id",
			inputRoomId: 1,
			mockBehavior: func(r *mock_service.MockRoom, roomId int) {
				r.EXPECT().Delete(roomId).Return(ErrWrongRoomId)
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: fmt.Sprintf(`{"error":"%s"}`, ErrWrongRoomId),
		},
		{
			name:        "Service Error",
			inputRoomId: 1,
			mockBehavior: func(r *mock_service.MockRoom, roomId int) {
				r.EXPECT().Delete(roomId).Return(ErrInternalService)
			},
			expectedStatusCode:   fiber.StatusInternalServerError,
			expectedResponseBody: fmt.Sprintf(`{"error":"%s"}`, ErrInternalService),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockRoom(c)
			test.mockBehavior(repo, test.inputRoomId)

			services := &service.Service{Room: repo}
			handler := Handler{services}

			r := fiber.New()
			handler.InitRoutes(r)

			req := httptest.NewRequest(
				"DELETE",
				"/rooms/"+strconv.Itoa(test.inputRoomId),
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

func TestHandler_getAllRooms(t *testing.T) {
	type mockBehavior func(r *mock_service.MockRoom, sortField string)

	tests := []struct {
		name                 string
		inputSort            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputSort: "id",
			mockBehavior: func(r *mock_service.MockRoom, sortField string) {
				rooms := []*model.Room{
					{Id: 1, Description: "description1", Price: 1000},
					{Id: 2, Description: "description2", Price: 5000},
					{Id: 3, Description: "description3", Price: 3000},
				}
				r.EXPECT().GetAll(sortField).Return(rooms, nil)
			},
			expectedStatusCode: fiber.StatusOK,
			expectedResponseBody: `[{"room_id":1,"description":"description1","price":1000},` +
				`{"room_id":2,"description":"description2","price":5000},` +
				`{"room_id":3,"description":"description3","price":3000}]`,
		},
		{
			name:      "Ok Empty List",
			inputSort: "id",
			mockBehavior: func(r *mock_service.MockRoom, sortField string) {
				r.EXPECT().GetAll(sortField).Return(nil, nil)
			},
			expectedStatusCode:   fiber.StatusOK,
			expectedResponseBody: `null`,
		},
		{
			name:      "Wrong Sort Field",
			inputSort: "wrong",
			mockBehavior: func(r *mock_service.MockRoom, sortField string) {
				r.EXPECT().GetAll(sortField).Return(nil, ErrWrongSortField)
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: fmt.Sprintf(`{"error":"%s"}`, ErrWrongSortField),
		},
		{
			name:      "Service Error",
			inputSort: "id",
			mockBehavior: func(r *mock_service.MockRoom, sortField string) {
				r.EXPECT().GetAll(sortField).Return(nil, ErrInternalService)
			},
			expectedStatusCode:   fiber.StatusInternalServerError,
			expectedResponseBody: fmt.Sprintf(`{"error":"%s"}`, ErrInternalService),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockRoom(c)
			test.mockBehavior(repo, test.inputSort)

			services := &service.Service{Room: repo}
			handler := Handler{services}

			r := fiber.New()
			handler.InitRoutes(r)

			req := httptest.NewRequest(
				"GET",
				"/rooms/?sort="+test.inputSort,
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
