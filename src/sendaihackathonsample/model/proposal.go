package model

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Proposals struct {
	Items []Proposal `json:"items"`
}

type Proposal struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	LikeCount int64  `json:"like_count"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type ByProposalsCreatedAt Proposals

func (p ByProposalsCreatedAt) Len() int {
	return len(p.Items)
}

func (p ByProposalsCreatedAt) Less(i, j int) bool {
	return p.Items[i].CreatedAt > p.Items[j].CreatedAt
}

func (p ByProposalsCreatedAt) Swap(i, j int) {
	p.Items[i], p.Items[j] = p.Items[j], p.Items[i]
}

func (p *Proposals) GetAll() (*Proposals, error) {
	return p.getAll()
}

func (p *Proposals) GetAllSortByCreatedAt() (*Proposals, error) {
	res, err := p.getAll()
	if err != nil {
		return nil, err
	}
	sort.Sort(ByProposalsCreatedAt(*res))
	return res, nil
}

type ByProposalsLikeCount Proposals

func (p ByProposalsLikeCount) Len() int {
	return len(p.Items)
}

func (p ByProposalsLikeCount) Less(i, j int) bool {
	return p.Items[i].LikeCount > p.Items[j].LikeCount
}

func (p ByProposalsLikeCount) Swap(i, j int) {
	p.Items[i], p.Items[j] = p.Items[j], p.Items[i]
}

func (p Proposals) GetAllSortByLikeCount() (*Proposals, error) {
	res, err := p.getAll()
	if err != nil {
		return nil, err
	}
	sort.Sort(ByProposalsLikeCount(*res))
	return res, nil
}

func (p *Proposals) getAll() (*Proposals, error) {
	f, err := os.Open("./data/proposals.json")
	if err != nil {
		return nil, fmt.Errorf("[model.Proposals.GetAll] Error: %v", err)
	}

	proposals := &Proposals{}
	if err := json.NewDecoder(f).Decode(proposals); err != nil {
		return nil, fmt.Errorf("[model.Proposals.GetAll] Error: %v", err)
	}

	return proposals, nil
}
