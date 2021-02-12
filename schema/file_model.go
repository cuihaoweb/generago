package schema

import (
	"io/ioutil"
	"os"
)

// ModelFile xx
type ModelFile struct {
	DirName string `json:"dirName"`
}

func (m *ModelFile) getPath(fileName string) string {
	var len = len(m.DirName)

	if m.DirName[len-1] != '/' {
		m.DirName += "/"
	}

	return m.DirName + fileName + ".go"
}

// OutputFile xx
func (m *ModelFile) OutputFile(fileName string, content string) {
	var path = m.getPath(fileName)

	//创建目录
	os.MkdirAll(m.DirName, os.ModePerm)

	if err := ioutil.WriteFile(path, []byte(content), 0666); err != nil {
		panic("写入文件失败\n--------------------------------------------" + err.Error())
	}
}
