package hiskio

import (
	"course-downloader/lib"
	"course-downloader/models"
	"fmt"
	"net/http"
	"strings"
)

// 0: res: 1080
// 1: res: 720
// 2: res: 540
// 3: res: 360
func (h *Hiskio) downloadVideo(source []models.LectureContentSource, path, name string) error {
	if len(source) < 2 {
		return fmt.Errorf("no source found")
	}
	s := source[1]
	if s.Type != "video/mp4" {
		fmt.Println("not video/mp4 file")
	}

	name = strings.ReplaceAll(name, "/", "_")
	fileName := fmt.Sprintf("%s/%s.mp4", path, name)
	req := lib.NewHttpRequest(http.MethodGet, s.Src, nil).WithToken(h.token)
	if err := req.Download(fileName); err != nil {
		return fmt.Errorf("failed to download video: %v", err)
	}
	return nil
}
