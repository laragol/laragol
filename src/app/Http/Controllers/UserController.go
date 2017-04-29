package Controllers

import (
	"app/Models"
	"encoding/json"
	"strconv"

	"github.com/laragol/framework/Database/Connection"
	"github.com/laragol/framework/Response"
	"github.com/qiangxue/fasthttp-routing"
)

type UserController struct {
}

func (UserController) Index(c *routing.Context) error {
	db := Connection.Get()
	var users []Models.User
	db.Preload("Country").Find(&users)
	return Response.Json(c, users)
}

func (UserController) Show(c *routing.Context) error {
	db := Connection.Get()
	var user Models.User
	db.Preload("Country").Find(&user, c.Param("userId"))
	return Response.Json(c, user)
}

func (UserController) Store(c *routing.Context) error {
	db := Connection.Get()
	body := c.PostBody()
	user := &Models.User{}
	json.Unmarshal(body, user)
	user.CountryID = user.Country.ID
	db.Set("gorm:save_associations", false).Save(&user)
	return Response.Json(c, user)
}

func (UserController) Update(c *routing.Context) error {
	db := Connection.Get()
	body := c.PostBody()
	user := &Models.User{}
	json.Unmarshal(body, user)
	userId, _ := strconv.ParseUint(c.Param("userId"), 10, 64)
	user.ID = userId
	user.CountryID = user.Country.ID
	db.Save(&user)
	return Response.Json(c, user)
}
