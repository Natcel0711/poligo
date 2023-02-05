package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

type UserController struct {
	storage *UserStorage
}

func NewUserController(storage *UserStorage) *UserController {
	return &UserController{
		storage: storage,
	}
}

type createUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type createUserResponse struct {
	ID string `json:"id"`
}

func (c *createUserRequest) Bind(r *http.Request) error {
	return nil
}

func (t *UserController) create(w http.ResponseWriter, r *http.Request) {
	data := &createUserRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	id, err := t.storage.createUser(data.Email, data.Password, data.Username, context.Background())
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	j, err := json.Marshal(createUserResponse{
		ID: id,
	})
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	w.Write(j)
	return
}

func (t *UserController) getAll(w http.ResponseWriter, r *http.Request) {
	users, err := t.storage.getAllUsers(context.Background())
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("error: %v", err)))
		return
	}
	j, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("error: %v", err)))
		return
	}
	w.Write([]byte(j))
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}
