// @BeeOverwrite YES
// @BeeGenerateTime 20200911_220506
package controllers

import (
	"beegoTest/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

//  ExampleController operations for Example
type ExampleController struct {
	beego.Controller
}

// URLMapping ...
func (c *ExampleController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Example
// @Param	body		body 	models.Example	true		"body for Example content"
// @Success 201 {int} models.Example
// @Failure 403 body is empty
// @router / [post]
func (c *ExampleController) Post() {
	var v models.Example
	c.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	_, err := models.AddExample(&v)
	if err != nil {
		c.Data["json"] = ResponseData{
			Code:    1,
			Message: err.Error(),
		}
		c.ServeJSON()
		return
	}
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = ResponseData{
		Code:    0,
		Message: "success",
		Data:    v,
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Example by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Example
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ExampleController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetExampleById(id)
	if err != nil {
		c.Data["json"] = ResponseData{
			Code:    1,
			Message: err.Error(),
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = ResponseData{
		Code:    0,
		Message: "success",
		Data:    v,
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Example
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Example
// @Failure 403
// @router / [get]
func (c *ExampleController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var pageSize int = 20
	var current int

	if v := c.GetString("id"); v != "" {
		query["id"] = v
	}

	if v := c.GetString("row1"); v != "" {
		query["row1"] = v
	}

	if v := c.GetString("row2"); v != "" {
		query["row2"] = v
	}

	if v := c.GetString("row3"); v != "" {
		query["row3"] = v
	}

	if v := c.GetString("row4"); v != "" {
		query["row4"] = v
	}

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt("pageSize"); err == nil {
		pageSize = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt("current"); err == nil {
		current = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	l, _ := models.GetAllExample(query, fields, sortby, order, current, pageSize)
	c.Data["json"] = l
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Example
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Example	true		"body for Example content"
// @Success 200 {object} models.Example
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ExampleController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Example{
		Id: id,
	}

	c.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	err := models.UpdateExampleById(&v)
	if err != nil {
		c.Data["json"] = ResponseData{
			Code:    1,
			Message: err.Error(),
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = ResponseData{
		Code:    0,
		Message: "success",
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Example
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ExampleController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	err := models.DeleteExample(id)
	if err != nil {
		c.Data["json"] = ResponseData{
			Code:    1,
			Message: err.Error(),
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = ResponseData{
		Code:    0,
		Message: "success",
	}
	c.ServeJSON()
}
