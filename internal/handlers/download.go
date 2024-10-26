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
		http.Error(w, "Erro ao obter informações do vídeo", http.StatusInternalServerError)
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
		http.Error(w, "Formato não encontrado", http.StatusInternalServerError)
		return
	}

	stream, _, err := youtubeClient.GetStream(video, format)
	if err != nil {
		http.Error(w, "Erro ao obter o stream do vídeo", http.StatusInternalServerError)
		return
	}

	fileName := fmt.Sprintf("%s.mp4", video.Title)
	w.Header().Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")

	_, err = io.Copy(w, stream)
	if err != nil {
		http.Error(w, "Erro ao enviar o vídeo", http.StatusInternalServerError)
		return
	}
}
