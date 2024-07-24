package structure

type Link struct {
	LongLink string `json:"longlink"`
	ShortLink string `json:"-"`
}
