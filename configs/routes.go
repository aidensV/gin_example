package configs

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aidensV/gin_example/helpers"
	"github.com/aidensV/gin_example/models"
	"github.com/aidensV/gin_example/repositories"
	"github.com/aidensV/gin_example/services"
)

func SetupRoutes(contactRepository *repositories.ContactRepository) *gin.Engine {
	route := gin.Default()

	//create route /create endpoint
	route.POST("/create", func(context *gin.Context) {
		//initialization contact model
		var contact models.Contact

		//validation json
		err := context.ShouldBind(&contact)

		//validation errors
		if err != nil {
			//generate  validation errors response
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)
			return
		}
		//default http status code = 200
		code := http.StatusOK

		//save contact & get its response
		response := services.CreateContact(&contact, *contactRepository)

		//save contact failed
		if !response.Success {
			// change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/", func(context *gin.Context) {
		code := http.StatusOK

		response := services.FindAllContacts(*contactRepository)
		if !response.Success {
			code = http.StatusBadRequest
		}
		context.JSON(code, response)
	})

	route.GET("/show/:id", func(context *gin.Context) {
		id := context.Param("id")

		code := http.StatusOK

		response := services.FindByIdContact(id, *contactRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}
		context.JSON(code, response)
	})

	route.PUT("/update/:id", func(context *gin.Context) {
		id := context.Param("id")

		var contact models.Contact

		err := context.ShouldBindJSON(&contact)

		//validate errors
		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := services.UpdateContactById(id, &contact, *contactRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)

	})

	route.DELETE("/delete/:id", func(context *gin.Context) {
		id := context.Param("id")

		code := http.StatusOK

		response := services.DeleteContactById(id, *contactRepository)
		if !response.Success {
			code = http.StatusBadRequest
		}
		context.JSON(code, response)
	})
	return route
}
