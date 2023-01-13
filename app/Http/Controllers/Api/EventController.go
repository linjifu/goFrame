package Api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goFrame/app/Http/RequestValidation/ApiValidation"
	"goFrame/app/Models"
	"goFrame/app/Tools"
)

type EventController struct {
	Tools.Response
}

func (e *EventController) Index(c *gin.Context) {

	aModel := Models.AtableModel{}
	//Tools.DBConnection.Connections["mysql2"].Connection.Take(&aModel)
	//Tools.DBConnection.Db.Connection.Clauses(dbresolver.Use("default")).Take(&aModel)
	e.ReturnJsonSuccess(c, 200, "succ", aModel)

}

func (e *EventController) Activation(c *gin.Context) {
	//platform := c.DefaultQuery("send_platform", "")

	var jrttValidation ApiValidation.JrttValidation
	params, ok := Tools.Validation(c, &jrttValidation).(*ApiValidation.JrttValidation)
	if ok {
		fmt.Println("成功转换:", params.ProjectId)
	} else {
		e.ReturnJsonError(c, Tools.USER_AUTH_ERROR, "参数转换失败")
	}
}
