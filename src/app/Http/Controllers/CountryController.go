package Controllers

import (
	"app/Models"
	"encoding/json"
	"strconv"

	"github.com/laragol/framework/Database/Connection"
	"github.com/laragol/framework/Response"
	"github.com/qiangxue/fasthttp-routing"
)

type CountryController struct {
}

func (CountryController) Index(c *routing.Context) error {
	db := Connection.Get()
	var countries []Models.Country
	db.Find(&countries)
	return Response.Json(c, countries)
}

func (CountryController) Show(c *routing.Context) error {
	db := Connection.Get()
	var country Models.Country
	db.Find(&country, c.Param("countryId"))
	return Response.Json(c, country)
}

func (CountryController) Store(c *routing.Context) error {
	db := Connection.Get()
	body := c.PostBody()
	country := &Models.Country{}
	json.Unmarshal(body, country)
	db.Save(&country)
	return Response.Json(c, country)
}

func (CountryController) Update(c *routing.Context) error {
	db := Connection.Get()
	body := c.PostBody()
	country := &Models.Country{}
	json.Unmarshal(body, country)
	countryId, _ := strconv.ParseUint(c.Param("countryId"), 10, 64)
	country.ID = countryId
	db.Save(&country)
	return Response.Json(c, country)
}
