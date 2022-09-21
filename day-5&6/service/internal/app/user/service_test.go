package user

import (
	"context"
	"day2-task1/internal/factory"
	database_mock "day2-task1/internal/mocks/database"
	"testing"

	models "day2-task1/internal/model"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

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
	factory      *factory.FakeFactory
}

func (suite *ControllerTest) SetupTest() {
	suite.mockUserDb = new(database_mock.MockUser)
	suite.dtoTest = &models.User{
		Name:     "Postman",
		Email:    "alifipa5@gmail.com",
		Password: "You1",
	}
	suite.factory = &factory.FakeFactory{
		UserRepository: suite.mockUserDb,
	}
	suite.userUC = NewUserController((*factory.Factory)(suite.factory))
	suite.dtoTestFail = nil
	suite.dtoTestArray = append(suite.dtoTestArray, suite.dtoTest)
	suite.testLogin = `{"name":"Jon Snow","email":"jon@labstack.com"}`
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(ControllerTest))
}

func (uc *ControllerTest) TestGetUserController() {

	//mock
	uc.mockUserDb.Mock.On("GetUsers").Return(uc.dtoTestArray, nil)

	//usecase
	result := uc.userUC.GetUserControllers(context.Background())
	uc.Equal(result, uc.dtoTestArray)

}
