package handlers

import (
	"fmt"
	"github.com/kkdai/youtube/v2"
	"io"
	"net/http"
	youtubeClient "youtube-get-video/internal/youtube"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	videoURL := r.FormValue("url")
	formatItag := r.FormValue("format")
	includeAudio := r.FormValue("audio-"+formatItag) == "true"
	video, err := youtubeClient.GetVideo(videoURL)
	if err != nil {
		http.Error(w, "Error getting information from video", http.StatusInternalServerError)
		return
	}

	var format *youtube.Format
	for _, f := range video.Formats {
		if fmt.Sprintf("%d", f.ItagNo) == formatItag {
			if includeAudio && f.AudioChannels > 0 {
				format = &f
				break
			} else if !includeAudio && f.AudioChannels == 0 {
				format = &f
				break
			}
		}
	}

	if format == nil {
		http.Error(w, "Format not found", http.StatusInternalServerError)
		return
	}

	stream, _, err := youtubeClient.GetStream(video, format)
	if err != nil {
		http.Error(w, "Error getting video stream", http.StatusInternalServerError)
		return
	}

	fileName := fmt.Sprintf("%s.mp4", video.Title)
	w.Header().Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")

	_, err = io.Copy(w, stream)
	if err != nil {
		http.Error(w, "Error uploading video", http.StatusInternalServerError)
		return
	}
}
