package main

import (
	"pizzaria/internal/data"
	"pizzaria/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	data.LoadPizzas()
	router := gin.Default()
	router.GET("/pizzas", handler.GetPizzas)
	router.POST("/pizzas", handler.PostPizzas)
	router.GET("/pizzas/:id", handler.GetPizzasByID)
	router.DELETE("/pizzas/:id", handler.DeletePizzasByID)
	//editar ou atualizar uma pizza
	router.PUT("/pizzas/:id", handler.UptadePizzasByID)
	router.POST("/pizzas/:id/review", handler.PostReview)
	router.Run()
}
