package service

import (
	"stefano/api/models"

	"github.com/sirupsen/logrus"
)

type VideoService interface {
	GetVideos(string) models.Videos
}

type videoService struct {
}

func NewVideoService() VideoService {
	return &videoService{}
}

//NOTE Response is hardcoded for now as requested by Adnan
func (service *videoService) GetVideos(videoType string) models.Videos {
	logrus.Info("GetVideos Service Start")
	videos := models.Videos{}
	if videoType == "standard" {
		standardVideos := []models.Video{
			{Name: "standard1.mp4", Link: "/var/www/videos/standard/"},
			{Name: "standard2.mp4", Link: "/var/www/videos/standard/"},
			{Name: "standard3.mp4", Link: "/var/www/videos/standard/"},
		}

		videos = models.Videos{Videos: standardVideos}
	} else if videoType == "360" {
		videos360 := []models.Video{
			{Name: "video_360_1.mp4", Link: "/var/www/videos/standard/"},
			{Name: "video_360_2.mp4", Link: "/var/www/videos/standard/"},
			{Name: "video_360_3.mp4", Link: "/var/www/videos/standard/"},
		}

		videos = models.Videos{Videos: videos360}
	}
	logrus.Info("videos:", videos)
	return videos
}
