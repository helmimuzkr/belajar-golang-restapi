package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/helmimuzkr/golang-restapi/exception"
	"github.com/helmimuzkr/golang-restapi/model"
	"github.com/helmimuzkr/golang-restapi/presenter"
	"github.com/helmimuzkr/golang-restapi/service"
	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	CreateCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	DeleteCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetAllCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

type categoryController struct {
	service service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &categoryController{service: service}
}

// Create Handler
func (handler *categoryController) CreateCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Decode request body lalu jadikan JSON Data
	categoryRequest := &model.CreateCategoryRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(categoryRequest)
	if err != nil {
		exception.ErrorWebResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Melakukan operasi membuat category
	categoryResponse, err := handler.service.CreateCategory(r.Context(), categoryRequest)
	if err != nil {
		exception.ErrorWebResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Tambahkan content type application/json di header
	w.Header().Add("Content-Type", "application/json")

	// Mempersiapkan response untuk dijadikan response end point, lalu encode dan write ke http
	webResponse := &presenter.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	if err != nil {
		exception.ErrorWebResponse(w, http.StatusInternalServerError, err)
		return
	}
}

// Update Handler
func (handler *categoryController) UpdateCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Decode request body lalu jadikan JSON Data
	categoryRequest := new(model.UpdateCategoryRequest)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(categoryRequest)
	if err != nil {
		exception.ErrorWebResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Mengambil category id di parameter url dan Assign ke categoryRequest
	categoryID, _ := strconv.Atoi(params.ByName("id"))
	categoryRequest.ID = categoryID

	// Tambahkan content type application/json di header
	w.Header().Add("Content-Type", "application/json")

	// Melakukan operasi update category
	categoryResponse, err := handler.service.UpdateCategory(r.Context(), categoryRequest)
	// SEMALEM SAMPE SINI
	// NANTI COMMENTNYA HAPUS
	// if err != nil {
	// 	webResponse := &presenter.WebResponse{
	// 		Code:   http.StatusBadRequest,
	// 		Status: "Bad Request",
	// 		Data:   err.Error(),
	// 	}
	// 	encoder := json.NewEncoder(w)
	// 	err = encoder.Encode(webResponse)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// 	return
	// }
	if err != nil {
		exception.ErrorWebResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Mempersiapkan response untuk dijadikan response end point, lalu encode dan write ke http
	webResponse := &presenter.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	if err != nil {
		exception.ErrorWebResponse(w, http.StatusInternalServerError, err)
		return
	}
}

// Delete Handler
func (handler *categoryController) DeleteCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Mengambil id dari parameter path
	requestID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		exception.ErrorWebResponse(w, http.StatusBadRequest, err)
		return
	}

	// Melakukan proses delete category by id
	err = handler.service.DeleteCategory(r.Context(), requestID)
	if err != nil {
		exception.ErrorWebResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Tambahkan content type application/json di header
	w.Header().Add("Content-Type", "application/json")

	// Mempersiapkan response untuk dijadikan response end point, lalu encode dan write ke http
	webResponse := &presenter.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Berhasil dihapus!",
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	if err != nil {
		exception.ErrorWebResponse(w, http.StatusInternalServerError, err)
		return
	}
}

// Find Category Handler
func (handler *categoryController) GetCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Mengambil id category dari parameter path url http
	requestID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		exception.ErrorWebResponse(w, http.StatusBadRequest, err)
		return
	}

	// Melakukan operasi find category by id
	categoryResponse, err := handler.service.GetCategory(r.Context(), requestID)
	if err != nil {
		exception.ErrorWebResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Menambahkan content type application/json di header
	w.Header().Add("Content-Type", "application/json")

	// Mempersiapkan web response untuk dijadikan sebagai response endpoint
	webResponse := &presenter.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	// Encode dan write ke http data response yang sudah disiapkan
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	if err != nil {
		exception.ErrorWebResponse(w, http.StatusInternalServerError, err)
		return
	}
}

// Find All Category
func (handler *categoryController) GetAllCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Melakukan operasi ambil semua category
	categoriesResponse, err := handler.service.GetAllCategory(r.Context())
	if err != nil {
		exception.ErrorWebResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Menambahkan content type application/json di header
	w.Header().Add("Content-Type", "application/json")

	// Mempersiapkan web response untuk dijadikan sebagai response endpoint
	webResponse := &presenter.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoriesResponse,
	}
	// Encode dan write ke http data response yang sudah disiapkan
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	if err != nil {
		exception.ErrorWebResponse(w, http.StatusInternalServerError, err)
		return
	}
}
