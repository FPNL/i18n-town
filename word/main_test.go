package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/FPNL/i18n-town/src/lib/idatabase"
	"github.com/FPNL/i18n-town/src/lib/igin"
	"github.com/FPNL/i18n-town/src/tools"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// global router
var gRouter *gin.Engine

func init() {
	fmt.Println("Init...")
	{ // #DB
		err := idatabase.Go()
		if err != nil {
			panic("init problem: " + err.Error())
		}
		gRouter = igin.SetupRouter()
	}
	{ // # Data
		w := tools.ServeHTTP(gRouter, http.MethodDelete, "/word/deleteAll", nil)
		if w.Code != http.StatusOK {
			panic("init problem: " + w.Body.String())
		}
		reqBody := tools.MakeJsonBody([]map[string]string{
			{
				"tag":  "hello",
				"lang": "eng",
				"word": "hello",
			},
			{

				"tag":  "hello2",
				"lang": "eng",
				"word": "hello2",
			},
		})
		w = tools.ServeHTTP(gRouter, http.MethodPost, "/word/addMany", reqBody)
		if w.Code != http.StatusOK {
			panic("init problem: " + w.Body.String())
		}
	}
}

func TestPingPong(t *testing.T) {
	w := tools.ServeHTTP(gRouter, http.MethodGet, "/ping", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestFetchEmptyWords(t *testing.T) {
	w := tools.ServeHTTP(gRouter, http.MethodGet, "/word/all", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `[{"Id":1,"Tag":"hello","Lang":"eng","Word":"hello"},{"Id":2,"Tag":"hello2","Lang":"eng","Word":"hello2"}]`, w.Body.String())
}

func Test_Group_AddOneWord(t *testing.T) {
	tools.RunGroup(t, tools.TestGroup{
		"AddOneWord": func(t *testing.T) {
			w := tools.ServeHTTP(gRouter, http.MethodPost, "/word/addOne?tag=hello3&lang=eng&word=hello", nil)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "ok", w.Body.String())
		},
		"Err_BadQuery": func(t *testing.T) {
			w := tools.ServeHTTP(gRouter, http.MethodPost, "/word/addOne", nil)
			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, "missing query", w.Body.String())
		},
		"Err_Duplicate": func(t *testing.T) {
			w := tools.ServeHTTP(gRouter, http.MethodPost, "/word/addOne?tag=hello&lang=eng&word=hello", nil)
			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, "word duplicate", w.Body.String())
		},
	})
}

func Test_Group_AddManyWords(t *testing.T) {
	tools.RunGroup(t, tools.TestGroup{
		"AddManyWords": func(t *testing.T) {
			body := tools.MakeJsonBody([]map[string]string{
				{
					"word": "go",
					"lang": "eng",
					"tag":  "go",
				},
				{
					"word": "go",
					"lang": "chi",
					"tag":  "go",
				},
			})
			w := tools.ServeHTTP(gRouter, http.MethodPost, "/word/addMany", body)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "ok", w.Body.String())
		},
		"Err_Duplicate": func(t *testing.T) {
			body := tools.MakeJsonBody([]map[string]string{
				{
					"word": "go",
					"lang": "eng",
					"tag":  "hello",
				},
				{
					"word": "go",
					"lang": "chi",
					"tag":  "go",
				},
			})
			w := tools.ServeHTTP(gRouter, http.MethodPost, "/word/addMany", body)
			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, "word duplicate", w.Body.String())
		},
		"Err_BadQuery": func(t *testing.T) {
			body := tools.MakeJsonBody([]map[string]string{
				{},
				{
					"word": "go",
					"lang": "chi",
					"tag":  "go",
				},
			})
			w := tools.ServeHTTP(gRouter, http.MethodPost, "/word/addMany", body)
			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, "missing query", w.Body.String())
		},
	})
}

