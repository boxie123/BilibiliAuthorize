package app

type AppVersionResponse struct {
	Data []AppVersionData `json:"data"`
}

type AppVersionData struct {
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
