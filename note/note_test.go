package note

import (
	"context"
	"database/sql"
	"testing"

	mockdb "shadowsoni666/db/mock"
	db "shadowsoni666/db/sqlc"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {
	args := db.RegisterUserParams{
		Username: "user1",
		Password: "password1",
		Email:    "user1@user.com",
	}

	testCases := []struct {
		name              string
		mockdbCreateUser  func(mockdb *mockdb.MockQuerier, args *db.RegisterUserParams)
		checkReturnValues func(t *testing.T, username string, err error)
	}{
		{
			name: "user registration OK",
			mockdbCreateUser: func(mockdb *mockdb.MockQuerier, args *db.RegisterUserParams) {
				mockdb.EXPECT().RegisterUser(gomock.Any(), args).Times(1).Return(args.Username, nil)
			},
			checkReturnValues: func(t *testing.T, username string, err error) {
				require.Equal(t, username, args.Username)
				require.Nil(t, err)
			},
		},
		{
			name: "user registration returns ErrUserAlreadyExists",
			mockdbCreateUser: func(mockdb *mockdb.MockQuerier, args *db.RegisterUserParams) {
				mockdb.EXPECT().RegisterUser(gomock.Any(), args).Times(1).Return("", ErrUserAlreadyExists)
			},
			checkReturnValues: func(t *testing.T, username string, err error) {
				require.ErrorIs(t, err, ErrUserAlreadyExists)
				require.Empty(t, username)
			},
		},
		{
			name: "user registration returns ErrDBInternal",
			mockdbCreateUser: func(mockdb *mockdb.MockQuerier, args *db.RegisterUserParams) {
				mockdb.EXPECT().RegisterUser(gomock.Any(), args).Times(1).Return("", ErrDBInternal)
			},
			checkReturnValues: func(t *testing.T, username string, err error) {
				require.ErrorIs(t, err, ErrDBInternal)
				require.Empty(t, username)
			},
		},
	}

	for c := range testCases {
		tc := testCases[c]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockdb := mockdb.NewMockQuerier(ctrl)
			ns := NewService(mockdb)

			tc.mockdbCreateUser(mockdb, &args)
			u, err := ns.q.RegisterUser(context.Background(), &args)
			tc.checkReturnValues(t, u, err)
		})
	}
}

func TestGetUser(t *testing.T) {
	const username = "user1"

	testCases := []struct {
		name              string
		username          string
		mockdbGetUser     func(mockdb *mockdb.MockQuerier)
		checkReturnValues func(t *testing.T, user db.User, err error)
	}{
		{
			name:     "getting user OK",
			username: username,
			mockdbGetUser: func(mockdb *mockdb.MockQuerier) {
				mockdb.EXPECT().GetUser(gomock.Any(), username).Times(1).Return(db.User{Username: username}, nil)
			},
			checkReturnValues: func(t *testing.T, user db.User, err error) {
				require.Equal(t, user.Username, username)
				require.Nil(t, err)
			},
		},
		{
			name:     "getting user returns ErrUserNotFound",
			username: username,
			mockdbGetUser: func(mockdb *mockdb.MockQuerier) {
				mockdb.EXPECT().GetUser(gomock.Any(), username).Times(1).Return(db.User{}, ErrUserNotFound)
			},
			checkReturnValues: func(t *testing.T, user db.User, err error) {
				require.Empty(t, user.Email)
				require.Empty(t, user.Password)
				require.Empty(t, user.Username)
				require.ErrorIs(t, err, ErrUserNotFound)
			},
		},
		{
			name:     "getting user returns ErrDBInternal",
			username: username,
			mockdbGetUser: func(mockdb *mockdb.MockQuerier) {
				mockdb.EXPECT().GetUser(gomock.Any(), username).Times(1).Return(db.User{}, ErrDBInternal)
			},
			checkReturnValues: func(t *testing.T, user db.User, err error) {
				require.Empty(t, user.Email)
				require.Empty(t, user.Password)
				require.Empty(t, user.Username)
				require.ErrorIs(t, err, ErrDBInternal)
			},
		},
	}

	for c := range testCases {
		tc := testCases[c]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockdb := mockdb.NewMockQuerier(ctrl)
			ns := NewService(mockdb)

			tc.mockdbGetUser(mockdb)
			user, err := ns.q.GetUser(context.Background(), tc.username)
			tc.checkReturnValues(t, user, err)
		})
	}
}

