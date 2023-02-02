package model

type AppLifeCycle struct {
	Event   string `json:"event"`
	StoreId string `json:"store"`
}

type AppInstalled struct {
	AppLifeCycle
	AppAccessKey string `json:"accessKey"`
}
