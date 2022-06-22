package delivery

import (
	"github.com/FPNL/i18n-town/src/core/entity"
	"github.com/FPNL/i18n-town/src/core/service"
	"github.com/FPNL/i18n-town/src/lib/ierror"
	pb "github.com/FPNL/i18n-town/src/lib/igrpc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IWordDelivery interface {
	AddCommittedWords(*gin.Context)
	FetchCommittedWords(*gin.Context)
	UpdateCommittedWords(*gin.Context)
	DeleteCommittedWords(*gin.Context)
	//AdviseWords(*gin.Context)
	//FetchStageWords(*gin.Context)
	//CommitStageWords(*gin.Context)
	//DiscardStageWords(*gin.Context)
}

type wordDelivery struct {
	wordService service.IWord
}

var singletonWord = wordDelivery{}

// Word prefix is /api/v1
func Word(wordService service.IWord) IWordDelivery {
	singletonWord.wordService = wordService
	return &singletonWord
}

func (hd *wordDelivery) AddCommittedWords(ctx *gin.Context) {
	var err error

	user, err := fetchUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	var words []entity.CommittedWord
	err = ctx.BindJSON(&words)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = hd.wordService.AddCommittedWords(ctx, user, words)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.Status(http.StatusOK)
}

func (hd *wordDelivery) FetchCommittedWords(ctx *gin.Context) {
	var err error

	user, err := fetchUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	var conditions *entity.SearchCondition_CommittedWord
	err = ctx.BindJSON(conditions)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	words, total, err := hd.wordService.FetchCommittedWords(ctx, user, conditions)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, struct {
		words []entity.CommittedWord
		total uint
	}{words, total})
}

func (hd *wordDelivery) UpdateCommittedWords(ctx *gin.Context) {
	var err error

	user, err := fetchUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	var words = make(map[uint]string)
	err = ctx.BindJSON(&words)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = hd.wordService.UpdateCommittedWords(ctx, user, words)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (hd *wordDelivery) DeleteCommittedWords(ctx *gin.Context) {
	var err error
	user, err := fetchUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	var ids []uint
	err = ctx.BindJSON(&ids)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = hd.wordService.DeleteCommittedWords(ctx, user, ids)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
}

func (hd *wordDelivery) AdviseWords(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (hd *wordDelivery) FetchStageWords(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (hd *wordDelivery) CommitStageWords(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (hd *wordDelivery) DiscardStageWords(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

//func (hd *wordDelivery) AddOneWord(c *gin.Context) {
//	w := &entity.Word{}
//
//	if c.BindQuery(w) != nil {
//		c.String(http.StatusBadRequest, "missing query")
//		return
//	}
//
//	err := hd.wordService.AddOneWord(w)
//	if err == nil {
//		c.String(http.StatusOK, "ok")
//	} else if errors.As(err, &ierror.Validate) {
//		c.String(http.StatusBadRequest, err.Error())
//	} else {
//		c.String(http.StatusInternalServerError, err.Error())
//
//	}
//}
//
//func (hd *wordDelivery) AddManyWords(c *gin.Context) {
//	var s = make([]entity.Word, 0)
//	err := c.BindJSON(&s)
//	if err != nil {
//		c.Writer.WriteString("missing query")
//		return
//	}
//
//	err = hd.wordService.AddManyWords(s)
//	if err == nil {
//		c.String(http.StatusOK, "ok")
//	} else if errors.As(err, &ierror.Validate) {
//		c.String(http.StatusBadRequest, err.Error())
//	} else {
//		c.String(http.StatusInternalServerError, err.Error())
//	}
//}
//
//func (hd *wordDelivery) UpdateOneWord(c *gin.Context) {
//	id, err := strconv.Atoi(c.DefaultQuery("id", "-1"))
//	if err != nil {
//		c.String(http.StatusBadRequest, err.Error())
//		return
//	}
//
//	word, ok := c.GetQuery("word")
//	if !ok {
//		c.String(http.StatusBadRequest, "missing query")
//		return
//	}
//
//	err = hd.wordService.UpdateOneWord(id, word)
//	if err == nil {
//		c.String(http.StatusOK, "ok")
//	} else if errors.As(err, &ierror.Validate) {
//		c.String(http.StatusBadRequest, err.Error())
//	} else {
//		// inside project, just directly send error msg
//		c.String(http.StatusInternalServerError, err.Error())
//	}
//}
//
//func (hd *wordDelivery) UpdateManyWords(c *gin.Context) {
//	q := make(map[int]string)
//	if c.BindJSON(&q) != nil {
//		c.Writer.WriteString("bad body")
//		return
//	}
//
//	err := hd.wordService.UpdateManyWords(q)
//	if err == nil {
//		c.String(http.StatusOK, "ok")
//	} else if errors.As(err, &ierror.Validate) {
//		c.String(http.StatusBadRequest, err.Error())
//	} else {
//		// inside project, just directly send error msg
//		c.String(http.StatusInternalServerError, err.Error())
//	}
//}
//
//func (hd *wordDelivery) DeleteOneWord(c *gin.Context) {
//	id, err := strconv.Atoi(c.DefaultQuery("id", "-1"))
//	if err != nil {
//		c.String(http.StatusBadRequest, err.Error())
//		return
//	}
//
//	err = hd.wordService.DeleteOneWord(id)
//	if err == nil {
//		c.String(http.StatusOK, "ok")
//	} else {
//		c.String(http.StatusInternalServerError, err.Error())
//	}
//}
//
//func (hd *wordDelivery) DeleteManyWord(c *gin.Context) {
//	s := make([]int, 0)
//	err := c.BindJSON(&s)
//	if err != nil {
//		c.Writer.WriteString(err.Error())
//		return
//	}
//
//	err = hd.wordService.DeleteManyWord(s)
//	if err == nil {
//		c.String(http.StatusOK, "ok")
//	} else {
//		// inside project, just directly send error msg
//		c.String(http.StatusInternalServerError, err.Error())
//	}
//}
//
//func (hd *wordDelivery) DeleteAll(c *gin.Context) {
//	if err := hd.wordService.DeleteAll(); err == nil {
//		c.String(http.StatusOK, "ok")
//	} else {
//		// inside project, just directly send error msg
//		c.String(http.StatusInternalServerError, err.Error())
//	}
//}

func fetchUser(ctx *gin.Context) (*pb.User, error) {
	val, ok := ctx.Get("user_from_admin_delivery")
	if !ok {
		return nil, ierror.NewServerErr("auth already but there is no user")
	}

	user, ok := val.(*pb.User)
	if !ok {
		return nil, ierror.NewServerErr("auth already but there is no user type")
	}

	return user, nil
}