func TestCreateNote(t *testing.T) {
	args := db.CreateNoteParams{
		ID:       uuid.New(),
		Title:    "testtitle1",
		Username: "testuser1",
		Text:     sql.NullString{String: "testtext", Valid: true},
	}

	testCases := []struct {
		name              string
		mockdbCreateNote  func(mockdb *mockdb.MockQuerier, args *db.CreateNoteParams)
		checkReturnValues func(t *testing.T, args *db.CreateNoteParams, id uuid.UUID, err error)
	}{
		{
			name: "creating note OK",
			mockdbCreateNote: func(mockdb *mockdb.MockQuerier, args *db.CreateNoteParams) {
				mockdb.EXPECT().CreateNote(gomock.Any(), args).Times(1).Return(args.ID, nil)
			},
			checkReturnValues: func(t *testing.T, args *db.CreateNoteParams, id uuid.UUID, err error) {
				require.Equal(t, args.ID, id)
				require.Nil(t, err)
			},
		},
		{
			name: "creating note returns ErrAlreadyExist",
			mockdbCreateNote: func(mockdb *mockdb.MockQuerier, args *db.CreateNoteParams) {
				mockdb.EXPECT().CreateNote(gomock.Any(), args).Times(1).Return(uuid.Nil, ErrAlreadyExists)
			},
			checkReturnValues: func(t *testing.T, args *db.CreateNoteParams, id uuid.UUID, err error) {
				require.Equal(t, id, uuid.Nil)
				require.ErrorIs(t, err, ErrAlreadyExists)
			},
		},
		{
			name: "creating note returns ErrDBInternal",
			mockdbCreateNote: func(mockdb *mockdb.MockQuerier, args *db.CreateNoteParams) {
				mockdb.EXPECT().CreateNote(gomock.Any(), args).Times(1).Return(uuid.Nil, ErrDBInternal)
			},
			checkReturnValues: func(t *testing.T, args *db.CreateNoteParams, id uuid.UUID, err error) {
				require.Equal(t, id, uuid.Nil)
				require.ErrorIs(t, err, ErrDBInternal)
			},
		},
	}

	for c := range testCases {
		tc := testCases[c]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockdb := mockdb.NewMockQuerier(ctrl)
			ns := NewService(mockdb)

			tc.mockdbCreateNote(mockdb, &args)
			id, err := ns.q.CreateNote(context.Background(), &args)
			tc.checkReturnValues(t, &args, id, err)
		})
	}
}

func TestGetAllNotesFromUser(t *testing.T) {
	const username = "user1"

	testCases := []struct {
		name                   string
		username               string
		mockdbGetNotesFromUser func(mockdb *mockdb.MockQuerier)
		checkReturnValues      func(t *testing.T, notes []db.Note, err error)
	}{
		{
			name:     "getting notes from user OK",
			username: username,
			mockdbGetNotesFromUser: func(mockdb *mockdb.MockQuerier) {
				mockdb.EXPECT().GetAllNotesFromUser(gomock.Any(), username).Times(1).Return([]db.Note{}, nil)
			},
			checkReturnValues: func(t *testing.T, notes []db.Note, err error) {
				require.Nil(t, err)
			},
		},
		{
			name:     "getting notes from user returns ErrDBInternal",
			username: username,
			mockdbGetNotesFromUser: func(mockdb *mockdb.MockQuerier) {
				mockdb.EXPECT().GetAllNotesFromUser(gomock.Any(), username).Times(1).Return([]db.Note{}, ErrDBInternal)
			},
			checkReturnValues: func(t *testing.T, notes []db.Note, err error) {
				require.ErrorIs(t, err, ErrDBInternal)
			},
		},
	}
	for c := range testCases {
		tc := testCases[c]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockdb := mockdb.NewMockQuerier(ctrl)
			ns := NewService(mockdb)

			tc.mockdbGetNotesFromUser(mockdb)
			notes, err := ns.q.GetAllNotesFromUser(context.Background(), tc.username)
			tc.checkReturnValues(t, notes, err)
		})
	}
}

