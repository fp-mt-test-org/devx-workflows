package cmd

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	exec "devx-workflows/pkg/exec"
)

type BuildSuite struct {
	suite.Suite
	execMock *exec.Mock
}

func (suite *BuildSuite) SetupTest() {
	suite.execMock = new(exec.Mock)
	suite.execMock.On("ExecFn", mock.Anything, mock.Anything).Return(nil)
	viper.Set("build.command", "build my-app")
}

func (suite *BuildSuite) TestBuildCommandIsExecuted() {
	build(suite.execMock)
	suite.execMock.AssertCalled(suite.T(), "ExecFn", "build my-app", mock.Anything)
}

func (suite *BuildSuite) TestBuildCommandNoError() {
	assert.Nil(suite.T(), build(suite.execMock), "Unexpected error")
}

func (suite *BuildSuite) TestEmptyBuildCommandErrors() {
	viper.Set("build.command", "")
	assert.NotNil(suite.T(), build(suite.execMock), "No error with unspecified build command")
}

func (suite *BuildSuite) TestEmptyBuildCommandNoBuildExecution() {
	viper.Set("build.command", "")
	suite.execMock.AssertNotCalled(suite.T(), "ExecFn", mock.Anything)
}

func TestBuildSuite(t *testing.T) {
	suite.Run(t, new(BuildSuite))
}
