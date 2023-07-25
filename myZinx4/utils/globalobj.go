package utils

import (
	"encoding/json"
	"myZinx4/ziface"
	"os"
)

/*
*
存储全局参数，
*/
type GlobalObj struct {
	TcpServer      ziface.IServer
	Host           string
	TcpPort        int
	Name           string
	Version        string
	MaxConn        int
	MaxPackageSize uint32
}

func (g *GlobalObj) Reload() {
	data, err := os.ReadFile("../myDemo/ZinxV0.4/config/zinx.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &GroubleObject)

	if err != nil {
		panic(err)
	}
}

// 定义一个全局的对外对象
var GroubleObject *GlobalObj

func init() {
	GroubleObject = &GlobalObj{
		Name:           "Dustdjj",
		Version:        "V0.4",
		TcpPort:        8369,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	GroubleObject.Reload()
}
