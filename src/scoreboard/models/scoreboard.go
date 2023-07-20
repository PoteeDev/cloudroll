package models

import (
	"gorm.io/gorm"
)

type Score struct {
	gorm.Model
	Value  int64
	TeamID int
}

type TaskColumn struct {
	gorm.Model
	ColumnName string
	Card       []*Card
}

type Board struct {
	gorm.Model
	TeamID  int
	Columns []*Column
}

type Column struct {
	gorm.Model
	BoardID int
	Name    string
	Cards   []*Card
}

type Card struct {
	gorm.Model
	ColumnID int
	Title    string
	//Labels   []*Label `gorm:"many2many:task_label;"`
	Tasks []*TeamTask
}

type TeamTask struct {
	gorm.Model
	CardID    int
	TaskID    int
	Task      *Task
	Completed bool
}

type Task struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex"`
	Description string
	Points      int64
}

type Label struct {
	gorm.Model
	Name  string
	Color string
}

type User struct {
	gorm.Model
	UserID string `gorm:"uniqueIndex"`
	TeamID int
}

type Team struct {
	gorm.Model
	Name  string `gorm:"uniqueIndex"`
	Users []User
	Score Score
	Board Board
}
