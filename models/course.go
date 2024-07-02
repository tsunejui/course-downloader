package models

import "fmt"

type CoursesDataResponse struct {
	Id      uint     `json:"id"`
	Title   string   `json:"title"`
	Feature []string `json:"feature"`
}

type CoursesResponse struct {
	Data []CoursesDataResponse `json:"data"`
}

type LectureContentSource struct {
	Res  int    `json:"res"`
	Size int    `json:"size"`
	Src  string `json:"src"`
	Type string `json:"type"`
}

type LectureContent struct {
	Sources []LectureContentSource `json:"sources"`
}

type LectureResponse struct {
	Title   string         `json:"title"`
	Content LectureContent `json:"content"`
}

type LecturesResponse struct {
	Id       uint   `json:"id"`
	Title    string `json:"title"`
	CourseId uint   `json:"course_id"`
}

type ChapterResponse struct {
	Id       uint               `json:"id"`
	Lectures []LecturesResponse `json:"lectures"`
	Title    string             `json:"title"`
}

type ChaptersResponse struct {
	Title    string            `json:"title"`
	Chapters []ChapterResponse `json:"chapters"`
}

func (resp CoursesResponse) FindData(id uint) (CoursesDataResponse, error) {
	for _, d := range resp.Data {
		if d.Id == id {
			return d, nil
		}
	}
	return CoursesDataResponse{}, fmt.Errorf("failed to find the course: %d", id)
}
