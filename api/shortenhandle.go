package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"shortLink/data"
)

func Shorten(c *gin.Context) {
	sReq := data.ShortenReq{}
	if err := c.Bind(&sReq); err != nil {
		log.Printf("[ERROR] Context bind ShortenReq err: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
		c.Abort()
		return
	}
	fmt.Println(sReq.Url, sReq.Ttl)
	s, err := data.Shorten(sReq.Url, sReq.Ttl)
	if err != nil {
		log.Printf("[ERROR] data.Shorten err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"short_link": c.Request.Host + "/s/" + s,
	})
	c.Next()
}

func Redirect(c *gin.Context) {
	sReq := data.ShortLinkReq{}
	if err := c.ShouldBindUri(&sReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
		c.Abort()
		return
	}
	url, err := data.Redirect(sReq.ShortLink)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		c.Abort()
		return
	}
	if url == "" {
		c.JSON(http.StatusNotFound, gin.H{})
		c.Abort()
		return
	}
	c.Redirect(http.StatusFound, url)
}
