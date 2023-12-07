package dependencies

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	// h := SocketInitiate()
	// go h.SocketRun()

	// router.GET("/ws", gin.WrapF(http.HandlerFunc(SocketAllowUpgrade(h))))
	// router.GET("/ws/bid", gin.WrapF(http.HandlerFunc(BidScore(h))))
	router.POST("/run", RunHandler)
	router.POST("/terminate", TerminateHandler)
	router.POST("/restart", RestartHandler)

}

type RequestBody struct {
	Ws string `json:"ws"`
}

func RunHandler(c *gin.Context) {
	// Get path from request body
	var body RequestBody
	// Bind the JSON request body to the struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ws := body.Ws
	wsPath, err := ReadWs(ws)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if exe is already running or not
	_, ok := RunningExes[wsPath]
	if ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Exe is already running"})
		return
	}

	err = RunningExe(wsPath, ws)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Running " + ws})
}

func TerminateHandler(c *gin.Context) {
	// Get path from request body
	var body RequestBody
	// Bind the JSON request body to the struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ws := body.Ws
	wsPath, err := ReadWs(ws)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if exe is already terminated or not
	_, ok := RunningExes[wsPath]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Exe is already terminated"})
		return
	}

	err = TerminateExe(wsPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Terminating " + ws})

}

func RestartHandler(c *gin.Context) {
	// Get path from request body
	var body RequestBody
	// Bind the JSON request body to the struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ws := body.Ws
	wsPath, err := ReadWs(ws)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if exe is already terminated or not, if not just run the command
	_, ok := RunningExes[wsPath]
	if !ok {
		err = RunningExe(wsPath, ws)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Running " + ws})
		return
	}

	err = TerminateExe(wsPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	time.Sleep(2 * time.Second)

	err = RunningExe(wsPath, ws)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Restarting " + ws})
}
