package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/axelspringer/swerve/model"

	"github.com/axelspringer/swerve/mocks"
	"github.com/golang/mock/gomock"
)

func TestLogin(t *testing.T) {
	conf := Config{Version: "v1"}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	d := mocks.NewMockDatabaseAdapter(mockCtrl)
	m := model.NewModel(d)
	api := NewAPIServer(m, conf)

	var tests = []struct {
		name               string
		user               string
		mockPlainPass      string
		body               []byte
		handler            http.HandlerFunc
		ExpectedStatusCode int
	}{
		{
			"plain pw matches encypted password",
			"dkern",
			"mytestpw",
			[]byte(`{ "username": "dkern", "pwd": "$2a$12$gh.TtSizoP0JFLHACOdIouPr42713m6k/8fH8jKPl0xQAUBk0OIdS" }`),
			api.login,
			200,
		},
		{
			"plain pw  does not match encypted password",
			"dkern",
			"unknown",
			[]byte(`{ "username": "dkern", "pwd": "$2a$12$gh.TtSizoP0JFLHACOdIouPr42713m6k/8fH8jKPl0xQAUBk0OIdS" }`),
			api.login,
			401,
		},
	}

	for _, test := range tests {
		d.EXPECT().GetPwdHash(test.user).Return(test.mockPlainPass, nil).Times(1)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/login", bytes.NewReader(test.body))

		api.login(w, req)
		resp := w.Result()

		assert.Equal(t, test.ExpectedStatusCode, resp.StatusCode, test.name)
	}
}
