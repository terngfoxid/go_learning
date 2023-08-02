package Controllers

import (
	"first-api/Models"
	"first-api/httputil"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers ... Get All user
//
//	@Summary		Get All User Data
//	@Description	get string
//	@Tags			User
//	@Produce		json
//	@Success		200	{array}	Models.User
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/user [get]
func GetUsers(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUsers(&user)
	if err != nil {
		httputil.NewError(c, http.StatusNotFound, err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// CreateUser ... Create User
//
//	@Summary		Create User
//	@Description	Add user by JSON
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user body Models.User true "Create User"
//	@Success		200	{object}	Models.User
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/user [post]
func CreateUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		httputil.NewError(c, http.StatusNotFound, err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetUserByID ... Get the user by id
//
//	@Summary		Get User Data
//	@Description	get string by ID
//	@Tags			User
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	Models.User
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/user/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUserByID(&user, id)
	if err != nil {
		httputil.NewError(c, http.StatusNotFound, err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// UpdateUser ... Update the user information
//
//	@Summary		Update User
//	@Description	Update user by JSON
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user body Models.User true "Update User"
//	@Success		200	{object}	Models.User
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/user [put]
func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&user, id)
	if err != nil {
		httputil.NewError(c, http.StatusNotFound, err)
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil {
		httputil.NewError(c, http.StatusNotFound, err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// DeleteUser ... Delete the user
//
//	@Summary		Delete User
//	@Description	Delete user by ID
//	@Tags			User
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	httputil.HTTPSuccess
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/user/{id} [delete]
func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		httputil.NewError(c, http.StatusNotFound, err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "status ok",
			"Detail":  "id " + id + " is deleted"})
	}
}
