package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type UnAssignedTask struct {
	gorm.Model
	Title       string    `gorm:"size:100; not null" json:"title"`
	Description string    `gorm:"size:500" json:"description"`
	DueDate     time.Time `gorm:"not null" json:"due_date"`
	IsDone      bool      `gorm:"default: false" json:"is_done"`
	//Creator       User      `gorm:"foreignKey:UserId" json:"-"`
	UserId        uint   `gorm:"not null; OnDelete: CASCADE" json:"user_id"`
	AssigneeEmail string `gorm:"type:varchar(100); not null" json:"assignee_email"`
}

func (t *UnAssignedTask) Prepare() {
	t.Title = strings.TrimSpace(t.Title)
	t.Description = strings.TrimSpace(t.Description)
}

func (t *UnAssignedTask) Validate() error {
	if t.Title == "" {
		return errors.New("title for task is required")
	}

	if t.DueDate.IsZero() {
		return errors.New("dueDare is required")
	}

	if t.AssigneeEmail == "" {
		return errors.New("assignee email is required")
	}

	if err := checkmail.ValidateFormat(t.AssigneeEmail); err != nil {
		return errors.New("email format is incorrect")
	}

	return nil
}

func GetTasksByEmail(email string, db *gorm.DB) (*[]UnAssignedTask, error) {
	tasks := []UnAssignedTask{}
	if err := db.Debug().Table("un_assigned_tasks").Where("assignee_email = ?", email).Find(&tasks).Error; err != nil {
		return &[]UnAssignedTask{}, err
	}

	return &tasks, nil
}

func GenerateTasks(unassignedTasks []UnAssignedTask, userId uint, db *gorm.DB) error {

	//var newTasks []Task
	var UnAssignedIds []int
	for i := range unassignedTasks {
		t := unassignedTasks[i]
		task := Task{
			Title:       t.Title,
			Description: t.Description,
			DueDate:     t.DueDate,
			IsDone:      t.IsDone,
			UserId:      userId,
			Assignor:    t.UserId,
		}
		UnAssignedIds = append(UnAssignedIds, int(t.ID))
		db.Debug().Create(&task)
		//newTasks = append(newTasks, task)
	}
	db.Debug().Delete(&UnAssignedTask{}, UnAssignedIds)
	return nil
}

func (t *UnAssignedTask) Save(db *gorm.DB) (*UnAssignedTask, error) {
	var err error

	err = db.Debug().Create(&t).Error
	if err != nil {
		return &UnAssignedTask{}, err
	}

	return t, nil
}
