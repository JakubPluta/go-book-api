package book

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"

	e "github.com/JakubPluta/go-book-api/api/resource/common/errors"
	validatorUtil "github.com/JakubPluta/go-book-api/util/validator"
)

type API struct {
	repository *Repository
	validator  *validator.Validate
}

func New(db *gorm.DB, v *validator.Validate) *API {
	return &API{
		repository: NewRepository(db),
		validator:  v,
	}
}

// List godoc
//
//	@summary        List books
//	@description    List books
//	@tags           books
//	@accept         json
//	@produce        json
//	@success        200 {array}     DTO
//	@failure        500 {object}    err.Error
//	@router         /books [get]
func (a *API) List(w http.ResponseWriter, r *http.Request) {
	books, err := a.repository.List()
	if err != nil {
		e.ServerError(w, e.RespDBDataAccessFailure)
		return
	}
	if len(books) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if err := json.NewEncoder(w).Encode(books.ToDto()); err != nil {
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}
}

// Read godoc
//
//	@summary        Read book
//	@description    Read book
//	@tags           books
//	@accept         json
//	@produce        json
//	@param          id	path        string  true    "Book ID"
//	@success        200 {object}    DTO
//	@failure        400 {object}    err.Error
//	@failure        404
//	@failure        500 {object}    err.Error
//	@router         /books/{id} [get]
func (a *API) Read(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.ServerError(w, e.RespInvalidURLParamID)
		return
	}

	book, err := a.repository.Read(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		e.ServerError(w, e.RespDBDataAccessFailure)

		return
	}
	dto := book.ToDto()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}
}

// Create godoc
//
//	@summary        Create book
//	@description    Create book
//	@tags           books
//	@accept         json
//	@produce        json
//	@param          body    body    Form    true    "Book form"
//	@success        201
//	@failure        400 {object}    err.Error
//	@failure        422 {object}    err.Errors
//	@failure        500 {object}    err.Error
//	@router         /books [post]
func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	form := &Form{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		e.ServerError(w, e.RespJSONDecodeFailure)
		return
	}

	if err := a.validator.Struct(form); err != nil {
		fmt.Println(err)
		return
	}

	newBook := form.ToModel()
	newBook.ID = uuid.New()

	_, err := a.repository.Create(newBook)
	if err != nil {
		e.ServerError(w, e.RespDBDataInsertFailure)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Update godoc
//
//	@summary        Update book
//	@description    Update book
//	@tags           books
//	@accept         json
//	@produce        json
//	@param          id      path    string  true    "Book ID"
//	@param          body    body    Form    true    "Book form"
//	@success        200
//	@failure        400 {object}    err.Error
//	@failure        404
//	@failure        422 {object}    err.Errors
//	@failure        500 {object}    err.Error
//	@router         /books/{id} [put]
func (a *API) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	form := &Form{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		e.ServerError(w, e.RespJSONDecodeFailure)
		return
	}

	if err := a.validator.Struct(form); err != nil {
		respBody, err := json.Marshal(validatorUtil.ToErrResponse(err))
		if err != nil {
			e.ServerError(w, e.RespJSONEncodeFailure)
			return
		}
		e.ValidationErrors(w, respBody)
		return
	}

	book := form.ToModel()
	book.ID = id
	rows, err := a.repository.Update(book)
	if err != nil {
		e.ServerError(w, e.RespDBDataUpdateFailure)
		return
	}

	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete godoc
//
//	@summary        Delete book
//	@description    Delete book
//	@tags           books
//	@accept         json
//	@produce        json
//	@param          id  path    string  true    "Book ID"
//	@success        200
//	@failure        400 {object}    err.Error
//	@failure        404
//	@failure        500 {object}    err.Error
//	@router         /books/{id} [delete]
func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.ServerError(w, e.RespInvalidURLParamID)

		return
	}

	rows, err := a.repository.Delete(id)
	if err != nil {
		e.ServerError(w, e.RespDBDataRemoveFailure)
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
