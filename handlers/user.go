package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"notes/models"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

/*

/api/user/signup POST

body:
{
	"username": "...",
	"password": "..."
}

*/

func (h *Handlers) UserSignupHandler(w http.ResponseWriter, r *http.Request) {
	newUserReq := new(models.UserRequest)
	if err := json.NewDecoder(r.Body).Decode(newUserReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(newUserReq.Username) < 3 {
		http.Error(w, "Username should be 3 symbols or longer", http.StatusBadRequest)
		return
	} else if len(newUserReq.Username) > 50 {
		http.Error(w, "Username is too long", http.StatusBadRequest)
		return
	}

	if len(newUserReq.Password) < 6 {
		http.Error(w, "Password should be 6 symbols or longer", http.StatusBadRequest)
		return
	} else if len(newUserReq.Password) > 50 {
		http.Error(w, "Password is too long", http.StatusBadRequest)
		return
	}

	if err := h.db.CreateUser(newUserReq); err != nil {
		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "23505":
				http.Error(w, "Username already exists", http.StatusBadRequest)
				return
			}
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	respondWithJSON(w, http.StatusCreated, "Created")
}

/*

/api/user/login POST

body:
{
	"username": "...",
	"password": "..."
}

returns User.ID, User.Username

*/

func (h *Handlers) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	loginUserReq := new(models.UserRequest)
	if err := json.NewDecoder(r.Body).Decode(loginUserReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.db.GetUserByUsername(loginUserReq.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if loginUserReq.Password != user.Password {
		http.Error(w, "Password doesn't match", http.StatusBadRequest)
		return
	}

	tokenString, err := h.jwt.CreateJWT(user)
	if err != nil {
		log.Println("Error when trying to create JWT token:", err)
		return
	}
	//cache the token
	log.Println(tokenString)

	respondWithJSON(w, http.StatusOK, map[string]any{
		"id":       user.ID,
		"username": user.Username,
	})
}

/*

/api/user/{userId} GET

returns User.ID, User.Username

*/

func (h *Handlers) UserGetHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(mux.Vars(r)["userId"])
	if err != nil {
		http.Error(w, "userId is not a number", http.StatusBadRequest)
		return
	}

	user, err := h.db.GetUser(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]any{
		"id":       user.ID,
		"username": user.Username,
	})
}

/*

/api/user/delete/{userId} DELETE

*/

func (h *Handlers) UserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(mux.Vars(r)["userId"])
	if err != nil {
		http.Error(w, "userId is not a number", http.StatusBadRequest)
		return
	}

	if err := h.db.DeleteUser(userId); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	respondWithJSON(w, http.StatusOK, "Deleted")
}

/*

/api/user/logout/{userId} POST

*/

func (h *Handlers) UserLogoutHandler(w http.ResponseWriter, r *http.Request) {
	// invalidate jwt token
	fmt.Println("UserLogoutHandler")
}
