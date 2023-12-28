package app

type BilibiliVersionResponse struct {
	Data []BilibiliVersionData `json:"data"`
}

type BilibiliVersionData struct {
	Build   int    `json:"build"`
	Version string `json:"version"`
}

type Device struct {
	BilibiliBuvid string `json:"BilibiliBuvid"`
	AndroidModel  string `json:"AndroidModel"`
	AndroidBuild  string `json:"AndroidBuild"`
	VersionName   string `json:"VersionName"`
	VersionCode   string `json:"VersionCode"`
}
