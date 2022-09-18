package controllers

import (
	"day2-task1/config"
	database_mock "day2-task1/mocks/database"
	"day2-task1/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

//init echo
func initEcho() *echo.Echo {
	//setup
	config.InitDB()
	e := echo.New()

	return e
}

type MockUserController struct {
	mock.Mock
}

type ControllerTest struct {
	suite.Suite
	userUC       UserUseCase
	mockUserDb   *database_mock.MockUser
	dtoTest      *models.User
	dtoTestFail  *models.User
	dtoTestArray []*models.User
	testLogin    string
}

func (suite *ControllerTest) SetupTest() {
	suite.mockUserDb = new(database_mock.MockUser)
	suite.dtoTest = &models.User{
		Name:     "Postman",
		Email:    "alifipa5@gmail.com",
		Password: "You1",
	}
	suite.userUC = NewUserController(suite.mockUserDb)
	suite.dtoTestFail = nil
	suite.dtoTestArray = append(suite.dtoTestArray, suite.dtoTest)
	suite.testLogin = `{"name":"Jon Snow","email":"jon@labstack.com"}`
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(ControllerTest))
}

func (suite *ControllerTest) TestGetUserController() {

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	//mock
	suite.mockUserDb.Mock.On("GetUsers").Return(suite.dtoTestArray, nil)
	// Assertions
	if assert.NoError(&testing.T{}, suite.userUC.GetUserControllers(c)) {

		assert.Equal(&testing.T{}, http.StatusOK, rec.Code)
		assert.Equal(&testing.T{}, suite.dtoTest, rec.Body.String())
	}

}

func (suite *ControllerTest) TestGetUserByIdController() {

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("4")

	//mock
	suite.mockUserDb.Mock.On("GetUsers", c.ParamNames()).Return(suite.dtoTest, nil)

	// Assertions
	if assert.NoError(&testing.T{}, suite.userUC.GetUserById(c)) {
		assert.Equal(&testing.T{}, http.StatusOK, rec.Code)
		// assert.Equal(&testing.T{}, suite.dtoTest, rec.Body.String())
	}

}

func (suite *ControllerTest) TestLoginController() {

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(suite.testLogin))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	//mock
	suite.mockUserDb.Mock.On("LoginUser", suite.dtoTest).Return(suite.dtoTest, nil)

	// Assertions
	if assert.NoError(&testing.T{}, suite.userUC.LoginController(c)) {
		assert.Equal(&testing.T{}, http.StatusOK, rec.Code)
		assert.Equal(&testing.T{}, suite.dtoTest, rec.Body.String())
	}

}
