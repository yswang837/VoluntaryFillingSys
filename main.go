package main

import (
	"fmt"
	"github.com/VoluntaryFillingSys/pkg/snowflake"
	"github.com/VoluntaryFillingSys/utils"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func main() {
	if err := snowflake.Init(utils.StartTime, utils.MachinedId); err != nil {
		return
	}
	fmt.Println(snowflake.Node.Generate())
	e := email.NewEmail()
	e.From = "wys <1714113169@qq.com>"
	e.To = []string{"1714113169@qq.com"}
	e.Bcc = []string{"1714113169@qq.com"}
	e.Cc = []string{"1714113169@qq.com"}
	e.Text = []byte("本文邮件内容")
	e.HTML = []byte("<h1>html 邮件内容!</h1>")
	if err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "1714113169@qq.com", "bxvjwuyyeoqfdced", "smtp.qq.com")); err != nil {
		fmt.Println(err)
	}

}

//bxvjwuyyeoqfdced
