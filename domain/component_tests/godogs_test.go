package component_tests

import (
	"burskey/mydailylife/domain/package/domain"
	"context"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type TestData struct {
	task            *domain.Task
	tasks           map[string]*domain.Task
	tasksInProgress map[string]*domain.TaskInProgress
}

type ctxKey struct{}

func thatIHaveAStatus(ctx context.Context) (context.Context, error) {
	return ctx, godog.ErrPending
}

func theStatusHasAStatus(ctx context.Context) (context.Context, error) {
	return ctx, godog.ErrPending
}

func theStatusHasATimestamp(ctx context.Context) (context.Context, error) {
	return ctx, godog.ErrPending
}

func theStatusHasAnID(ctx context.Context) (context.Context, error) {
	return ctx, godog.ErrPending
}

func aTaskInProgressNamedIsInStatus(ctx context.Context, arg1, arg2 string) (context.Context, error) {
	return ctx, godog.ErrPending
}

func theTaskNamedIsInStatus(ctx context.Context, arg1, arg2 string) (context.Context, error) {
	return ctx, godog.ErrPending
}
func thatIHaveATask(ctx context.Context, arg1 *godog.Table) (context.Context, error) {
	return ctx, godog.ErrPending
}

func theTaskNamedHasADescription(ctx context.Context, arg1 string) (context.Context, error) {
	_, ok := getTaskHavingNameFromContext(ctx, arg1)
	if !ok {
		return ctx, errors.New("Unable to find a task using description")
	}

	return ctx, nil
}

func getTaskHavingNameFromContext(ctx context.Context, name string) (*domain.Task, bool) {

	if anObject, ok := ctx.Value(ctxKey{}).(*TestData); ok {
		if anObject.task != nil && strings.EqualFold(anObject.task.Name, name) {
			return anObject.task, ok
		}
	}
	return nil, false
}

func theTaskNamedHasAnID(ctx context.Context, arg1 string) (context.Context, error) {

	atask, ok := getTaskHavingNameFromContext(ctx, arg1)
	if !ok {
		return ctx, errors.New("Task not saved in context")
	}
	if atask.ID.String() == "" {
		return ctx, errors.New("GUID not set on task id")
	}
	return ctx, nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a new task$`, aNewTask)
	ctx.Step(`^I complete the task$`, iCompleteTheTask)
	ctx.Step(`^I start the task$`, iStartTheTask)
	ctx.Step(`^that I have a status$`, thatIHaveAStatus)
	ctx.Step(`^the status has a status$`, theStatusHasAStatus)
	ctx.Step(`^the status has a timestamp$`, theStatusHasATimestamp)
	ctx.Step(`^the status has an ID$`, theStatusHasAnID)
	ctx.Step(`^the task status is "([^"]*)"$`, theTaskStatusIs)
	ctx.Step(`^a task in progress named "([^"]*)" is in status "([^"]*)"$`, aTaskInProgressNamedIsInStatus)
	ctx.Step(`^the task named "([^"]*)" is in status "([^"]*)"$`, theTaskNamedIsInStatus)
	ctx.Step(`^the task named "([^"]*)" has a description$`, theTaskNamedHasADescription)
	ctx.Step(`^the task named "([^"]*)" has an ID$`, theTaskNamedHasAnID)

}

func aNewTask(ctx context.Context, data *godog.Table) (context.Context, error) {

	var testData TestData

	if data != nil {
		for j := 1; j < len(data.Rows[0].Cells); j++ {
			name := data.Rows[j].Cells[0].Value
			description := data.Rows[j].Cells[1].Value

			testData.task = domain.NewTask(name, description)

			ctx = context.WithValue(ctx, ctxKey{}, &testData)
		}
	}

	return ctx, nil
}

func iCompleteTheTask(ctx context.Context) (context.Context, error) {
	return ctx, godog.ErrPending
}

func iStartTheTask(ctx context.Context) (context.Context, error) {
	return ctx, godog.ErrPending
}

func theTaskStatusIs(arg1 string, ctx context.Context) (context.Context, error) {
	return ctx, godog.ErrPending
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

// assertExpectedAndActual is a helper function to allow the step function to call
// assertion functions where you want to compare an expected and an actual value.
func assertExpectedAndActual(a expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, expected, actual, msgAndArgs...)
	return t.err
}

type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

// assertActual is a helper function to allow the step function to call
// assertion functions where you want to compare an actual value to a
// predined state like nil, empty or true/false.
func assertActual(a actualAssertion, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, actual, msgAndArgs...)
	return t.err
}

type actualAssertion func(t assert.TestingT, actual interface{}, msgAndArgs ...interface{}) bool

// asserter is used to be able to retrieve the error reported by the called assertion
type asserter struct {
	err error
}

// Errorf is used by the called assertion to report an error
func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}
