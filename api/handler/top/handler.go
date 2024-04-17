package top

import (
	"io"
	"net/http"
	"proxy-service/conf"

	"github.com/gin-gonic/gin"
)

type TopHandler interface {
	Proxy(c *gin.Context)
}

func (h *TopInstance) Proxy(c *gin.Context) {
	// 创建代理请求的URL
	proxyURL := conf.Conf.OpenTaoBao.ServerUrl + c.Param("proxyPath")
	proxyReq, err := http.NewRequest(c.Request.Method, proxyURL, c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// 复制原始请求的Header到代理请求
	for h, v := range c.Request.Header {
		proxyReq.Header[h] = v
	}

	// 创建HTTP客户端并发送请求
	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	// 将响应Header设置到Gin上下文
	for h, vv := range resp.Header {
		for _, v := range vv {
			c.Writer.Header().Add(h, v)
		}
	}

	// 设置响应状态码
	c.Status(resp.StatusCode)

	// 复制响应内容到Gin响应中
	io.Copy(c.Writer, resp.Body)
}
