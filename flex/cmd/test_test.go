package cmd

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	exec "github.flexport.io/flexport/devx-workflow-scripts/pkg/exec"
)

type TestSuite struct {
	suite.Suite
	execMock *exec.Mock
}

func (suite *TestSuite) SetupTest() {
	suite.execMock = new(exec.Mock)
	suite.execMock.On("ExecFn", mock.Anything, mock.Anything).Return(nil)
	viper.Set("test.command", "test")
}

func (suite *TestSuite) TestTestCommandIsConfiguredCommandIsExecuted() {
	test(suite.execMock)
	suite.execMock.AssertCalled(suite.T(), "ExecFn", "test", mock.Anything)
}

func (suite *TestSuite) TestTestCommandIsNotConfiguredCommandIsNotExecuted() {
	viper.Set("test.command", "")
	test(suite.execMock)
	suite.execMock.AssertNotCalled(suite.T(), "ExecFn", mock.Anything)
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
