package main

import (
	"strings"
	// "bytes"
	// "encoding/json"
	"flag"
	"fmt"

	// "io"
	"io/ioutil"
	"net/http"
	"os/exec"

	// "strings"

	"gopkg.in/yaml.v2"
)

//一级目录的结构体
type MyYaml struct {
	Version   string   `yaml:"version"`
	DevOwner  string   `yaml:"devOwner"`
	Backends  Backend  `yaml:"backend"`
	Sqls      Sql      `yaml:"sql"`
	Configs   Config   `yaml:"config"`
	Gateways  Gateway  `yaml:"gateway"`
	Frontends Frontend `yaml:"frontend"`
}

//Backend的二级目录
type Backend struct {
	Name string `yaml:"name"`
	Poms Pom    `yaml:"pom"`
	Svns Svn    `yaml:"svn"`
}

//Pom的三级目录
type Pom struct {
	Version string `yaml:"version"`
}

//Svn的三级目录
type Svn struct {
	Branches string `yaml:"branches"`
}

//Sql的二级目录
type Sql struct {
	Databases Database `yaml:"database"`
}

//Database的三级目录
type Database struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

//Config的二级目录
type Config struct {
	Url string `yaml:"url"`
}

//Gateway的二级目录
type Gateway struct {
	Apis Api `yaml:"api"`
}

//Api的三级目录
type Api struct {
	Url string `yaml:"url"`
}

//Frontend的二级目录
type Frontend struct {
	Svns    Svn2   `yaml:"svn"`
	Uadmins Uadmin `yaml:"uadmin"`
}

//Api的三级目录
type Svn2 struct {
	Branches string `yaml:"branches"`
}

//Uadmin的三级目录
type Uadmin struct {
	Modules Module `yaml:"module"`
}

//Uadmin的三级目录
type Module struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type HeartbeatResponse struct {
	API    string `json:"api"`
	CODE   string `json:"code"`
	MSG    string `json:"msg"`
	RESULT string `json:"result"`
}

var (
	automation string
)

func main() {

	flag.StringVar(&automation, "AOTOMATION", "conf.yaml", "automation")

	data, _ := ioutil.ReadFile(automation)
	// fmt.Println(string(data))
	t := MyYaml{}
	//把yaml形式的字符串解析成struct类型
	yaml.Unmarshal(data, &t)
	if t.Version == "" {
		fmt.Println("配置文件不对")
		return
	}
	// fmt.Println(t.Version)

	myCmd := "http:///"
	str3 := `"create branch"`
	str := "svn cp -m " + str3 + "  " + t.Backends.Svns.Branches + "  " + myCmd
	fmt.Println(str)

	cmd := exec.Command("/bin/ls")
	// 执行命令，并返回结果
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	// 因为结果是字节数组，需要转换成string
	fmt.Println(string(output))

	jenkinsUrl := ""
	httpPost(jenkinsUrl)
}

func httpPost(url string) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader("SVN_SQL_URL=cjb&SVN=sdlk"))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}
