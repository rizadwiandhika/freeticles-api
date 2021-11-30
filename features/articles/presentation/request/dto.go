package request

import (
	"github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"
)

type QueryParams struct {
	Keyword string `query:"q"`
	Today   bool   `query:"today"`
	Limit   int    `query:"limit"`
	Offset  int    `query:"offset"`
}

type Article struct {
	ID       uint   `form:"id"`
	Title    string `form:"title"`
	Subtitle string `form:"subtitle"`
	Content  string `form:"content"`
	Tags     string `form:"tags"`
}

func ToQueryParamsCore(q *QueryParams) articles.QueryParams {
	return articles.QueryParams{
		Keyword: q.Keyword,
		Today:   q.Today,
		Limit:   q.Limit,
		Offset:  q.Offset,
	}
}
