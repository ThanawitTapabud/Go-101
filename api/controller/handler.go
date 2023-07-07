package controller

import (
	"context"
	"fmt"
	"net/http"

	"api/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetValorantTeam(c *gin.Context) {
	name := "Pond"
	c.JSON(http.StatusOK, name)

}

func GetValorantTeamByID(c *gin.Context) {
	paramID := c.Param("id")
	for _, valorantTeam := range ValorantTeams {
		if valorantTeam.ID == paramID {
			c.JSON(http.StatusOK, valorantTeam)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Data not found"})
}

func AddValorantTeam(c *gin.Context) {
	var newValorantTeam model.ValorantTeamDetail

	if err := c.ShouldBindJSON(&newValorantTeam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Add Failed"})
		return
	}
	clientOption := options.Client().ApplyURI("mongodb://root:example@localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOption)
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
	}
	db := client.Database("valorantTeamdb")
	collection := db.Collection("valorantTeam")
	result, err := collection.InsertOne(context.Background(), newValorantTeam)

	c.JSON(http.StatusCreated, gin.H{
		"id":   result.InsertedID,
		"data": newValorantTeam})
}

func UpdateValorantTeam(c *gin.Context) {
	var editValorantTeam model.ValorantTeamDetail

	if err := c.ShouldBindJSON(&editValorantTeam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Update Failed"})
		return
	}

	paramID := c.Param("id")
	for i := 0; i <= len(ValorantTeams)-1; i++ {
		if ValorantTeams[i].ID == paramID {
			ValorantTeams[i].Name = editValorantTeam.Name
			ValorantTeams[i].Location = editValorantTeam.Location
			ValorantTeams[i].Description = editValorantTeam.Description

			c.IndentedJSON(http.StatusOK, ValorantTeams[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Data not found"})
}

func DeleteValorantTeamByID(c *gin.Context) {
	paramID := c.Param("id")
	for i := 0; i <= len(ValorantTeams)-1; i++ {
		if ValorantTeams[i].ID == paramID {
			ValorantTeams = append(ValorantTeams[:i], ValorantTeams[i+1:]...)
			c.JSON(http.StatusOK, "delete success")
			return
		}
	}
	c.JSON(http.StatusNotFound, "data not found")
}