func TestDeleteNote(t *testing.T) {
	id := uuid.New()

	testCases := []struct {
		name              string
		mockdbDeleteNote  func(mockdb *mockdb.MockQuerier)
		checkReturnValues func(t *testing.T, id uuid.UUID, err error)
	}{
		{
			name: "deleting note OK",
			mockdbDeleteNote: func(mockdb *mockdb.MockQuerier) {
				mockdb.EXPECT().DeleteNote(gomock.Any(), id).Times(1).Return(id, nil)
			},
			checkReturnValues: func(t *testing.T, retID uuid.UUID, err error) {
				require.Equal(t, id, retID)
				require.Nil(t, err)
			},
		},
		{
			name: "deleting note returns ErrDBInternal",
			mockdbDeleteNote: func(mockdb *mockdb.MockQuerier) {
				mockdb.EXPECT().DeleteNote(gomock.Any(), id).Times(1).Return(uuid.Nil, ErrDBInternal)
			},
			checkReturnValues: func(t *testing.T, retID uuid.UUID, err error) {
				require.Equal(t, retID, uuid.Nil)
				require.ErrorIs(t, err, ErrDBInternal)
			},
		},
	}
	for c := range testCases {
		tc := testCases[c]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockdb := mockdb.NewMockQuerier(ctrl)
			ns := NewService(mockdb)

			tc.mockdbDeleteNote(mockdb)
			id, err := ns.q.DeleteNote(context.Background(), id)
			tc.checkReturnValues(t, id, err)
		})
	}
}

func TestUpdateNote(t *testing.T) {
	args := db.UpdateNoteParams{
		ID:    uuid.New(),
		Title: sql.NullString{String: "testtitle1", Valid: true},
		Text:  sql.NullString{String: "testtext", Valid: true},
	}

	testCases := []struct {
		name              string
		mockdbUpdateNote  func(mockdb *mockdb.MockQuerier, args *db.UpdateNoteParams)
		checkReturnValues func(t *testing.T, args *db.UpdateNoteParams, id uuid.UUID, err error)
	}{
		{
			name: "updating note OK",
			mockdbUpdateNote: func(mockdb *mockdb.MockQuerier, args *db.UpdateNoteParams) {
				mockdb.EXPECT().UpdateNote(gomock.Any(), args).Times(1).Return(args.ID, nil)
			},
			checkReturnValues: func(t *testing.T, args *db.UpdateNoteParams, id uuid.UUID, err error) {
				require.Equal(t, args.ID, id)
				require.Nil(t, err)
			},
		},
		{
			name: "updating note returns ErrNotFound",
			mockdbUpdateNote: func(mockdb *mockdb.MockQuerier, args *db.UpdateNoteParams) {
				mockdb.EXPECT().UpdateNote(gomock.Any(), args).Times(1).Return(uuid.Nil, ErrNotFound)
			},
			checkReturnValues: func(t *testing.T, args *db.UpdateNoteParams, id uuid.UUID, err error) {
				require.Equal(t, uuid.Nil, id)
				require.ErrorIs(t, err, ErrNotFound)
			},
		},
		{
			name: "updating note returns ErrDBInternal",
			mockdbUpdateNote: func(mockdb *mockdb.MockQuerier, args *db.UpdateNoteParams) {
				mockdb.EXPECT().UpdateNote(gomock.Any(), args).Times(1).Return(uuid.Nil, ErrDBInternal)
			},
			checkReturnValues: func(t *testing.T, args *db.UpdateNoteParams, id uuid.UUID, err error) {
				require.Equal(t, uuid.Nil, id)
				require.ErrorIs(t, err, ErrDBInternal)
			},
		},
	}
	for c := range testCases {
		tc := testCases[c]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockdb := mockdb.NewMockQuerier(ctrl)
			ns := NewService(mockdb)

			tc.mockdbUpdateNote(mockdb, &args)
			id, err := ns.q.UpdateNote(context.Background(), &args)
			tc.checkReturnValues(t, &args, id, err)
		})
	}
}
