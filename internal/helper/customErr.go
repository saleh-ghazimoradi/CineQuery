package helper

import (
	"fmt"
	"log/slog"
	"net/http"
)

type CustomErr struct {
	logger *slog.Logger
}

func (c *CustomErr) LogError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)
	c.logger.Error(err.Error(), "method", method, "uri", uri)
}

func (c *CustomErr) ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := Envelope{"error": message}

	if err := WriteJSON(w, status, env, nil); err != nil {
		c.LogError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *CustomErr) ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	c.LogError(r, err)
	message := "the server encountered a problem and could not process your request"
	c.ErrorResponse(w, r, http.StatusInternalServerError, message)
}

func (c *CustomErr) NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	c.ErrorResponse(w, r, http.StatusNotFound, message)
}

func (c *CustomErr) MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	c.ErrorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (c *CustomErr) BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	c.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (c *CustomErr) FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	c.ErrorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func NewCustomErr() *CustomErr {
	return &CustomErr{}
}
