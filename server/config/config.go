package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"sync"
)
var (
	Conf = New()
	cfg *toml.Tree
	once sync.Once
)


func New() *toml.Tree  {
	once.Do(func() {
		var err error
		cfg,err=toml.LoadFile("./config/config.toml")
		if err!=nil{
			fmt.Println("TomlError",err.Error())
		}
	})
	return cfg
}

