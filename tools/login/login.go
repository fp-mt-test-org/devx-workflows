package login

import (
	"fmt"
	"strings"

	"github.com/stretchr/testify/mock"
	exec "github.flexport.io/flexport/devx-workflow-scripts/pkg/exec"
)

const loginCommand = "aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin "

// L possesses the functionality to authenticate to requires services
type L interface {
	ECRLogin(exec.E, string) error
}

// Obj implements Login
type Obj struct{}

// ECRLogin authenticates to specifed aws account parsed from repo
func (e *Obj) ECRLogin(exec exec.E, destination string) error {
	registry := ""
	index := strings.IndexByte(destination, '/')
	if index > 0 {
		registry = destination[:index]
	} else {
		return fmt.Errorf("ECR destination incorrectly formatted. Example: 484990879900.dkr.ecr.us-east-1.amazonaws.com/fpos-visibility")
	}
	return exec.ExecFn(loginCommand + registry)
}

// Mock allows mocking authentication attempts
type Mock struct {
	mock.Mock
}

// ECRLogin mock
func (e *Mock) ECRLogin(exec exec.E, destination string) error {
	args := e.Called(exec, destination)
	return args.Error(0)
}
