package models

type VideosRequest struct {
	Type      string `form:"type" binding:"required"`
	StartDate string `form:"type"`
	EndDate   string `form:"type"`
}

type Video struct {
	Link string `json:"link"`
	Name string `json:"name"`
}

type Videos struct {
	Videos []Video `json:"videos"`
}
