package cmd

import (
	"fmt"
	"testing"

	exec "devx-workflows/pkg/exec"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CmdSuite struct {
	suite.Suite
	execMock *exec.Mock
}

const testCmdName = "testing"
const testCmdCommand = "testing123"

func (suite *CmdSuite) SetupTest() {
	suite.execMock = new(exec.Mock)
	suite.execMock.On("ExecFn", mock.Anything, mock.Anything).Return(nil)
	viper.Set(fmt.Sprintf("%s.%s.command", cmdKey, testCmdName), testCmdCommand)
}

func (suite *CmdSuite) TestCmdCommandIsExecuted() {
	cmdExec(suite.execMock, testCmdName)
	suite.execMock.AssertCalled(suite.T(), "ExecFn", testCmdCommand, mock.Anything)
}

func (suite *CmdSuite) TestCmdCommandNoError() {
	assert.Nil(suite.T(), cmdExec(suite.execMock, testCmdName), "Unexpected error")
}

func (suite *CmdSuite) TestEmptyCmdCommandErrors() {
	viper.Set(fmt.Sprintf("%s.%s.command", cmdKey, testCmdName), "")
	assert.NotNil(suite.T(), cmdExec(suite.execMock, testCmdName), "No error with unspecified build command")
}

func (suite *CmdSuite) TestEmptyCmdCommandNoCmdExecution() {
	viper.Set(fmt.Sprintf("%s.%s.command", cmdKey, testCmdName), "")
	cmdExec(suite.execMock, testCmdName)
	suite.execMock.AssertNotCalled(suite.T(), "ExecFn", mock.Anything)
}

func (suite *CmdSuite) TestNoCommandDefWillError() {
	assert.NotNil(suite.T(), cmdExec(suite.execMock, "fake"), "No error when command not defined")
}

func TestCmdSuite(t *testing.T) {
	suite.Run(t, new(CmdSuite))
}
