package component_tests

import (
	"burskey/mydailylife/domain/package/domain"
	"context"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type TestData struct {
	tasks           map[string]*domain.Task
	tasksInProgress map[string]*domain.TaskInProgress
	person          *domain.Person
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

		for _, aTask := range anObject.tasks {
			if aTask != nil && strings.EqualFold(aTask.Name, name) {
				return aTask, ok
			}
		}

	}
	return nil, false
}

func getTaskInProgressForTaskFromContext(ctx context.Context, task *domain.Task) (*domain.TaskInProgress, bool) {

	if anObject, ok := ctx.Value(ctxKey{}).(*TestData); ok {

		for _, aTaskInProgress := range anObject.tasksInProgress {
			if aTaskInProgress != nil && strings.EqualFold(aTaskInProgress.TaskID, task.ID) {
				return aTaskInProgress, ok
			}
		}

	}
	return nil, false
}

func getTasksInProgressForTaskFromContext(ctx context.Context, task *domain.Task) ([]*domain.TaskInProgress, bool) {
	var tasksInProgress []*domain.TaskInProgress
	if anObject, ok := ctx.Value(ctxKey{}).(*TestData); ok {

		for _, aTaskInProgress := range anObject.tasksInProgress {
			if aTaskInProgress != nil && strings.EqualFold(aTaskInProgress.TaskID, task.ID) {
				tasksInProgress = append(tasksInProgress, aTaskInProgress)
			}
		}

	}
	return tasksInProgress, (len(tasksInProgress) > 0)
}

func theTaskNamedHasAnID(ctx context.Context, arg1 string) (context.Context, error) {

	atask, ok := getTaskHavingNameFromContext(ctx, arg1)
	if !ok {
		return ctx, errors.New("Task not saved in context")
	}
	if atask.ID == "" {
		return ctx, errors.New("GUID not set on task id")
	}
	return ctx, nil
}

func aTaskInProgressForTaskIsInStatus(ctx context.Context, arg1, arg2 string) (context.Context, error) {

	atask, ok := getTaskHavingNameFromContext(ctx, arg1)
	if !ok {
		return ctx, errors.New("Task not saved in context")
	}

	taskInProgress, ok := getTaskInProgressForTaskFromContext(ctx, atask)
	if !ok {
		return ctx, errors.New("Task in progress not saved in context")
	}
	statusToCheckAgainst := domain.StatusFactory(arg2)
	if taskInProgress.Status.Status != statusToCheckAgainst {
		return ctx, errors.New("Wrong Status is assigned")
	}

	return ctx, nil
}

func theFollowingPersonIsDefined(ctx context.Context, arg1 *godog.Table) (context.Context, error) {

	aRow := arg1.Rows[1]
	first := aRow.Cells[0].Value
	last := aRow.Cells[1].Value

	party := &domain.Person{
		ID:    uuid.New().String(),
		First: first,
		Last:  last,
	}

	if anObject, ok := ctx.Value(ctxKey{}).(*TestData); ok {
		anObject.person = party

	}
	return ctx, nil
}

func theTaskHasTasksInProgress(ctx context.Context, arg1 string, arg2 int) (context.Context, error) {
	atask, ok := getTaskHavingNameFromContext(ctx, arg1)
	if !ok {
		return ctx, errors.New("Task not saved in context")
	}
	var tasksInProgress []*domain.TaskInProgress
	if tasksInProgress, ok = getTasksInProgressForTaskFromContext(ctx, atask); !ok {
		return ctx, errors.New("Task in progress not saved in context")
	}
	if len(tasksInProgress) > arg2 {
		return ctx, errors.New("Too many tasks are in progress")
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
	ctx.Step(`^I choose to work with a task named "([^"]*)"$`, iChooseToWorkWithATaskNamed)

	ctx.Step(`^a task in progress for task "([^"]*)" is in status "([^"]*)"$`, aTaskInProgressForTaskIsInStatus)
	ctx.Step(`^the following person is defined$`, theFollowingPersonIsDefined)
	ctx.Step(`^the task "([^"]*)" has (\d+) tasks in progress$`, theTaskHasTasksInProgress)

}

func aNewTask(ctx context.Context, data *godog.Table) (context.Context, error) {

	var testData TestData

	if data != nil {
		for j := 1; j < len(data.Rows[0].Cells); j++ {
			name := data.Rows[j].Cells[0].Value
			description := data.Rows[j].Cells[1].Value
			if testData.tasks == nil {
				testData.tasks = make(map[string]*domain.Task)
				testData.tasksInProgress = make(map[string]*domain.TaskInProgress)
			}
			task := domain.NewTask(name, description)
			testData.tasks[task.ID] = task

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

func iChooseToWorkWithATaskNamed(ctx context.Context, arg1 string) (context.Context, error) {
	atask, ok := getTaskHavingNameFromContext(ctx, arg1)
	if !ok {
		return ctx, errors.New("Task not saved in context")
	}
	var tip *domain.TaskInProgress
	var err error
	if tip, err = atask.Start(""); err != nil {
		panic("Unable to start work on a task")
	}
	//if err := tip.Start(); err != nil {
	//	panic("Unable to start a task")
	//}

	if anObject, ok := ctx.Value(ctxKey{}).(*TestData); ok {
		anObject.tasksInProgress[tip.ID] = tip
	}
	return ctx, nil
}
