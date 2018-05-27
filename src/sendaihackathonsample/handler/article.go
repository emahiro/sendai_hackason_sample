package handler

import (
	"fmt"
	"net/http"

	"render"
	"sendaihackathonsample/model"

)

func Index(w http.ResponseWriter, r *http.Request) {
	sort := r.URL.Query().Get("sort")

	prob := &model.Problems{}
	prop := &model.Proposals{}

	var err error

	switch sort {
	case "createdAt":
		prob, err = prob.GetAllSortByCreatedAt()
		prop, err = prop.GetAllSortByCreatedAt()
	case "like_count":
		prob, err = prob.GetAllSortByLikeCount()
		prop, err = prop.GetAllSortByLikeCount()
	default:
		prob, err = prob.GetAll()
		prop, err = prop.GetAll()
	}

	if err != nil {
		fmt.Printf("[Index] err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	d := struct {
		Problems  *model.Problems
		Proposals *model.Proposals
	}{
		prob,
		prop,
	}

	render.HTMLRender("article/index.tmpl", w, d)
	return
}
