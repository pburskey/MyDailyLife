package domain

import (
	"testing"
)

func TestTask(t *testing.T) {

	task := NewTask("a name", "a description")

	if task.Description == "" {
		t.Fatal("Expecting 'Simple'")
	}

}
