package dependencies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	h := SocketInitiate()
	go h.SocketRun()
	gin.SetMode(gin.ReleaseMode)
	router.GET("/ws", gin.WrapF(http.HandlerFunc(SocketAllowUpgrade(h))))
	router.GET("/ws/bid", gin.WrapF(http.HandlerFunc(BidScore(h))))

}
