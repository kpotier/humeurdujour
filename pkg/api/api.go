package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kpotier/humeurdujour/pkg/store"
)

type API struct {
	Log   *log.Logger
	Store store.Store
}

func (a *API) Routes(g *gin.RouterGroup) {
	g.POST("/update", func(ctx *gin.Context) {
		t, err := time.Parse(time.RFC3339, ctx.PostForm("date"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			a.Log.Println(err)
			return
		}

		mood := &store.Mood{
			Date:         t,
			Sleep:        ctx.PostForm("sleep"),
			Global:       ctx.PostForm("global"),
			Anxiety:      ctx.PostForm("anxiety"),
			DarkThoughts: ctx.PostForm("darkthoughts"),
			Ruminations:  ctx.PostForm("ruminations"),
			Notes:        ctx.PostForm("notes"),
		}

		err = a.Store.SetMood(mood)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			a.Log.Println(err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	})

	g.POST("/get", func(ctx *gin.Context) {
		from, err := time.Parse(time.RFC3339, ctx.PostForm("from"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			a.Log.Println(err)
			return
		}

		to, err := time.Parse(time.RFC3339, ctx.PostForm("to"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			a.Log.Println(err)
			return
		}

		mood, err := a.Store.GetMood(from, to)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			a.Log.Println(err)
			return
		}
		ctx.JSON(http.StatusOK, mood)
	})
}
