Feature: Sunny Day
  This is a sunny day test

  Background:
    Given the following person is defined
      | First | Last |
      | A     | Test |

  Scenario: Tasks have an ID
    Given a new task
      | Name        | Description              |
      | Brush Teeth | Brush my teeth very well |
    Then the task named "Brush Teeth" has an ID


  Scenario: Tasks may have a description
    Given a new task
      | Name        | Description              |
      | Brush Teeth | Brush my teeth very well |
    Then the task named "Brush Teeth" has a description


  Scenario: Tasks have to be selected in order to be worked with
    Given a new task
      | Name        | Description              |
      | Brush Teeth | Brush my teeth very well |
    When I choose to work with a task named "Brush Teeth"
    Then a task in progress for task "Brush Teeth" is in status "Not Started"
    Then the task "Brush Teeth" has 1 tasks in progress

#  Scenario: Only one instance of a task can be in progress by a person
#    Given a new task
#      | Name        | Description              |
#      | Brush Teeth | Brush my teeth very well |
#    When I choose to work with a task named "Brush Teeth"
#    Then a task in progress for task "Brush Teeth" is in status "Not Started"
#    When I choose to work with a task named "Brush Teeth"
#    Then the task "Brush Teeth" has 1 tasks in progress
