package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/FPNL/i18n-town/src/core/entity"
	"github.com/FPNL/i18n-town/src/core/service"
	"github.com/FPNL/i18n-town/src/lib/ierror"

	"github.com/gin-gonic/gin"
)

type IWordHandler interface {
	FetchAllWords(c *gin.Context)
	AddOneWord(c *gin.Context)
	AddManyWords(c *gin.Context)
	UpdateOneWord(c *gin.Context)
	UpdateManyWords(c *gin.Context)
	DeleteOneWord(c *gin.Context)
	DeleteManyWord(c *gin.Context)
	DeleteAll(c *gin.Context)
	Auth(*gin.Context)
	Login(*gin.Context)
}

type wordHandler struct {
	wordService service.IWordService
}

var singleWord = wordHandler{}

// Word prefix is /api/v1
func Word(wordService service.IWordService) IWordHandler {
	singleWord.wordService = wordService
	return &singleWord
}

func (hd *wordHandler) FetchAllWords(c *gin.Context) {
	if resp, err := hd.wordService.FetchAllWords(); err == nil {
		c.JSON(http.StatusOK, resp)
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func (hd *wordHandler) AddOneWord(c *gin.Context) {
	w := &entity.Word{}

	if c.BindQuery(w) != nil {
		c.String(http.StatusBadRequest, "missing query")
		return
	}

	err := hd.wordService.AddOneWord(w)
	if err == nil {
		c.String(http.StatusOK, "ok")
	} else if errors.As(err, &ierror.Validate) {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.String(http.StatusInternalServerError, err.Error())

	}
}

func (hd *wordHandler) AddManyWords(c *gin.Context) {
	var s = make([]entity.Word, 0)
	err := c.BindJSON(&s)
	if err != nil {
		c.Writer.WriteString("missing query")
		return
	}

	err = hd.wordService.AddManyWords(s)
	if err == nil {
		c.String(http.StatusOK, "ok")
	} else if errors.As(err, &ierror.Validate) {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func (hd *wordHandler) UpdateOneWord(c *gin.Context) {
	id, err := strconv.Atoi(c.DefaultQuery("id", "-1"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	word, ok := c.GetQuery("word")
	if !ok {
		c.String(http.StatusBadRequest, "missing query")
		return
	}

	err = hd.wordService.UpdateOneWord(id, word)
	if err == nil {
		c.String(http.StatusOK, "ok")
	} else if errors.As(err, &ierror.Validate) {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		// inside project, just directly send error msg
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func (hd *wordHandler) UpdateManyWords(c *gin.Context) {
	q := make(map[int]string)
	if c.BindJSON(&q) != nil {
		c.Writer.WriteString("bad body")
		return
	}

	err := hd.wordService.UpdateManyWords(q)
	if err == nil {
		c.String(http.StatusOK, "ok")
	} else if errors.As(err, &ierror.Validate) {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		// inside project, just directly send error msg
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func (hd *wordHandler) DeleteOneWord(c *gin.Context) {
	id, err := strconv.Atoi(c.DefaultQuery("id", "-1"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = hd.wordService.DeleteOneWord(id)
	if err == nil {
		c.String(http.StatusOK, "ok")
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func (hd *wordHandler) DeleteManyWord(c *gin.Context) {
	s := make([]int, 0)
	err := c.BindJSON(&s)
	if err != nil {
		c.Writer.WriteString(err.Error())
		return
	}

	err = hd.wordService.DeleteManyWord(s)
	if err == nil {
		c.String(http.StatusOK, "ok")
	} else {
		// inside project, just directly send error msg
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func (hd *wordHandler) DeleteAll(c *gin.Context) {
	if err := hd.wordService.DeleteAll(); err == nil {
		c.String(http.StatusOK, "ok")
	} else {
		// inside project, just directly send error msg
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func (hd wordHandler) Auth(*gin.Context) {

}

func (hd wordHandler) Login(*gin.Context) {

}
