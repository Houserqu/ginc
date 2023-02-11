package ginc

import (
	"encoding/json"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

// 发送系统告警消息
func SendAlertMessage(title string, message string) (err error) {
	hookurl := viper.GetString("alert.feishu")
	if hookurl == "" {
		log.Println(title, message)
		return
	}

	params := map[string]any{
		"msg_type": "post",
		"content": map[string]any{
			"post": map[string]any{
				"zh_cn": map[string]any{
					"title": title,
					"content": [][]map[string]any{
						{
							{
								"tag":  "text",
								"text": message,
							},
						},
					},
				},
			},
		},
	}
	str, _ := json.Marshal(params)

	client := resty.New()
	_, err = client.R().SetBody(str).Post(hookurl)
	return
}
