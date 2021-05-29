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
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}

	user.TempPassword = tempPassword

	merr := mgm.Coll(user).Create(user)
	if merr != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `"}`))
	} else {
		ret := &struct {
			ID    primitive.ObjectID `json:"id"`
			Email string             `json:"email"`
		}{ID: user.ID, Email: user.Email}
		// fmt.Printf("%v %T\n", user.ID, user.ID)
		// fmt.Printf("%v %T\n", user.Email, user.Email)
		// fmt.Printf("%v %T\n", ret, ret)
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
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + result.Err().Error() + `"}`))
		return
	}
	ret := &models.User{}
	err := result.Decode(ret)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	if ret.Email != strings.ToLower(user.Email) {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message":  Unable To Login"}`))
		return
	}
	perr := utils.ComparePassword(ret.Password, user.Password)
	if perr != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `"}`))
	} else {
		pret := &struct {
			ID    primitive.ObjectID `json:"id"`
			Email string             `json:"email"`
		}{ID: ret.ID, Email: ret.Email}
		// fmt.Printf("%v %T\n", user.ID, user.ID)
		// fmt.Printf("%v %T\n", user.Email, user.Email)
		// fmt.Printf("%v %T\n", ret, ret)
		json.NewEncoder(res).Encode(pret)
	}
}
