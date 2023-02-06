package router

import (
	"net/http"
	models_rep "q1/models/repository"
	svc "q1/service/interface"

	"github.com/gin-gonic/gin"
)

type IRouter interface {
	InitRouter() *gin.Engine
}

type Router struct {
	CommentSvc svc.ICommentSvc
}

func NewRouter(ICommentSvc svc.ICommentSvc) IRouter {
	return &Router{
		CommentSvc: ICommentSvc,
	}
}

func (router *Router) InitRouter() *gin.Engine {
	r := gin.Default()
	g1 := r.Group("/quiz/v1/")
	g1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	// Comment
	g1.POST("/comment", router.createComment)
	g1.GET("/comment/:uuid", router.getComment)
	g1.PUT("/comment/:uuid", router.updateComment)
	g1.DELETE("/comment/:uuid", router.deleteComment)
	return r
}

// region CRUD Comment
func (router *Router) createComment(c *gin.Context) {
	var payload models_rep.Comment
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": err.Error()})
		return
	}
	result, errRsp := router.CommentSvc.CreateComment(&payload)
	if errRsp != nil {
		c.JSON(errRsp.StatusCode, errRsp)
		return
	}
	c.JSON(http.StatusOK, result)
}
func (router *Router) getComment(c *gin.Context) {
	uuid := c.Param("uuid")
	result, errRsp := router.CommentSvc.GetComment(uuid)
	if errRsp != nil {
		c.JSON(errRsp.StatusCode, errRsp)
		return
	}
	c.JSON(http.StatusOK, result)
}
func (router *Router) updateComment(c *gin.Context) {
	uuid := c.Param("uuid")
	var payload models_rep.Comment
	payload.Uuid = uuid
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": err.Error()})
		return
	}
	result, errRsp := router.CommentSvc.UpdateComment(&payload)
	if errRsp != nil {
		c.JSON(errRsp.StatusCode, errRsp)
		return
	}
	c.JSON(http.StatusOK, result)
}
func (router *Router) deleteComment(c *gin.Context) {
	uuid := c.Param("uuid")
	errRsp := router.CommentSvc.DeleteComment(uuid)
	if errRsp != nil {
		c.JSON(http.StatusInternalServerError, errRsp)
		return
	}
	c.JSON(http.StatusOK, gin.H{"description": "ok"})
}

// endregion
