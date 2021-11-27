package request

type Request struct {
	ArticleID    uint   `param:"id"`
	Commentar    string `json:"commentar"`
	ReportTypeID uint   `json:"reportTypeId"`
}
