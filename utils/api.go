package utils

import "net/url"

func GetAPIPath(path string) string{
    config:=GetConfig()
    u, _ := url.Parse(config.Server)
    u.Path = path
    return u.String()
}
