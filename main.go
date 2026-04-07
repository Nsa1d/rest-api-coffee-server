package main

import (
	"net/http"
	"rest-api-coffee-server/menu"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/drinks", GetAllDrinks)
	r.GET("/drinks/in-stock", AllAvailableDrinks)
	r.GET("/drinks/:id", GetDrinkByID)
	r.POST("/drinks", CreateDrinks)
	r.DELETE("/drinks/:id", DeleteDrinks)
	r.PATCH("/drinks/:id", UpdateDrinks)

	r.Run(":8080")
}

func GetAllDrinks(c *gin.Context) {
	drinksList, err := menu.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, drinksList)
}

func AllAvailableDrinks(c *gin.Context) {
	drinksList, err := menu.AvailableDrinks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, drinksList)
}

func GetDrinkByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID должен быть числом",
		})
		return
	}

	drink, err := menu.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, drink)
}

func CreateDrinks(c *gin.Context) {
	var req menu.DrinkCreate

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	drink, err := menu.Add(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, drink)
}

func DeleteDrinks(c *gin.Context) {
	id := c.Param("id")

	err := menu.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "напиток успешно удалён",
	})
}

func UpdateDrinks(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "неверный формат ID",
		})
		return
	}
	var req menu.DrinkUpdate

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.ID = id

	drink, err := menu.Update(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, drink)
}
