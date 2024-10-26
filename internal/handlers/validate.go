package handlers

import (
	"github.com/kkdai/youtube/v2"
	"html/template"
	"net/http"
	youtubeClient "youtube-get-video/internal/youtube"
)

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	videoURL := r.FormValue("url")
	video, err := youtubeClient.GetVideo(videoURL)
	if err != nil {
		http.Error(w, "Error getting information from video", http.StatusInternalServerError)
		return
	}

	var formats []youtube.Format
	for _, f := range video.Formats {
		formats = append(formats, f)
	}

	data := struct {
		URL       string
		Formats   []youtube.Format
		Thumbnail string
	}{
		URL:       videoURL,
		Formats:   formats,
		Thumbnail: video.Thumbnails[len(video.Thumbnails)-1].URL,
	}

	tmpl := template.Must(template.ParseFiles("web/templates/formats.html"))
	tmpl.Execute(w, data)
}
