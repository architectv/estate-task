package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/architectv/property-task/pkg/model"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestRoomPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewRoomPostgres(db)

	type args struct {
		room *model.Room
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
				room: &model.Room{
					Description: "test description",
					Price:       1000,
				},
			},
			mock: func(args args) {
				room := args.room
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery(fmt.Sprintf("INSERT INTO %s", roomsTable)).
					WithArgs(room.Description, room.Price).
					WillReturnRows(rows)
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
			mock: func(args args) {
				room := args.room
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery(fmt.Sprintf("INSERT INTO %s", roomsTable)).
					WithArgs(room.Description, room.Price).
					WillReturnRows(rows)
			},
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
			mock: func(args args) {
				room := args.room
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery(fmt.Sprintf("INSERT INTO %s", roomsTable)).
					WithArgs(room.Description, room.Price).
					WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock(test.input)

			got, err := r.Create(test.input.room)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}

func TestRoomPostgres_Delete(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewRoomPostgres(db)

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
				mock.ExpectExec(fmt.Sprintf("DELETE FROM %s WHERE (.+)", roomsTable)).
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
				mock.ExpectExec(fmt.Sprintf("DELETE FROM %s WHERE (.+)", roomsTable)).
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

func TestRoomPostgres_GetAll(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewRoomPostgres(db)

	type args struct {
		sortField string
		desc      bool
	}
	type mockBehavior func()

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
				desc:      false,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "description", "price"}).
					AddRow(1, "description1", 1000).
					AddRow(2, "description2", 5000).
					AddRow(3, "description3", 3000)

				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s ORDER BY (.+)", roomsTable)).
					WillReturnRows(rows)
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
				desc:      false,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "description", "price"}).
					AddRow(1, "description1", 1000).
					AddRow(3, "description3", 3000).
					AddRow(2, "description2", 5000)
				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s ORDER BY (.+)", roomsTable)).
					WillReturnRows(rows)
			},
			want: []*model.Room{
				{Id: 1, Description: "description1", Price: 1000},
				{Id: 3, Description: "description3", Price: 3000},
				{Id: 2, Description: "description2", Price: 5000},
			},
			wantErr: false,
		},
		{
			name: "Ok Sort By Id Reverse",
			input: args{
				sortField: "id",
				desc:      true,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "description", "price"}).
					AddRow(3, "description3", 3000).
					AddRow(2, "description2", 5000).
					AddRow(1, "description1", 1000)

				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s ORDER BY (.+)", roomsTable)).
					WillReturnRows(rows)
			},
			want: []*model.Room{
				{Id: 3, Description: "description3", Price: 3000},
				{Id: 2, Description: "description2", Price: 5000},
				{Id: 1, Description: "description1", Price: 1000},
			},
			wantErr: false,
		},
		{
			name: "Ok Sort By Price Reverse",
			input: args{
				sortField: "price",
				desc:      true,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "description", "price"}).
					AddRow(2, "description2", 5000).
					AddRow(3, "description3", 3000).
					AddRow(1, "description1", 1000)

				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s ORDER BY (.+)", roomsTable)).
					WillReturnRows(rows)
			},
			want: []*model.Room{
				{Id: 2, Description: "description2", Price: 5000},
				{Id: 3, Description: "description3", Price: 3000},
				{Id: 1, Description: "description1", Price: 1000},
			},
			wantErr: false,
		},
		{
			name: "Empty Sosrt Field",
			input: args{
				sortField: "",
				desc:      false,
			},
			mock: func() {
				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s ORDER BY (.+)", roomsTable)).
					WillReturnError(errors.New("some error"))
			},
			wantErr: true,
		},
		{
			name: "Wrong Sosrt Field",
			input: args{
				sortField: "wrong",
				desc:      false,
			},
			mock: func() {
				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s ORDER BY (.+)", roomsTable)).
					WillReturnError(errors.New("some error"))
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()

			got, err := r.GetAll(test.input.sortField, test.input.desc)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}

func TestRoomPostgres_GetById(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewRoomPostgres(db)

	type args struct {
		id int
	}
	type mockBehavior func(args args)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    *model.Room
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				id: 1,
			},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"id", "description", "price"}).
					AddRow(1, "description1", 1000)

				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s WHERE (.+)", roomsTable)).
					WithArgs(args.id).WillReturnRows(rows)
			},
			want: &model.Room{
				Id:          1,
				Description: "description1",
				Price:       1000,
			},
			wantErr: false,
		},
		{
			name: "Not Found",
			input: args{
				id: 1,
			},
			mock: func(args args) {
				rows := sqlmock.NewRows([]string{"id", "description", "price"})

				mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM %s WHERE (.+)", roomsTable)).
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
