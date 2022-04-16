package data

type ShortLinkRes struct {
	ShortLink string `json:"short_link"`
}

type ShortLinkInfo struct {
	Url       string `json:"url,omitempty"`
	CreatedAt string `json:"created_at"`
	Ttl       int    `json:"ttl,omitempty"`
}
