package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/models"
	"github.com/alexbednarczyk/crud-gopher/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *models.UserResponse) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByGUID(ctx context.Context, guid string) (*models.UserResponse, error) {
	args := m.Called(ctx, guid)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.UserResponse), args.Error(1)
}

func (m *MockUserRepository) GetAll(ctx context.Context) ([]models.UserResponse, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.UserResponse), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *models.UserResponse) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(ctx context.Context, guid string) error {
	args := m.Called(ctx, guid)
	return args.Error(0)
}

func TestGetAllUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(MockUserRepository)

	users := []models.UserResponse{
		{GUID: "1", DisplayName: "User1"},
		{GUID: "2", DisplayName: "User2"},
	}

	mockRepo.On("GetAll", mock.Anything).Return(users, nil)

	r := gin.New()
	r.GET("/users", GetAllUsers(mockRepo))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
	assert.Equal(t, "User1", response[0].DisplayName)

	mockRepo.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(MockUserRepository)

	newUser := models.UserResponse{
		DisplayName: "NewUser",
		FirstName:   "John",
		LastName:    "Doe",
	}

	mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(u *models.UserResponse) bool {
		return u.DisplayName == "NewUser" && u.FirstName == "John" && u.GUID != ""
	})).Return(nil)

	r := gin.New()
	r.POST("/users", CreateUser(mockRepo))

	body, _ := json.Marshal(newUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetSpecificUser_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(MockUserRepository)

	mockRepo.On("GetByGUID", mock.Anything, "non-existent").Return(nil, repository.ErrNotFound)

	r := gin.New()
	r.GET("/users/:userId", GetSpecificUser(mockRepo))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/non-existent", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockRepo.AssertExpectations(t)
}
