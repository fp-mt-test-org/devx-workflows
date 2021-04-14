package login

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	exec "github.flexport.io/flexport/devx-workflow-scripts/pkg/exec"
)

type ECRSuite struct {
	suite.Suite
	login       L
	execMock    *exec.Mock
	destination string
}

func (suite *ECRSuite) SetupSuite() {
	suite.login = &Obj{}
}

func (suite *ECRSuite) SetupTest() {
	suite.execMock = new(exec.Mock)
	suite.execMock.On("ExecFn", mock.Anything, mock.Anything).Return(nil)
	suite.destination = "391779066913.dkr.ecr.us-east-1.amazonaws.com/dev-tools"
}

func (suite *ECRSuite) TestCorrectlyFormattedDestinationLogsIn() {
	suite.login.ECRLogin(suite.execMock, suite.destination)
	suite.execMock.AssertCalled(suite.T(), "ExecFn", loginCommand+"391779066913.dkr.ecr.us-east-1.amazonaws.com", mock.Anything)
}

func (suite *ECRSuite) TestCorrectlyFormattedDestinationNoError() {
	assert.Nil(suite.T(), suite.login.ECRLogin(suite.execMock, suite.destination), "Unexpected error")
}

func (suite *ECRSuite) TestIncorrectlyFormattedDestinationNoLogin() {
	suite.destination = "391779066913.dkr.ecr.us-east-1.amazonaws.comdev-tools"
	suite.login.ECRLogin(suite.execMock, suite.destination)
	suite.execMock.AssertNotCalled(suite.T(), "ExecFn", mock.Anything)
}

func (suite *ECRSuite) TestIncorrectlyFormattedDestinationReturnsError() {
	suite.destination = "391779066913.dkr.ecr.us-east-1.amazonaws.comdev-tools"
	assert.NotNil(suite.T(), suite.login.ECRLogin(suite.execMock, suite.destination), "No error for incorrect format!")
}

func TestECRSuite(t *testing.T) {
	suite.Run(t, new(ECRSuite))
}
