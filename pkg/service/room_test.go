package service

import (
	"testing"

	. "github.com/architectv/property-task/pkg/error"
	"github.com/architectv/property-task/pkg/model"
	mock_repository "github.com/architectv/property-task/pkg/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRoomService_Create(t *testing.T) {
	type args struct {
		room *model.Room
	}
	type mockBehavior func(r *mock_repository.MockRoom, args args)

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
				room: &model.Room{
					Description: "test description",
					Price:       1000,
				},
			},
			mock: func(r *mock_repository.MockRoom, args args) {
				r.EXPECT().Create(args.room).Return(1, nil)
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Empty Description",
			input: args{
				room: &model.Room{
					Description: "",
					Price:       1000,
				},
			},
			mock:    func(r *mock_repository.MockRoom, args args) {},
			wantErr: true,
		},
		{
			name: "Wrong Price",
			input: args{
				room: &model.Room{
					Description: "test description",
					Price:       -1,
				},
			},
			mock:    func(r *mock_repository.MockRoom, args args) {},
			wantErr: true,
		},
		{
			name: "DB Error",
			input: args{
				room: &model.Room{
					Description: "test description",
					Price:       1000,
				},
			},
			mock: func(r *mock_repository.MockRoom, args args) {
				r.EXPECT().Create(args.room).Return(0, ErrInternalService)
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repository.NewMockRoom(c)
			test.mock(repo, test.input)
			s := &RoomService{repo: repo}

			got, err := s.Create(test.input.room)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}

func TestRoomService_Delete(t *testing.T) {
	type args struct {
		id int
	}
	type mockBehavior func(r *mock_repository.MockRoom, args args)

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
			mock: func(r *mock_repository.MockRoom, args args) {
				r.EXPECT().GetById(args.id).Return(&model.Room{}, nil)
				r.EXPECT().Delete(args.id).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Wrong Room Id",
			input: args{
				id: 1,
			},
			mock: func(r *mock_repository.MockRoom, args args) {
				r.EXPECT().GetById(args.id).Return(nil, ErrWrongRoomId)
			},
			wantErr: true,
		},
		{
			name: "DB Error",
			input: args{
				id: 1,
			},
			mock: func(r *mock_repository.MockRoom, args args) {
				r.EXPECT().GetById(args.id).Return(&model.Room{}, nil)
				r.EXPECT().Delete(args.id).Return(ErrInternalService)
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repository.NewMockRoom(c)
			test.mock(repo, test.input)
			s := &RoomService{repo: repo}

			err := s.Delete(test.input.id)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRoomService_GetAll(t *testing.T) {
	type args struct {
		sortField string
	}
	type mockBehavior func(r *mock_repository.MockRoom)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    []*model.Room
		wantErr bool
	}{
		{
			name: "Ok Sort By Id",
			input: args{
				sortField: "id",
			},
			mock: func(r *mock_repository.MockRoom) {
				rooms := []*model.Room{
					{Id: 1, Description: "description1", Price: 1000},
					{Id: 2, Description: "description2", Price: 5000},
					{Id: 3, Description: "description3", Price: 3000},
				}
				r.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(rooms, nil)
			},
			want: []*model.Room{
				{Id: 1, Description: "description1", Price: 1000},
				{Id: 2, Description: "description2", Price: 5000},
				{Id: 3, Description: "description3", Price: 3000},
			},
			wantErr: false,
		},
		{
			name: "Ok Sort By Price",
			input: args{
				sortField: "price",
			},
			mock: func(r *mock_repository.MockRoom) {
				rooms := []*model.Room{
					{Id: 1, Description: "description1", Price: 1000},
					{Id: 2, Description: "description2", Price: 5000},
					{Id: 3, Description: "description3", Price: 3000},
				}
				r.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(rooms, nil)
			},
			want: []*model.Room{
				{Id: 1, Description: "description1", Price: 1000},
				{Id: 2, Description: "description2", Price: 5000},
				{Id: 3, Description: "description3", Price: 3000},
			},
			wantErr: false,
		},
		{
			name: "Ok Sort By Id Reverse",
			input: args{
				sortField: "-id",
			},
			mock: func(r *mock_repository.MockRoom) {
				rooms := []*model.Room{
					{Id: 1, Description: "description1", Price: 1000},
					{Id: 3, Description: "description3", Price: 3000},
					{Id: 2, Description: "description2", Price: 5000},
				}
				r.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(rooms, nil)
			},
			want: []*model.Room{
				{Id: 1, Description: "description1", Price: 1000},
				{Id: 3, Description: "description3", Price: 3000},
				{Id: 2, Description: "description2", Price: 5000},
			},
			wantErr: false,
		},
		{
			name: "Ok Sort By Price Reverse",
			input: args{
				sortField: "-price",
			},
			mock: func(r *mock_repository.MockRoom) {
				rooms := []*model.Room{
					{Id: 2, Description: "description3", Price: 5000},
					{Id: 3, Description: "description2", Price: 3000},
					{Id: 1, Description: "description1", Price: 1000},
				}
				r.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(rooms, nil)
			},
			want: []*model.Room{
				{Id: 2, Description: "description3", Price: 5000},
				{Id: 3, Description: "description2", Price: 3000},
				{Id: 1, Description: "description1", Price: 1000},
			},
			wantErr: false,
		},
		{
			name: "Empty Sort Field",
			input: args{
				sortField: "",
			},
			mock: func(r *mock_repository.MockRoom) {
				rooms := []*model.Room{
					{Id: 1, Description: "description1", Price: 1000},
					{Id: 2, Description: "description2", Price: 5000},
					{Id: 3, Description: "description3", Price: 3000},
				}
				r.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(rooms, nil)
			},
			want: []*model.Room{
				{Id: 1, Description: "description1", Price: 1000},
				{Id: 2, Description: "description2", Price: 5000},
				{Id: 3, Description: "description3", Price: 3000},
			},
			wantErr: false,
		},
		{
			name: "Wrong Sort Field",
			input: args{
				sortField: "wrong",
			},
			mock:    func(r *mock_repository.MockRoom) {},
			wantErr: true,
		},
		{
			name: "Wrong Sort Field (Reverse)",
			input: args{
				sortField: "-wrong",
			},
			mock:    func(r *mock_repository.MockRoom) {},
			wantErr: true,
		},
		{
			name: "Wrong Sort Field (one symbol)",
			input: args{
				sortField: "w",
			},
			mock:    func(r *mock_repository.MockRoom) {},
			wantErr: true,
		},
		{
			name: "DB Error",
			input: args{
				sortField: "id",
			},
			mock: func(r *mock_repository.MockRoom) {
				r.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(nil, ErrInternalService)
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_repository.NewMockRoom(c)
			test.mock(repo)
			s := &RoomService{repo: repo}

			got, err := s.GetAll(test.input.sortField)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}
