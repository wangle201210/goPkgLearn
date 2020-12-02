package routers

import (
	"beegoTest/controllers"

	"net/http"

	"github.com/astaxie/beego"
)

func handleFun(w http.ResponseWriter)  {

}
func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Handler("/used/abc", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("abc"))
	//}))
	//beego.Handler("/used/*abc", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("*abc"))
	//}))
}
