package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/clshu/go-mgm/models"
	"github.com/clshu/go-mgm/utils"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser /usr/create
func CreateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")

	user := &models.User{}
	json.NewDecoder(req.Body).Decode(user)

	// ctx := mgm.Ctx()
	// defer ctx.Done()
	tempPassword, err := utils.CreateTempPassword()
	if err != nil {
		ReturnError(http.StatusInternalServerError, err, "", &res)
		return
	}
	user.TempPassword = tempPassword

	merr := mgm.Coll(user).Create(user)
	if merr != nil {
		ReturnError(http.StatusInternalServerError, merr, "", &res)
	} else {
		viwer := models.UserViewer{ID: user.ID, Email: user.Email, FirstName: user.FirstName, LastName: user.LastName}
		ret := &models.UserAuth{User: viwer, Token: "token"}
		json.NewEncoder(res).Encode(ret)
	}

}

// LogIn log in an user with email and password
func LogIn(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")

	user := &models.User{}
	json.NewDecoder(req.Body).Decode(user)

	ctx := mgm.Ctx()
	defer ctx.Done()

	result := mgm.Coll(user).FindOne(ctx, bson.M{"email": user.Email})
	if result.Err() != nil {
		ReturnError(http.StatusNotFound, result.Err(), "", &res)
		return
	}
	ret := &models.User{}
	err := result.Decode(ret)
	if err != nil {
		ReturnError(http.StatusInternalServerError, err, "", &res)
		return
	}
	if ret.Email != strings.ToLower(user.Email) {
		ReturnError(http.StatusUnauthorized, nil, "Unable to Login", &res)
		return
	}
	perr := utils.ComparePassword(ret.Password, user.Password)
	if perr != nil {
		ReturnError(http.StatusUnauthorized, nil, "Unable to Login", &res)
	} else {
		pret := &struct {
			ID    primitive.ObjectID `json:"id"`
			Email string             `json:"email"`
		}{ID: ret.ID, Email: ret.Email}
		json.NewEncoder(res).Encode(pret)
	}
}
