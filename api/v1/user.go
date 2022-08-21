package v1

//import (
//	dao "github.com/VoluntaryFillingSys/dao/user"
//	"github.com/gin-gonic/gin"
//)
//
//func AddUser(c *gin.Context) {
//	var data dao.User
//	_ = c.ShouldBindJSON(&data)
//	//msg, code := validator.Validate(&data)
//	//if code != errmsg.Success {
//	//	c.JSON(http.StatusOK, gin.H{
//	//		"status":  code,
//	//		"message": msg,
//	//	})
//	//	return
//	//}
//	//code = model.CheckUser(data.UserName)
//	//if code == errmsg.Success {
//	//	model.CreateUser(&data)
//	//}
//	//if code == errmsg.ErrorUsernameUsed {
//	//	code = errmsg.ErrorUsernameUsed
//	//}
//	//c.JSON(http.StatusOK, gin.H{
//	//	"status": code,
//	//	"msg":    errmsg.GetErrMsg(code),
//	//})
//}
