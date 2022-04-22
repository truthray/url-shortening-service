package router

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/truthray/url-shortening-service/internal/app/storage"
)

func TestHandlePost(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		url         string
	}
	testCases := []struct {
		name    string
		request string
		url     string
		want    want
	}{
		{
			name:    "Basic case",
			request: "/",
			url:     "yandex.ru",
			want: want{
				contentType: "application/json",
				statusCode:  http.StatusCreated,
				url:         "http://localhost:8080/0",
			},
		},
		{
			name:    "Empty body case",
			request: "/",
			url:     "",
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusBadRequest,
				url:         "Body is missing\n",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, tc.request, strings.NewReader(tc.url))
			w := httptest.NewRecorder()

			s := storage.New()
			h := New(s)

			h.ServeHTTP(w, request)
			r := w.Result()
			defer r.Body.Close()

			assert.Equal(t, tc.want.statusCode, r.StatusCode)
			assert.Equal(t, tc.want.contentType, r.Header.Get("Content-Type"))

			b, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Body is missing", http.StatusBadRequest)
				return
			}
			url := string(b)

			assert.Equal(t, tc.want.url, url)
		})
	}
}

func TestHandleGet(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		location    string
		response    string
	}
	testCases := []struct {
		name    string
		request string
		want    want
	}{
		{
			name:    "Basic case",
			request: "/0",
			want: want{
				contentType: "application/json",
				statusCode:  http.StatusTemporaryRedirect,
				location:    "yandex.ru",
				response:    "yandex.ru",
			},
		},
		{
			name:    "Not found",
			request: "/2",
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusNotFound,
				location:    "",
				response:    "Id not found\n",
			},
		},
		{
			name:    "Wrong id",
			request: "/abc",
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusBadRequest,
				location:    "",
				response:    "Query parameter is not integer\n",
			},
		},
		{
			name:    "Wrong path",
			request: "/",
			want: want{
				contentType: "",
				statusCode:  http.StatusMethodNotAllowed,
				location:    "",
				response:    "",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := storage.New()
			s.AddURL("yandex.ru")
			h := New(s)
			w := httptest.NewRecorder()

			getRequest := httptest.NewRequest(http.MethodGet, tc.request, nil)
			h.ServeHTTP(w, getRequest)
			r := w.Result()
			defer r.Body.Close()

			assert.Equal(t, tc.want.statusCode, r.StatusCode)
			assert.Equal(t, tc.want.contentType, r.Header.Get("Content-Type"))
			assert.Equal(t, tc.want.location, r.Header.Get("Location"))

			b, err := io.ReadAll(r.Body)
			if err != nil {
				t.Errorf("Body required")
				return
			}
			response := string(b)

			assert.Equal(t, tc.want.response, response)
		})
	}
}
