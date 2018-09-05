package utils

import (
	"math"
)

type Pagination struct {
	perPage     int64
	totalAmount int64
	currentPage int64
	totalPage   int64
	baseUrl     string

	firstPart  []string
	middlePart []string
	lastPart   []string
}

func NewPagination(totalAmount, perPage, currentPage int64, baseUrl string) *Pagination {
	if currentPage == 0 {
		currentPage = 1
	}

	n := int64(math.Ceil(float64(totalAmount) / float64(perPage)))
	if currentPage > n {
		currentPage = n
	}

	return &Pagination{
		perPage:     perPage,
		totalAmount: totalAmount,
		currentPage: currentPage,
		totalPage:   int64(math.Ceil(float64(totalAmount) / float64(perPage))),
		baseUrl:     baseUrl,
	}
}

func (p *Pagination) TotalPages() int64 {
	return p.totalPage
}

func (p *Pagination) HasPages() bool {
	return p.TotalPages() > 1
}
