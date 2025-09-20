package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Xib1uvXi/pumpimggtw"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ipfsNode, err := pumpimggtw.NewIPFSNode()
	if err != nil {
		log.Fatalf("Failed to create IPFS node: %v", err)
	}

	gtw, err := pumpimggtw.NewGateway(ipfsNode, pumpimggtw.NewPublicGateway("https://ipfs.io"))
	if err != nil {
		log.Fatalf("Failed to create gateway: %v", err)
	}

	r := gin.Default()
	r.GET("/ipfs/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		body, err := gtw.Get(cid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Data(http.StatusOK, "image/jpeg", body)
	})

	r.Run(":" + port)
}
