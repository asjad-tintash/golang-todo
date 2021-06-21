package controllers

import (
	"../../models"
	"../../utils"
	"../responses"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)


func (a * App) Register(w http.ResponseWriter, r * http.Request) {
	var resp = map[string]interface{}{
		"status": "Success",
		"message": "Registration Successful"}

	user := &models.User{}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//userDb, _ = user.GetUser(a.Db)
	//if userD

	user.Prepare()
	err = user.Validate("")
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userCreated, err := user.Save(a.Db)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["user"] = userCreated
	responses.JSON(w, http.StatusCreated, resp)

	return
}


func (a *App) Login (w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("Content-Type"))
	var resp = map[string]interface{}{
		"status": "Success",
		"message": "Login successfull"}

	user := models.User{}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if len(body) == 0 {
		responses.ERROR(w, http.StatusBadRequest, errors.New("no data received"))
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userDb, err := user.GetUser(a.Db)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if userDb == nil{
		resp["status"] = "failed"
		resp["message"] = "no user found with these credentials"
		responses.JSON(w, http.StatusBadRequest, resp)
	}

	err = models.CheckPasswordHash(user.Password, userDb.Password)
	if err != nil {
		resp["status"] = "failed"
		resp["message"] = "Login failed, please try again"
		responses.JSON(w, http.StatusForbidden, resp)
		return
	}

	token, err := utils.EncodeAuthToken(userDb.ID)

	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["token"] = token
	responses.JSON(w, http.StatusOK, resp)
	return
}