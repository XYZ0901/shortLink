package data

type ShortenReq struct {
	Url string `json:"url,omitempty" bind:"required,url"`
	Ttl int    `json:"ttl,omitempty" bind:"min=0"`
}

type ShortLinkReq struct {
	ShortLink string `json:"short_link" uri:"short_link" bind:"required"`
}
