package controllers
//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/gorilla/mux"
//	"io/ioutil"
//	"net/http"
//	"strconv"
//	"../../models"
//	"../responses"
//)
//
//
//
//func (a *App) AssignTask(w http.ResponseWriter, r *http.Request){
//	var resp = map[string]interface{}{
//		"status": "Success",
//		"Message": "Task created successfully"}
//
//	user := r.Context().Value("userID").(float64)
//	unassignedTask := models.UnAssignedTask{}
//
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		responses.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//
//	err = json.Unmarshal(body, &unassignedTask)
//	if err != nil {
//		responses.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//
//	unassignedTask.Prepare()
//
//	if err = unassignedTask.Validate(); err != nil {
//		responses.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//
//	unassignedTask.UserId = uint(user)
//	taskCreated, err := unassignedTask.Save(a.Db)
//	if err != nil{
//		responses.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//
//	resp["task"] = taskCreated
//	responses.JSON(w, http.StatusCreated, resp)
//
//	return
//}
//
//
//func (a *App) UpdateTask (w http.ResponseWriter, r *http.Request) {
//
//	fmt.Println("coming here")
//	resp := map[string]interface{}{
//		"status": "Success",
//		"message": "Task updated successfully"}
//
//	vars := mux.Vars(r)
//
//	user := r.Context().Value("userID").(float64)
//
//	userID := uint(user)
//
//	id, _ := strconv.Atoi(vars["id"])
//	task, err := models.GetTaskById(id, a.Db)
//
//	if err != nil {
//		responses.JSON(w, http.StatusNotFound, err)
//		return
//	}
//
//	if task.UserId != userID {
//		resp["status"] = "failed"
//		resp["message"] = "you are not allowed to perform this action"
//		responses.JSON(w, http.StatusUnauthorized, resp)
//		return
//	}
//
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		responses.ERROR(w, http.StatusBadRequest, err)
//	}
//
//	taskUpdate := models.Task{}
//	if err = json.Unmarshal(body, &taskUpdate); err != nil {
//		responses.ERROR(w, http.StatusBadRequest, err)
//
//		return
//	}
//
//	taskUpdate.Prepare()
//	_, err = taskUpdate.Update(id, a.Db)
//	if err != nil {
//		responses.ERROR(w, http.StatusInternalServerError, err)
//		return
//	}
//
//	responses.JSON(w, http.StatusOK, resp)
//
//	return
//}
//
//
//func (a *App) GetTasks(w http.ResponseWriter, r *http.Request) {
//	var resp = map[string]interface{}{
//		"status": "success",
//		"message": "tasks"}
//
//	user := r.Context().Value("userID").(float64)
//	userID := int(user)
//	tasks, err := models.TasksOfUser(userID, a.Db)
//
//	if err != nil {
//		responses.ERROR(w, http.StatusInternalServerError, err)
//		return
//	}
//
//	resp["tasks"] = tasks
//	responses.JSON(w, http.StatusOK, resp)
//	return
//}
//
//
//
//func (a *App) DeleteTask(w http.ResponseWriter, r *http.Request) {
//	var resp = map[string]interface{}{
//		"status": "success",
//		"message": "task deleted successfully"}
//
//	vars := mux.Vars(r)
//	user := r.Context().Value("userID").(float64)
//	userID := uint(user)
//
//	id, _ := strconv.Atoi(vars["id"])
//	venue, err := models.GetTaskById(id, a.Db)
//
//	if venue.UserId != userID {
//		resp["status"] = "failed"
//		resp["message"] = "you are not allowed to delete this object"
//		return
//	}
//
//	err = models.DeleteVenue(id, a.Db)
//	if err != nil {
//		responses.ERROR(w, http.StatusInternalServerError, err)
//		return
//	}
//
//	responses.JSON(w, http.StatusOK, resp)
//	return
//}
