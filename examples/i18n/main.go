package main

import (
	"fmt"
	"suzaku/pkg/i18n"
)

func main() {
	i18n.Initialization()
	T := i18n.GetUserTranslations("zh-CN")
	title := T("api.templates.username_change_body.title")
	fmt.Println(title)
}
