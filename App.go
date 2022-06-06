package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"tools/config"
)

//Notion的Crypto每日反省添加Message

func AddCryptoMessage(c *gin.Context) {

	jsonData := CryptoMessage{}
	err := c.BindJSON(&jsonData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mess := jsonData.Mess
	url := fmt.Sprintf("https://api.notion.com/v1/blocks/%s/children", config.Settings.Notion.CryptoBlockId)

	mapdata := make(map[string]interface{})
	mapdata["children"] = []interface{}{
		map[string]interface{}{
			"object": "block",
			"type":   "paragraph",
			"paragraph": map[string]interface{}{
				"rich_text": []interface{}{
					map[string]interface{}{
						"type": "text",
						"text": map[string]interface{}{
							"content": mess,
							"link":    nil,
						},
					},
				},
			},
		},
	}
	b, _ := json.Marshal(mapdata)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(b))

	req.Header.Add("Authorization", fmt.Sprintf("%s", config.Settings.Notion.Auth))
	req.Header.Add("Notion-Version", fmt.Sprintf("%s", config.Settings.Notion.Version))
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	c.JSON(http.StatusOK, gin.H{"message": string(body)})
}
