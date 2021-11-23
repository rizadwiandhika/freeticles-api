package request

import "github.com/rizadwiandhika/miniproject-backend-alterra/features/articles"

type QueryParams struct {
	Keyword string `query:"q"`
	Today   bool   `query:"today"`
	Limit   int    `query:"limit"`
	Offset  int    `query:"offset"`
}

func ToQueryParamsCore(q *QueryParams) articles.QueryParams {
	return articles.QueryParams{
		Keyword: q.Keyword,
		Today:   q.Today,
		Limit:   q.Limit,
		Offset:  q.Offset,
	}
}
