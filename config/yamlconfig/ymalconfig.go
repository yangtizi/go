package yamlconfig

import (
	"io/ioutil"

	log "github.com/yangtizi/go/log/zaplog"
	"gopkg.in/yaml.v2"
)

// YMALParsing 解析YAML文件
func YMALParsing(strPath string, v interface{}) {
	data, err := ioutil.ReadFile(strPath)

	if err != nil {
		log.Errorf("%v", err)
	}

	yaml.Unmarshal(data, v)

	if err != nil {
		log.Errorf("%v", err)
	}

}
