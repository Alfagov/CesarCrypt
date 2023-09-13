package handlers

import (
	"github.com/Alfagov/CesarCrypt/models"
	"github.com/Alfagov/CesarCrypt/pkg/store"
	"github.com/Alfagov/CesarCrypt/utils"
	"github.com/gin-gonic/gin"
)

func NewStoreBackend(storeBackends *models.ServerStores) gin.HandlerFunc {
	return func(c *gin.Context) {
		backendType := c.Param("storeBackend")
		if backendType == "" {
			c.JSON(400, gin.H{"error": "store backend type is required"})
			return
		}

		backendName := c.Param("backendName")
		if backendName == "" {
			c.JSON(400, gin.H{"error": "store backend name is required"})
			return
		}

		backendQualifiedName := utils.GenerateStoreBackendName(backendType, backendName)

		if storeBackends.StoreHandlers[backendQualifiedName] != nil &&
			storeBackends.StoreHandlers[backendQualifiedName].IsOpen() {
			c.JSON(400, gin.H{"error": "store backend already exists"})
			return
		}

		var request models.CreateStoreBackendRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		storeBackend := store.BackendConstructors[backendType](backendName, request.Settings)

		storeBackends.Mux.Lock()
		storeBackends.StoreHandlers[backendQualifiedName] = storeBackend
		storeBackends.Mux.Unlock()
	}
}
