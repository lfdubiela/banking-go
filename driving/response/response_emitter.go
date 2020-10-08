package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	w http.ResponseWriter
}

type Emitter interface {
	Json()
	Error()
	BadRequest()
	Created()
}

func NewResponseEmitter(w http.ResponseWriter) Response {
	return Response{w}
}

func (e Response) Json(code int, p interface{}) {
	e.w.Header().Set("Content-Type", "application/json")
	e.w.WriteHeader(code)

	if p != nil {
		e.w.Write(e.encode(p))
	}
}

func (e Response) BadRequest(p interface{}) {
	e.Json(http.StatusBadRequest, p)
}

func (e Response) InternalServerError(p interface{}) {
	e.Json(http.StatusInternalServerError, p)
}

func (e Response) UnprocessableEntity(p interface{}) {
	e.Json(http.StatusUnprocessableEntity, p)
}

func (e Response) NotFound(p interface{}) {
	e.Json(http.StatusNotFound, p)
}

func (e Response) Created(p interface{}) {
	e.Json(http.StatusCreated, p)
}

func (e Response) OK(p interface{}) {
	e.Json(http.StatusOK, p)
}

func (e Response) Conflict(p interface{}) {
	e.Json(http.StatusConflict, p)
}

func (e Response) encode(p interface{}) []byte {
	res, _ := json.Marshal(p)
	return res
}
