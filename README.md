# mail
邮件相关
## 引用示例

```
package main
 
import (
	"fmt"
	"github.com/w497525437/mail"
)

func main() {
    
    result := mail.SendMail(from, password, host, port, name, subject, filepath, to)
    if result == "" {
      fmt.Printf("success")
    }else {
      fmt.Printf(result)
    }
}
``` 

## Fun
```
SendMail

param
 * from     发送账户 
 * password 密码
 * host     SMTP服务器
 * port     端口号
 * name     发送人
 * subject  标题
 * filepath 模版文件路径
 * to       收件人邮箱
return	    成功返回空字符串，失败返回错误信息
 
 func SendMail(from string, password string, host string, port string, name string, subject string, filepath string, to []string) string {}
```
