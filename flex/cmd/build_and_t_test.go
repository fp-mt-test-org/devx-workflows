package cmd

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	exec "github.flexport.io/flexport/devx-workflow-scripts/pkg/exec"
)

type BuildAndTestSuite struct {
	suite.Suite
	execMock *exec.Mock
}

func (suite *BuildAndTestSuite) SetupTest() {
	suite.execMock = new(exec.Mock)
	suite.execMock.On("ExecFn", mock.Anything, mock.Anything).Return(nil)
	viper.Set("build-test.command", "build-test")
	viper.Set("build.command", "build")
	viper.Set("test.command", "test")
}

func (suite *BuildAndTestSuite) TestBuildAndTestNoErrorsReturnsWithoutError() {
	buildAndTest(suite.execMock)
	suite.execMock.AssertNumberOfCalls(suite.T(), "ExecFn", 2)
}

func TestBuildAndTestSuite(t *testing.T) {
	suite.Run(t, new(BuildAndTestSuite))
}
