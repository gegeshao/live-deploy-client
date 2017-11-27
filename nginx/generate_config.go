package nginx

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"live-deploy/database/bean"
	"live-deploy/database/schema"
	"live-deploy/utils"
	"os"
	"regexp"
	"text/template"
	"time"
)

func GenerateNginxConfig(data *schema.NginxAppFull) (string, error) {
	templateConfig, err := utils.GetNginxTemplate()
	if err != nil {
		return "", err
	}
	var templateBody bytes.Buffer
	templ, err := template.New("nginx").Parse(templateConfig)
	if err != nil {
		return "", err
	}
	if err := templ.Execute(&templateBody, data); err != nil {
		return "", err
	}
	body := templateBody.String()
	re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`)
	return re.ReplaceAllString(body, ""), nil
}

func GetPreviewContent(appid int64) (map[string]string, error) {
	app, err := bean.GetPreview(appid)
	if err != nil {
		return nil, err
	}
	content, err := GenerateNginxConfig(&app)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"content": content,
		"name":    app.Name,
	}, nil
}

// UndeployNginxConfig 下线服务， 当配置文件不存在时返回false
func UndeployNginxConfig(appid int64) (bool, error) {
	filename := utils.GetNginxTemplateFileName(appid)
	//
	if _, err := os.Stat(filename); err != nil {
		//文件不存在不需要做任何事
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	if reErr := os.Remove(filename); reErr != nil {
		return false, reErr
	}
	_, err := RestartNginx()
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetRunningConfig(appid int64) ([]byte, error) {
	filename := utils.GetNginxTemplateFileName(appid)
	return ioutil.ReadFile(filename)
}

func DeployNginxConfig(machineMD5ID string, appid int64, userid int64) (string, error) {
	data, err := GetPreviewContent(appid)
	if err != nil {
		return "", err
	}
	filename := utils.GetNginxTemplateFileName(appid)

	if err := ioutil.WriteFile(filename, []byte(data["content"]), 0644); err != nil {
		return "", err
	}
	ensureMsg, err := EnsureNginxConfig()
	if err != nil {
		//移除配置文件
		if reErr := os.Remove(filename); reErr != nil {
			return "", fmt.Errorf(err.Error() + "\n" + reErr.Error())
		}
		return "", err
	}
	reloadMsg, err := RestartNginx()
	if err != nil {
		return "", err
	}
	if reloadMsg == "" {
		reloadMsg = "nginx reload success"
	}

	backup := schema.NginxBackup{
		NginxID:      appid,
		Content:      data["content"],
		CreateTime:   time.Now().Unix(),
		AuthorID:     userid,
		MachineMD5ID: machineMD5ID,
	}
	if err := bean.BackupConfig(&backup); err != nil {
		return ensureMsg + "\n" + reloadMsg + "\n", err
	}
	return ensureMsg + "\n" + reloadMsg, nil
}
