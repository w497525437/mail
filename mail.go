package mail

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"strings"
)

/*发送邮件
 *from      发送账户
 *password  密码
 *host		SMTP服务器
 *port		端口号
 *name		发送人
 *subject	标题
 *filepath	模版文件路径
 *to 		收件人邮箱
 *return	成功返回空字符串，失败返回错误信息
 */
func SendMail(from string, password string, host string, port string, name string, subject string, filepath string, to []string) string {
	auth := smtp.PlainAuth("", from, password, host)
	content_type := "Content-Type: text/html; charset=UTF-8"
	body, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + name +
		"<" + from + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n")
	msg = BytesCombine(msg, body)
	error := smtp.SendMail(host+":"+port, auth, from, to, msg)
	if error != nil {
		return fmt.Sprintf("%v", error)
	} else {
		return ""
	}
}

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
