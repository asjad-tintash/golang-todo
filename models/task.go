package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type Task struct {
	gorm.Model
	Title       string    `gorm:"size:100; not null" json:"title"`
	Description string    `gorm:"size:500" json:"description"`
	DueDate     time.Time `gorm:"not null" json:"due_date"`
	IsDone      bool      `gorm:"default: false" json:"is_done"`
	Creator     User      `gorm:"foreignKey:UserId" json:"-"`
	UserId      uint      `gorm:"not null,OnDelete: CASCADE" json:"user_id"`
	//Assignee    User      `gorm:"foreignKey: AssigneeId" json:"-"`
	//AssigneeId  uint
}

func GetTaskById(id int, db *gorm.DB) (*Task, error) {
	task := &Task{}
	if err := db.Debug().Table("tasks").Where("id = ?", id).First(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (t *Task) Prepare() {
	t.Title = strings.TrimSpace(t.Title)
	t.Description = strings.TrimSpace(t.Description)
	t.Creator = User{}
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title for task is required")
	}

	if t.DueDate.IsZero() {
		return errors.New("dueDare is required")
	}

	return nil
}

func (t *Task) Save(db *gorm.DB) (*Task, error) {
	var err error

	err = db.Debug().Create(&t).Error
	if err != nil {
		return &Task{}, err
	}

	return t, nil
}

func (t *Task) Update(id int, db *gorm.DB) (*Task, error) {
	if err := db.Debug().Table("tasks").Where("id = ?", id).Updates(Task{
		Title:       t.Title,
		Description: t.Description,
		DueDate:     t.DueDate,
		IsDone:      t.IsDone}).Error; err != nil {
		return &Task{}, err
	}

	return t, nil
}

func TasksOfUser(UserId int, db *gorm.DB) (*[]Task, error) {
	tasks := []Task{}
	if err := db.Debug().Table("tasks").Where("user_id = ?", UserId).Find(&tasks).Error; err != nil {
		return &[]Task{}, err
	}

	return &tasks, nil
}

func DeleteVenue(id int, db *gorm.DB) error {
	if err := db.Debug().Table("tasks").Where("id = ?", id).Delete(&Task{}).Error; err != nil {
		return err
	}

	return nil
}
