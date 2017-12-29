package utils

import "net/url"

func GetAPIPath(path string) string{
    config:=GetConfig()
    u, _ := url.Parse(config.System.Server)
    u.Path = path
    return u.String()
}
