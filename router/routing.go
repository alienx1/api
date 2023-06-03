package router

import (
	"api/controller"
	"api/controller/auth"
	"api/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		admin := api.Group("/admin")
		{
			Auth := admin.Group("/auth")
			{
				Auth.POST("/login/admin", auth.AdminLogin)
				Auth.POST("/login/user", auth.UserLogin)
				Auth.POST("/register", auth.Register)
			}
			get := admin.Group("/get", middleware.LockAdmin())
			{
				// admin
				get.GET("/table/admin/:user/:name", controller.FindTableAdmin)
				get.GET("/ban/admin/:user/:status", controller.BanAdmin)
				// user
				get.GET("/table/user/:user/:name/:phone/:id_card", controller.FindTableUser)
				get.GET("/ban/user/:user", controller.BanUser)
				get.GET("/duration/user/:airtime/:user", controller.AddAirTime)
				//get animal
				get.GET("table/animal/:animal_id/:owner_name/:name/:type", controller.FindTableAnimal)
				get.GET("/pet/animal/:animal_id", controller.GetTableAnimalByAnimalIDWithPet)
				get.GET("/pet/animal/owner/:user", controller.UserProfile)
			}
			post := admin.Group("/post", middleware.LockAdmin())
			{
				// admin
				post.POST("/create/admin", controller.CreateAdmin)
				// user
				post.POST("/create/user", controller.CreateUser)
				// Animal
				post.POST("/craete/animal", controller.CreateAnimal)
			}
			put := admin.Group("/put", middleware.LockAdmin())
			{
				// admin
				put.PUT("/admin/:user", controller.UpdateAdmin)
				// user
				put.PUT("/user/:usre/:id", controller.UpdateUser)
			}
			delete := admin.Group("delete", middleware.LockAdmin())
			{
				// admin
				delete.DELETE("/admin/:user", controller.DeleteAdmin)
				// user
				delete.DELETE("/user/:user", controller.DeleteUser)
			}

		}
		user := api.Group("/user")
		{
			public := user.Group("/public")
			{
				public.GET("table-animal/:animal_id/:owner_name/:name/:type", controller.FindTableAnimal)
				public.GET("profile/:user", controller.UserProfile)

			}
			private := user.Group("/private")
			{
				private.GET("/profile/:user", controller.UserProfile)
				private.GET("/tableAnimal/:animal_id/:owner_name/:name/:type", controller.FindTableAnimal)
				private.POST("/animal-register", controller.CreateAnimal)
			}

		}
	}
	return r
}
