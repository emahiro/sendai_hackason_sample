package model

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Problems struct {
	Items []Problem `json:"items"`
}

type Problem struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	LikeCount int64  `json:"like_count"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type sortByCreatedAt Problems

func (p sortByCreatedAt) Len() int {
	return len(p.Items)
}

func (p sortByCreatedAt) Less(i, j int) bool {
	return p.Items[i].CreatedAt > p.Items[j].CreatedAt
}

func (p sortByCreatedAt) Swap(i, j int) {
	p.Items[i], p.Items[j] = p.Items[j], p.Items[i]
}

func (p *Problems) GetAll() (*Problems, error) {
	return p.getAll()
}

func (p *Problems) GetAllSortByCreatedAt() (*Problems, error) {
	prob := &Problems{}
	res, err := prob.getAll()
	if err != nil {
		return nil, err
	}

	sort.Sort(sortByCreatedAt(*res))
	return res, nil
}

type sortByLikeCount Problems

func (p sortByLikeCount) Len() int {
	return len(p.Items)
}

func (p sortByLikeCount) Less(i, j int) bool {
	return p.Items[i].LikeCount > p.Items[j].LikeCount
}

func (p sortByLikeCount) Swap(i, j int) {
	p.Items[i], p.Items[j] = p.Items[j], p.Items[i]
}

func (p *Problems)GetAllSortByLikeCount()(*Problems, error){
	prob := &Problems{}
	res, err := prob.getAll()
	if err != nil {
		return nil, err
	}
	sort.Sort(sortByLikeCount(*res))
	return res, nil
}

func (p *Problems) getAll() (*Problems, error) {
	f, err := os.Open("./data/problems.json")
	if err != nil {
		return nil, fmt.Errorf("[model.Problems.GetAll] Error: %v", err)

	}

	problems := &Problems{}
	if err := json.NewDecoder(f).Decode(problems); err != nil {
		return nil, fmt.Errorf("[ERROR] %v", err)
	}

	return problems, nil
}
