// @BeeOverwrite YES
// @BeeGenerateTime 20200911_220506
package routers

import (
	"beegoTest/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/example", &controllers.ExampleController{}, "get:GetAll")
	beego.Router("/api/example/:id", &controllers.ExampleController{}, "get:GetOne")
	beego.Router("/api/example", &controllers.ExampleController{}, "post:Post")
	beego.Router("/api/example/:id", &controllers.ExampleController{}, "put:Put")
	beego.Router("/api/example/:id", &controllers.ExampleController{}, "delete:Delete")
}
