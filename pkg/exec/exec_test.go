package exec

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StdoutSuite struct {
	suite.Suite
	GetExecStdout func(func(string, ...string) error, string, ...string) string
	exec          E
}

func (suite *StdoutSuite) SetupSuite() {
	createStdoutWrapper(&suite.GetExecStdout)
	suite.exec = &Obj{}
}

func (suite *StdoutSuite) TestBashCommandExecutes() {
	command := "echo testing123"
	commandOutput := suite.GetExecStdout(suite.exec.ExecFn, command)
	assert.Equal(suite.T(), "testing123\n", commandOutput)
}

func (suite *StdoutSuite) TestBashCommandWithEnvVarsExecutes() {
	command := "echo $ENV_VAR1 $ENV_VAR2"
	env := []string{"ENV_VAR1=testing123", "ENV_VAR2=testing456"}
	commandOutput := suite.GetExecStdout(suite.exec.ExecFn, command, env...)
	assert.Equal(suite.T(), "testing123 testing456\n", commandOutput)
}

func TestStdoutSuite(t *testing.T) {
	suite.Run(t, new(StdoutSuite))
}

type ErrorSuite struct {
	suite.Suite
	exec E
}

func (suite *ErrorSuite) SetupSuite() {
	suite.exec = &Obj{}
}

func (suite *ErrorSuite) TestFakeCommandReturnsError() {
	err := suite.exec.ExecFn("fakecommand")
	assert.NotNil(suite.T(), err, "Didn't return an error when running a fake command")
}

func (suite *ErrorSuite) TestBashCommandReturnsNoError() {
	err := suite.exec.ExecFn("echo testing123")
	assert.Nil(suite.T(), err, "Returned error for real command")
}

func TestErrorSuite(t *testing.T) {
	suite.Run(t, new(ErrorSuite))
}
