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

type WorkflowSuite struct {
	suite.Suite
	execMock *exec.Mock
}

const testWorkflowName = "testing"
const testWorkflowCommand = "testing123"

func (suite *WorkflowSuite) SetupTest() {
	suite.execMock = new(exec.Mock)
	suite.execMock.On("ExecFn", mock.Anything, mock.Anything).Return(nil)
	viper.Set(fmt.Sprintf("%s.%s.command", workflowKey, testWorkflowName), testWorkflowCommand)
}

func (suite *WorkflowSuite) TestWorkflowCommandIsExecuted() {
	workflowExec(suite.execMock, testWorkflowName)
	suite.execMock.AssertCalled(suite.T(), "ExecFn", testWorkflowCommand, mock.Anything)
}

func (suite *WorkflowSuite) TestWorkflowCommandNoError() {
	assert.Nil(suite.T(), workflowExec(suite.execMock, testWorkflowName), "Unexpected error")
}

func (suite *WorkflowSuite) TestEmptyWorkflowCommandErrors() {
	viper.Set(fmt.Sprintf("%s.%s.command", workflowKey, testWorkflowName), "")
	assert.NotNil(suite.T(), workflowExec(suite.execMock, testWorkflowName), "No error with unspecified build command")
}

func (suite *WorkflowSuite) TestEmptyWorkflowCommandNoworkflowExecution() {
	viper.Set(fmt.Sprintf("%s.%s.command", workflowKey, testWorkflowName), "")
	workflowExec(suite.execMock, testWorkflowName)
	suite.execMock.AssertNotCalled(suite.T(), "ExecFn", mock.Anything)
}

func (suite *WorkflowSuite) TestNoCommandDefWillError() {
	assert.NotNil(suite.T(), workflowExec(suite.execMock, "fake"), "No error when command not defined")
}

func TestWorkflowSuite(t *testing.T) {
	suite.Run(t, new(WorkflowSuite))
}
