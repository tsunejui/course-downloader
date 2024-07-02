package hiskio

import (
	"course-downloader/config"
	"course-downloader/models"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	HISKIO_URL = "https://api.hiskio.com"
	timeFormat = "20060102150405"
)

type Hiskio struct {
	account  string
	password string
	token    string
}

func New(conf config.Auth) (*Hiskio, error) {
	h := &Hiskio{
		account:  conf.Account,
		password: conf.Password,
	}
	if err := h.login(); err != nil {
		return nil, fmt.Errorf("failed to new struct: %v", err)
	}
	return h, nil
}

func (h *Hiskio) Download(path string) error {

	tPath := fmt.Sprintf("%s/%s", path, time.Now().Format(timeFormat))
	if err := os.Mkdir(tPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create root directory: %v", err)
	}

	// print courses
	completedCourses, err := h.getCourses(true)
	if err != nil {
		return fmt.Errorf("[completed] failed to print the courses: %v", err)
	}
	uncompletedCourses, err := h.getCourses(false)
	if err != nil {
		return fmt.Errorf("[uncompleted] failed to print the courses: %v", err)
	}

	for {
		courses := printCourses(completedCourses, uncompletedCourses)

		// scanf course id
		var courseId uint
		fmt.Printf("\n(type 0 to interrupt the session)\n")
		fmt.Printf("Please type the course id that you want to download: ")
		fmt.Scanf("%d", &courseId)
		if courseId == 0 {
			break
		}

		// print data
		data, err := findCourse(courses, courseId)
		if err != nil {
			return fmt.Errorf("course not found: %v", err)
		}
		printDataInfo(data)
		dTitle := strings.ReplaceAll(data.Title, "/", "_")
		dPath := fmt.Sprintf("%s/%s", tPath, dTitle)
		if err := os.Mkdir(dPath, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create data directory: %v", err)
		}

		// download
		fmt.Printf("\n\nstart to download...\n")
		chapters, err := h.getChapters(courseId)
		if err != nil {
			return fmt.Errorf("failed to get chapters: %v", err)
		}

		for _, c := range chapters.Chapters {
			fmt.Printf("---- %s ---\n", c.Title)
			title := strings.ReplaceAll(c.Title, "/", "_")
			cPath := fmt.Sprintf("%s/%s", dPath, title)
			if err := os.Mkdir(cPath, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create chapter directory: %v", err)
			}
			h.download(c.Lectures, cPath)
		}
	}
	return nil
}

func (h *Hiskio) download(lectures []models.LecturesResponse, path string) error {
	var wg sync.WaitGroup
	for _, l := range lectures {
		wg.Add(1)
		title := strings.ReplaceAll(l.Title, "/", "_")
		fmt.Printf("download: %d - %s\n", l.Id, title)
		lecture, err := h.getLecture(l.CourseId, l.Id)
		if err != nil {
			return fmt.Errorf("failed to get Lecture %d - %s: %v", l.Id, l.Title, err)
		}
		sources := lecture.Content.Sources
		go func(lsp models.LecturesResponse) {
			if err := h.downloadVideo(sources, path, lsp.Title); err != nil {
				log.Println(err)
			}
			wg.Done()
		}(l)
		time.Sleep(500 * time.Millisecond)
	}
	wg.Wait()
	return nil
}
