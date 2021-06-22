package controllers

import (
	"../../models"
	"../../utils"
	"../responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (a *App) AssignTask(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{
		"status":  "Success",
		"Message": "Task created successfully"}

	user := r.Context().Value("userID").(float64)
	unassignedTask := models.UnAssignedTask{}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &unassignedTask)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	unassignedTask.Prepare()

	if err = unassignedTask.Validate(); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Check if the user with this email exists

	userDb, err := models.GetUserByEmail(unassignedTask.AssigneeEmail, a.Db)
	fmt.Println(unassignedTask)

	if userDb != nil {
		// create a task for that user

		task := models.Task{
			Title:       unassignedTask.Title,
			Description: unassignedTask.Description,
			DueDate:     unassignedTask.DueDate,
			IsDone:      unassignedTask.IsDone,
			UserId:      userDb.ID,
			Assignor:    uint(user),
		}

		taskCreated, err := task.Save(a.Db)

		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		resp["task"] = taskCreated
		responses.JSON(w, http.StatusCreated, resp)

	} else {
		unassignedTask.UserId = uint(user)
		unassignedTaskCreated, err := unassignedTask.Save(a.Db)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		err = utils.SendEmail(unassignedTask.AssigneeEmail)
		if err != nil {
			fmt.Println("failed to send email")
		}
		resp["unassigned_task"] = unassignedTaskCreated
		responses.JSON(w, http.StatusCreated, resp)

	}

	return

}
