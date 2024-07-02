package hiskio

import (
	"course-downloader/lib"
	"course-downloader/models"
	"fmt"
	"net/http"
)

func (h *Hiskio) getCourses() (models.CoursesResponse, error) {
	url := fmt.Sprintf("%s/v2/me/courses?%s", HISKIO_URL, "search=&course_type=COURSE")
	req := lib.NewHttpRequest(http.MethodGet, url, nil).WithToken(h.token)

	var resp models.CoursesResponse
	if err := req.Run(&resp); err != nil {
		return models.CoursesResponse{}, fmt.Errorf("failed to invoke API courses: %v", err)
	}
	return resp, nil
}

func (h *Hiskio) getChapters(id uint) (models.ChaptersResponse, error) {
	url := fmt.Sprintf("%s/v2/courses/%d/chapters", HISKIO_URL, id)
	req := lib.NewHttpRequest(http.MethodGet, url, nil).WithToken(h.token)
	var resp models.ChaptersResponse
	if err := req.Run(&resp); err != nil {
		return models.ChaptersResponse{}, fmt.Errorf("failed to invoke API chapters: %v", err)
	}
	return resp, nil
}

func (h *Hiskio) getLecture(cid, lid uint) (models.LectureResponse, error) {
	url := fmt.Sprintf("%s/v2/courses/%d/lectures/%d", HISKIO_URL, cid, lid)
	req := lib.NewHttpRequest(http.MethodGet, url, nil).WithToken(h.token)
	var resp models.LectureResponse
	if err := req.Run(&resp); err != nil {
		return models.LectureResponse{}, fmt.Errorf("failed to invoke API lectures: %v", err)
	}
	return resp, nil
}

func printCourses(courses models.CoursesResponse) {
	for _, data := range courses.Data {
		fmt.Printf("%d - %s\n", data.Id, data.Title)
	}
}

func printDataInfo(data models.CoursesDataResponse) {
	fmt.Printf("\nyou choose [%d] %s \n", data.Id, data.Title)
	fmt.Println("feature:")
	for k, f := range data.Feature {
		fmt.Printf("%d - %s \n", k+1, f)
	}
}