func Test_Group_UpdateOneWord(t *testing.T) {
	tools.RunGroup(t, tools.TestGroup{
		"UpdateOneWord": func(t *testing.T) {
			w := tools.ServeHTTP(gRouter, http.MethodGet, "/word/all", nil)
			fmt.Println(w.Body.String())
			w = tools.ServeHTTP(gRouter, http.MethodPut, "/word/updateOne?id=1&word=go", nil)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "ok", w.Body.String())
		},
		"Err_WordNotExist": func(t *testing.T) {
			w := tools.ServeHTTP(gRouter, http.MethodPut, "/word/updateOne?id=0&word=hello", nil)
			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, "word not exits", w.Body.String())
		},
		"Err_BadQuery": func(t *testing.T) {
			w := tools.ServeHTTP(gRouter, http.MethodPut, "/word/updateOne?id=hello", nil)
			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, `strconv.Atoi: parsing "hello": invalid syntax`, w.Body.String())
		},
	})
}

func Test_Group_UpdateManyWords(t *testing.T) {
	tools.RunGroup(t, tools.TestGroup{
		"UpdateManyWords": func(t *testing.T) {
			body := tools.MakeJsonBody(map[int]string{
				1: "no",
				2: "no2",
			})
			w := tools.ServeHTTP(gRouter, http.MethodPut, "/word/updateMany", body)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "ok", w.Body.String())
		},
		"Err_": func(t *testing.T) {
			body := tools.MakeJsonBody(map[int]any{
				1: 3,
				2: "no2",
			})
			w := tools.ServeHTTP(gRouter, http.MethodPut, "/word/updateMany", body)
			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, "bad body", w.Body.String())
		},
	})
}

func Test_Group_DeleteOneWord(t *testing.T) {
	tools.RunGroup(t, tools.TestGroup{
		"DeleteOneWord": func(t *testing.T) {
			w := tools.ServeHTTP(gRouter, http.MethodDelete, "/word/deleteOne?id=1", nil)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "ok", w.Body.String())
		},
		"Err_NotExist": func(t *testing.T) {
			w := tools.ServeHTTP(gRouter, http.MethodDelete, "/word/deleteOne?id=0", nil)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "ok", w.Body.String())
		},
		"Err_BadQuery": func(t *testing.T) {
			w := tools.ServeHTTP(gRouter, http.MethodDelete, "/word/deleteOne?id=hello", nil)
			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, `strconv.Atoi: parsing "hello": invalid syntax`, w.Body.String())
		},
	})
}

func Test_Group_DeleteManyWords(t *testing.T) {
	tools.RunGroup(t, tools.TestGroup{
		"DeleteManyWords": func(t *testing.T) {
			body := tools.MakeJsonBody([]int{1, 2, 3})
			w := tools.ServeHTTP(gRouter, http.MethodDelete, "/word/deleteMany", body)
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "ok", w.Body.String())
		},
		"Err_BadBody": func(t *testing.T) {
			body := tools.MakeJsonBody([]string{"1", "2", "hello"})
			w := tools.ServeHTTP(gRouter, http.MethodDelete, "/word/deleteMany", body)
			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, "json: cannot unmarshal string into Go value of type int", w.Body.String())
		},
	})
}

func TestQuickStart(t *testing.T) {
	w := tools.ServeHTTP(gRouter, http.MethodDelete, "/word/deleteAll", nil)
	w = tools.ServeHTTP(gRouter, http.MethodGet, "/word/all", nil)
	assert.Equal(t, "[]", w.Body.String())
	tools.ServeHTTP(gRouter, http.MethodPost, "/word/addOne?tag=hello&lang=eng&word=go", nil)
	w = tools.ServeHTTP(gRouter, http.MethodGet, "/word/all", nil)
	assert.Equal(t, `[{"Id":1,"Tag":"hello","Lang":"eng","Word":"go"}]`, w.Body.String())
	tools.ServeHTTP(gRouter, http.MethodPut, "/word/updateOne?id=3&word=hello", nil)
	w = tools.ServeHTTP(gRouter, http.MethodGet, "/word/all", nil)
	assert.Equal(t, `[{"Id":1,"Tag":"hello","Lang":"eng","Word":"go"}]`, w.Body.String())
	tools.ServeHTTP(gRouter, http.MethodDelete, "/word/deleteOne?id=1", nil)
	w = tools.ServeHTTP(gRouter, http.MethodGet, "/word/all", nil)
	assert.Equal(t, "[]", w.Body.String())
}
