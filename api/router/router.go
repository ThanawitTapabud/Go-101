package router

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func RouterPath() {
	r := gin.Default()
	r.GET("/valorant/info", controller.GetValorantTeam)
	r.GET("/valorant/info/:id", controller.GetValorantTeamByID)
	r.POST("/valorant/info", controller.AddValorantTeam)
	r.PUT("/valorant/update/:id", controller.UpdateValorantTeam)
	r.DELETE("/valorant/delete/:id", controller.DeleteValorantTeamByID)
	r.Run(":8081")
}
