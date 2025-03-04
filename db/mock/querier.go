package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	db "shadowsoni666/db/sqlc"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CreateNote mocks base method.
func (m *MockQuerier) CreateNote(arg0 context.Context, arg1 *db.CreateNoteParams) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNote", arg0, arg1)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNote indicates an expected call of CreateNote.
func (mr *MockQuerierMockRecorder) CreateNote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNote", reflect.TypeOf((*MockQuerier)(nil).CreateNote), arg0, arg1)
}

// DeleteNote mocks base method.
func (m *MockQuerier) DeleteNote(arg0 context.Context, arg1 uuid.UUID) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNote", arg0, arg1)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNote indicates an expected call of DeleteNote.
func (mr *MockQuerierMockRecorder) DeleteNote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNote", reflect.TypeOf((*MockQuerier)(nil).DeleteNote), arg0, arg1)
}

// GetAllNotesFromUser mocks base method.
func (m *MockQuerier) GetAllNotesFromUser(arg0 context.Context, arg1 string) ([]db.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllNotesFromUser", arg0, arg1)
	ret0, _ := ret[0].([]db.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllNotesFromUser indicates an expected call of GetAllNotesFromUser.
func (mr *MockQuerierMockRecorder) GetAllNotesFromUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllNotesFromUser", reflect.TypeOf((*MockQuerier)(nil).GetAllNotesFromUser), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockQuerier) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockQuerierMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockQuerier)(nil).GetUser), arg0, arg1)
}

// ListUsers mocks base method.
func (m *MockQuerier) ListUsers(arg0 context.Context) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockQuerierMockRecorder) ListUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockQuerier)(nil).ListUsers), arg0)
}

// RegisterUser mocks base method.
func (m *MockQuerier) RegisterUser(arg0 context.Context, arg1 *db.RegisterUserParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockQuerierMockRecorder) RegisterUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockQuerier)(nil).RegisterUser), arg0, arg1)
}

// UpdateNote mocks base method.
func (m *MockQuerier) UpdateNote(arg0 context.Context, arg1 *db.UpdateNoteParams) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNote", arg0, arg1)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNote indicates an expected call of UpdateNote.
func (mr *MockQuerierMockRecorder) UpdateNote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNote", reflect.TypeOf((*MockQuerier)(nil).UpdateNote), arg0, arg1)
}
