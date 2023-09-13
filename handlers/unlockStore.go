package handlers

import (
	"github.com/Alfagov/CesarCrypt/pkg/shamir"
	"github.com/Alfagov/CesarCrypt/pkg/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UnlockStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		storeName := c.Param("name")
		if storeName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "store name is required"})
			return
		}

		storeInterface, exists := c.Get("storeBackend")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "store handler not found"})
			return
		}

		storeHandler, ok := storeInterface.(store.Handler)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "store handler not found"})
			return
		}

		var secretShards []string
		err := c.ShouldBindJSON(&secretShards)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		shardBytes := make([][]byte, len(secretShards))
		for i, shard := range secretShards {
			shardBytes[i] = []byte(shard)
		}

		secret, err := shamir.Combine(shardBytes)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = storeHandler.Unlock(secret)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}
