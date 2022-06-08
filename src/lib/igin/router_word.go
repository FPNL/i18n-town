package igin

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/FPNL/i18n-town/src/core/entity"
	"github.com/FPNL/i18n-town/src/core/service"
	"github.com/FPNL/i18n-town/src/lib/ierror"

	"github.com/gin-gonic/gin"
)

// inside project, just directly send error msg

func fetchAllWords(c *gin.Context) {
	if resp, err := service.WordService().FetchAllWords(); err == nil {
		c.JSON(http.StatusOK, resp)
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func addOneWord(c *gin.Context) {
	w := &entity.Word{}

	if c.BindQuery(w) != nil {
		c.String(http.StatusBadRequest, "missing query")
		return
	}

	err := service.WordService().AddOneWord(w)
	if err == nil {
		c.String(http.StatusOK, "ok")
	} else if errors.As(err, &ierror.Validate) {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.String(http.StatusInternalServerError, err.Error())

	}
}

func addManyWords(c *gin.Context) {
	var s = make([]entity.Word, 0)
	err := c.BindJSON(&s)
	if err != nil {
		c.Writer.WriteString("missing query")
		return
	}

	err = service.WordService().AddManyWords(s)
	if err == nil {
		c.String(http.StatusOK, "ok")
	} else if errors.As(err, &ierror.Validate) {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func updateOneWord(c *gin.Context) {
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

	err = service.WordService().UpdateOneWord(id, word)
	if err == nil {
		c.String(http.StatusOK, "ok")
	} else if errors.As(err, &ierror.Validate) {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		// inside project, just directly send error msg
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func updateManyWords(c *gin.Context) {
	q := make(map[int]string)
	if c.BindJSON(&q) != nil {
		c.Writer.WriteString("bad body")
		return
	}

	err := service.WordService().UpdateManyWords(q)
	if err == nil {
		c.String(http.StatusOK, "ok")
	} else if errors.As(err, &ierror.Validate) {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		// inside project, just directly send error msg
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func deleteOneWord(c *gin.Context) {
	id, err := strconv.Atoi(c.DefaultQuery("id", "-1"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = service.WordService().DeleteOneWord(id)
	if err == nil {
		c.String(http.StatusOK, "ok")
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func deleteManyWord(c *gin.Context) {
	s := make([]int, 0)
	err := c.BindJSON(&s)
	if err != nil {
		c.Writer.WriteString(err.Error())
		return
	}

	err = service.WordService().DeleteManyWord(s)
	if err == nil {
		c.String(http.StatusOK, "ok")
	} else {
		// inside project, just directly send error msg
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func deleteAll(c *gin.Context) {
	if err := service.WordService().DeleteAll(); err == nil {
		c.String(http.StatusOK, "ok")
	} else {
		// inside project, just directly send error msg
		c.String(http.StatusInternalServerError, err.Error())
	}
}
