Feature: Sunny Day
  This is a sunny day test

  Background:


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


  Scenario: Tasks have to selected in order to be worked with
    Given a new task
      | Name        | Description              |
      | Brush Teeth | Brush my teeth very well |
    When I choose to work with a task named "Brush Teeth"
    Then a task in progress named "Brush Teeth" is in status "Not Started"



#can a task have more than one task in progress?