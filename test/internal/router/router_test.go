package router

import (
	"encoding/json"
	"fmt"
	"github.com/FIY-pc/User-management-System/internal/config"
	"github.com/FIY-pc/User-management-System/internal/model"
	"github.com/FIY-pc/User-management-System/internal/router"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type LoginResponse struct {
	Msg   string `json:"msg"`
	Token string `json:"token"`
}

func GetToken() (string, error) {
	// 初始化测试环境
	e := echo.New()
	config.InitConfig()
	model.InitPostgres()
	router.InitRouter(e)

	admin := model.Admin{}
	admin.AdminName = "default"
	admin.AdminPass = "123456"

	req := httptest.NewRequest("GET", fmt.Sprintf("%s/tokens?username=%s&password=%s", config.Config.Server.TestURL, admin.AdminName, admin.AdminPass), nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		return "", fmt.Errorf("login failed")
	}
	// 解析响应
	var response LoginResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		return "", err
	}
	return response.Token, nil
}

func TestCreateAdmin(t *testing.T) {
	// 初始化测试环境
	e := echo.New()
	config.InitConfig()
	model.InitPostgres()
	router.InitRouter(e)
	// 获取token
	token, err := GetToken()
	if err != nil {
		t.Error(err)
	}
	// 设置测试参数
	admin := model.Admin{}
	admin.AdminName = "test1"
	admin.AdminPass = "test1"
	// 定义请求
	url := fmt.Sprintf("%s/%s/admins?username=%s&password=%s", config.Config.Server.TestURL, config.Config.Server.ApiPrefix, admin.AdminName, admin.AdminPass)
	req := httptest.NewRequest("POST", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()
	// 执行请求
	e.ServeHTTP(rec, req)
	// 断言响应
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateAdmin(t *testing.T) {
	// 初始化测试环境
	e := echo.New()
	config.InitConfig()
	model.InitPostgres()
	router.InitRouter(e)
	// 获取token
	token, err := GetToken()
	if err != nil {
		t.Error(err)
	}
	admin := model.Admin{}
	admin.AdminName = "test1"

	url := fmt.Sprintf("%s/%s/admins?username=%s&password=%s", config.Config.Server.TestURL, config.Config.Server.ApiPrefix, admin.AdminName, admin.AdminPass)
	req := httptest.NewRequest("PUT", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetAdmin(t *testing.T) {
	// 初始化测试环境
	e := echo.New()
	config.InitConfig()
	model.InitPostgres()
	router.InitRouter(e)
	// 获取token
	token, err := GetToken()
	if err != nil {
		t.Error(err)
	}

	admin := model.Admin{}
	admin.AdminName = "test1"

	url := fmt.Sprintf("%s/%s/admins?username=%s", config.Config.Server.TestURL, config.Config.Server.ApiPrefix, admin.AdminName)
	req := httptest.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteAdmin(t *testing.T) {
	// 初始化测试环境
	e := echo.New()
	config.InitConfig()
	model.InitPostgres()
	router.InitRouter(e)
	// 获取token
	token, err := GetToken()
	if err != nil {
		t.Error(err)
	}

	admin := model.Admin{}
	admin.AdminName = "test1"

	url := fmt.Sprintf("%s/%s/admins?username=%s", config.Config.Server.TestURL, config.Config.Server.ApiPrefix, admin.AdminName)
	req := httptest.NewRequest("DELETE", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUserCreate(t *testing.T) {
	// 初始化测试环境
	e := echo.New()
	config.InitConfig()
	model.InitPostgres()
	router.InitRouter(e)
	// 获取token
	token, err := GetToken()
	if err != nil {
		t.Error(err)
	}

	user := model.User{}
	user.Name = "test2"
	user.Password = "test2"
	user.Email = "EMAIL"

	url := fmt.Sprintf("%s/%s/users?username=%s&password=%s&email=%s", config.Config.Server.TestURL, config.Config.Server.ApiPrefix, user.Name, user.Password, user.Email)
	req := httptest.NewRequest("POST", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUserUpdate(t *testing.T) {
	// 初始化测试环境
	e := echo.New()
	config.InitConfig()
	model.InitPostgres()
	router.InitRouter(e)
	// 获取token
	token, err := GetToken()
	if err != nil {
		t.Error(err)
	}

	user := model.User{}
	user.Name = "test2"

	url := fmt.Sprintf("%s/%s/users?username=%s&password=%s&email=%s", config.Config.Server.TestURL, config.Config.Server.ApiPrefix, user.Name, user.Password, user.Email)
	req := httptest.NewRequest("PUT", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUserGetByName(t *testing.T) {
	// 初始化测试环境
	e := echo.New()
	config.InitConfig()
	model.InitPostgres()
	router.InitRouter(e)
	// 获取token
	token, err := GetToken()
	if err != nil {
		t.Error(err)
	}

	user := model.User{}
	user.Name = "test2"

	url := fmt.Sprintf("%s/%s/users?username=%s", config.Config.Server.TestURL, config.Config.Server.ApiPrefix, user.Name)
	req := httptest.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUserDeleteByName(t *testing.T) {
	// 初始化测试环境
	e := echo.New()
	config.InitConfig()
	model.InitPostgres()
	router.InitRouter(e)
	// 获取token
	token, err := GetToken()
	if err != nil {
		t.Error(err)
	}

	user := model.User{}
	user.Name = "test2"

	url := fmt.Sprintf("%s/%s/users?username=%s", config.Config.Server.TestURL, config.Config.Server.ApiPrefix, user.Name)
	req := httptest.NewRequest("DELETE", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}
