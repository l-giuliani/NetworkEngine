package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/l-giuliani/networkEngine"
	"github.com/l-giuliani/networkEngine/dto"
	"github.com/l-giuliani/networkEngine/libs"
	"github.com/l-giuliani/networkEngine/libs/gameLibs"
)

func InitGameControllers(engine *gin.Engine)  {
	engine.POST("/game/create", func(c *gin.Context) {
		
		//var json components.User
		/*err := c.ShouldBindJSON(&json)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}*/
		idUser := libs.GenerateUserId()
		idGame := libs.GenerateGameId()

		token, err := libs.GenerateGameToken(idUser, idGame)
		if err != nil{
			fmt.Println("Error signin token")
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		//FIXME
		game, err := gameLibs.CreateGame(idGame, 4)
		if err != nil {
			fmt.Println(err)
			return
		}
		states := networkEngine.GetInitialGameStates()
		for s := range states {
			game.AddState(s, states[s])
		}
		
		c.JSON(200, gin.H{
			"idGame": idGame,
			"token" : token,
		})
	})

	engine.POST("/game/partecipate", func(c *gin.Context) {

		var json dto.GameRequestDto
		err := c.ShouldBindJSON(&json)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		idUser := libs.GenerateUserId()

		token, err := libs.GenerateGameToken(idUser, json.IdGame)
		if err != nil{
			fmt.Println("Error signin token")
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"idGame": json.IdGame,
			"token" : token,
		})
	})
}