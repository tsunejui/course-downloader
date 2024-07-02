package hiskio

import (
	"course-downloader/lib"
	"course-downloader/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Hiskio) login() error {
	content, err := json.Marshal(models.LoginRequest{
		Account:  h.account,
		Password: h.password,
		Confirm:  true,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal data: %v", err)
	}

	url := fmt.Sprintf("%s/v2/auth/login", HISKIO_URL)
	req := lib.NewHttpRequest(http.MethodPost, url, content)
	var loginResp models.LoginResponse
	if err := req.Run(&loginResp); err != nil {
		return fmt.Errorf("failed to invoke API login: %v", err)
	}
	h.token = loginResp.AccessToken
	return nil
}
