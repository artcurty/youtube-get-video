package youtube

import (
	"github.com/kkdai/youtube/v2"
	"io"
)

var client *youtube.Client

func init() {
	SetClient(&youtube.Client{})
}

func SetClient(c *youtube.Client) {
	client = c
}

func GetVideo(url string) (*youtube.Video, error) {
	return client.GetVideo(url)
}

func GetStream(video *youtube.Video, format *youtube.Format) (io.ReadCloser, int64, error) {
	return client.GetStream(video, format)
}
