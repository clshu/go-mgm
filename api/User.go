package api

import (
	"encoding/json"
	"net/http"

	"github.com/clshu/go-mgm/models"
	"github.com/kamva/mgm/v3"
)

// CreateUser /usr/create
func CreateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")

	user := &models.User{}
	json.NewDecoder(req.Body).Decode(user)

	ctx := mgm.Ctx()
	defer ctx.Done()

	err := mgm.Coll(user).Create(user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `"}`))
	} else {
		json.NewEncoder(res).Encode(user)
	}

}
