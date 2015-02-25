package templating

import (
	"github.com/jamesyong/o3erp/go/config"
	"github.com/jamesyong/o3erp/go/helper"
	"github.com/unrolled/render"
	"html/template"
	"log"
	"strings"
)

var (
	Render  *render.Render
	Options = render.Options{}
)

func Setup() {
	if Options.Directory == "" {
		Options.Directory = config.PATH_BASE_GOLANG_TEMPLATES
	}
	Options.Funcs = []template.FuncMap{
		{
			"msg": prepareMessage,
		},
	}
	Render = render.New(Options)
}

func prepareMessage(userLoginId string, args ...interface{}) map[string]string {

	m := make(map[string]string)
	labels := []string{}
	for index := range args {
		s := strings.Split(args[index].(string), ":")
		labels = append(labels, s[1])
		m[s[0]] = s[1]
	}

	labelMap, err := helper.RunThriftService(helper.GetMessageMapFunction(userLoginId, labels))
	if err != nil {
		log.Println("error: ", err)
	}

	for key, value := range m {
		m[key] = labelMap[value]
	}

	return m
}
