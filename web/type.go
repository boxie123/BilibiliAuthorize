package web

type NavResp struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	TTL     int     `json:"ttl"`
	Data    NavData `json:"data"`
}

type NavWbiImg struct {
	ImgURL string `json:"img_url"`
	SubURL string `json:"sub_url"`
}

type NavData struct {
	IsLogin bool      `json:"isLogin"`
	WbiImg  NavWbiImg `json:"wbi_img"`
}
