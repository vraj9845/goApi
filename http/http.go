package http

import (
	"encoding/json"
	"goApi/models"
	"goApi/services"
	"goApi/types"
	"io/ioutil"
	"net/http"
	"strconv"
)

type productHttp struct {
	service services.Product
}

var jsonError = types.ErrInvalidParam{Param: []string{"JSON Request Body"}}

func (p *productHttp) Create(r *http.Request, w http.ResponseWriter) (interface{}, error) {
	// ctx, span := trace.StartSpan(r.Context(), "http.Create")
	// defer span.End()

	// var err error

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, jsonError
	}
	product := models.Product{}
	err = json.Unmarshal(body, &product)
	if err != nil {
		return nil, jsonError
	}

	if product.Id < 1 {
		return nil, types.ErrInvalidParam{Param: []string{"id"}}
	}

	if product.Name == "" {
		return nil, types.ErrInvalidParam{Param: []string{"name"}}
	}

	if product.Price == "" {
		return nil, types.ErrInvalidParam{Param: []string{"price"}}
	}

	err = p.service.Create(product)
	if err != nil {
		return nil, err
	}
	str := "Product successfully added!"
	json.NewEncoder(w).Encode(str)
	return "Product successfully added!", nil
}

func (p *productHttp) Read(w http.ResponseWriter) (interface{}, error) {
	products, err := p.service.Read()
	if err != nil {
		return nil, err
	}
	json.NewEncoder(w).Encode(products)
	return products, nil
}

func (p *productHttp) Update(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var product models.Product

	// id := path.Base(r.URL.Path)
	//id := filepath.Base(r.URL.Path) // id := r.URL.Query().Get("id")
	q := r.URL.Query()

	id := q.Get("id")
	numericId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, jsonError
	}

	err = json.Unmarshal(body, &product)
	if err != nil {
		return nil, jsonError
	}

	if product.Id < 1 {
		return nil, types.ErrInvalidParam{Param: []string{"id"}}
	}

	err = p.service.Update(product.Price, numericId)
	if err != nil {
		return nil, err
	}

	str := "Product updated successfully!"
	json.NewEncoder(w).Encode(str)
	return str, nil
}

func (p *productHttp) Delete(r *http.Request, w http.ResponseWriter) (interface{}, error) {
	// i := path.Base(r.URL.Path)
	// i := r.URL.Query().Get("id")
	q := r.URL.Query()

	i := q.Get("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		return nil, types.ErrInvalidParam{[]string{"id"}}
	}
	if id < 1 {
		return nil, types.ErrInvalidParam{[]string{"id"}}
	}

	err = p.service.Delete(id)
	if err != nil {
		return nil, err
	}
	str := "Successfully deleted!"
	json.NewEncoder(w).Encode(str)
	return "Successfully deleted!", nil
}

func New(service services.Product) productHttp {
	return productHttp{service: service}
}

func (p productHttp) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.Read(w)
	case http.MethodPost:
		p.Create(r, w)
	case http.MethodPut:
		p.Update(w, r)
	case http.MethodDelete:
		p.Delete(r, w)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
