package email

import (
	"fmt"
	"github.com/VoluntaryFillingSys/utils"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"time"
)

func SendEmail(to string) (string, error) {
	e := email.NewEmail()
	fmt.Println(utils.EmailFrom)
	e.From = utils.EmailFrom
	e.To = []string{to}
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	str := "<h3> 验证码：" + code + "</h3> <p>仅5分钟内有效,请勿转发他人!</p>"
	e.HTML = []byte(str)
	if err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "1714113169@qq.com", "bxvjwuyyeoqfdced", "smtp.qq.com")); err != nil {
		return "", err
	}
	return code, nil
}
