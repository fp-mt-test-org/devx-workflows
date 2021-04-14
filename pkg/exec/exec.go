package exec

import (
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"reflect"

	"github.com/stretchr/testify/mock"
)

// E possesses the functionality to execute bash commands.
type E interface {
	ExecFn(string, ...string) error
}

// Obj implements Exec.
type Obj struct{}

// ExecFn executes the specified command with specified environment variables.
func (e *Obj) ExecFn(command string, env ...string) error {
	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt)

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), env...)
	return cmd.Run()
}

// Mock allows mocking exec attempts.
type Mock struct {
	mock.Mock
}

// ExecFn mock
func (e *Mock) ExecFn(command string, env ...string) error {
	args := e.Called(command, env)
	return args.Error(0)
}

func functionStdoutWrapper(in []reflect.Value) []reflect.Value {
	backup := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	in[0].CallSlice(in[1:]) // Call original function with original params

	w.Close()
	commandOutput, _ := ioutil.ReadAll(r)
	os.Stdout = backup

	var output []reflect.Value
	return append(output, reflect.ValueOf(string(commandOutput)))
}

// Creates new function that returns stdout of passed in function.
func createStdoutWrapper(function interface{}) {
	fn := reflect.ValueOf(function).Elem()
	v := reflect.MakeFunc(reflect.TypeOf(function).Elem(), functionStdoutWrapper)
	fn.Set(v)
}
