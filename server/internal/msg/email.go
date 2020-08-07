/*
@Time : 2020/7/13 9:49 下午
@Author : lucbine
@File : email.go
*/
package msg

import (
	"bytes"
	"hotwheels/agent/entity"
	"hotwheels/agent/internal/config"
	"html/template"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

var mailTpl *template.Template

func init() {
	mailTpl, _ = template.New("mail_tpl").Parse(`
	你好 {{.username}}，<br/>

<p>以下是任务执行结果：</p>

<p>
任务ID：{{.task_id}}<br/>
任务名称：{{.task_name}}<br/>       
执行时间：{{.start_time}}<br />
执行耗时：{{.process_time}}秒<br />
执行状态：{{.status}}
</p>
<p>-------------以下是任务执行输出-------------</p>
<p>{{.output}}</p>
<p>
--------------------------------------------<br />
本邮件由系统自动发出，请勿回复<br />
如果要取消邮件通知，请登录到系统进行设置<br />
</p>
`)

}

type EmailNotice struct {
	email *email.Email
	auth  smtp.Auth
}

var emailNotice EmailNotice

func (en EmailNotice) Init() {
	//配置初始化
	emailNotice.email = email.NewEmail()
	emailNotice.email.From = config.GetString("config.email.from")
	emailNotice.auth = smtp.PlainAuth("", config.GetString("config.email.username"),
		config.GetString("config.email.password"), config.GetString("config.email.host"))

}

func (en EmailNotice) Send(nt entity.Notice) error {
	//https://studygolang.com/articles/26619?fr=sidebar
	content := new(bytes.Buffer)
	mailTpl.Execute(content, nt)
	//ccList := make([]string, 0)
	//if j.task.NotifyEmail != "" {
	//	ccList = strings.Split(j.task.NotifyEmail, "\n")
	//}
	en.email.To = []string{"719756455@qq.com"}
	en.email.Subject = "Awesome web"
	en.email.Text = []byte("Text Body is, of course, supported!")
	err := en.email.Send("smtp.126.com:25", en.auth)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
