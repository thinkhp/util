package cache

import (
	"util/thinkJson"
	"util/think"
	"io/ioutil"
	"os"
)

type Cache struct {
	// 磁盘存储位置
	InDisk
	// 内存缓冲位置
	InMemory
}

type InDisk struct {
	FileFullName string
}
type InMemory map[string][]string

// 初始化缓冲
func (c *Cache)Init(filePath, fileName string){
	// disk -> memory
	c.InDisk = InDisk{filePath+fileName}
	c.InMemory = make(map[string][]string)

	bs, err := ioutil.ReadFile(c.InDisk.FileFullName)
	think.IsNil(err)
	jo := thinkJson.MustGetJsonObject(bs)
	for key, _ := range jo {
		for _, v := range jo.MustGetStringList(key) {
			c.InMemory[key] = append(c.InMemory[key], v)
		}
	}
}

func (c *Cache)Get(key string) []string {
	return c.InMemory[key]
}

func (c *Cache)Add(key string, value string){
	list := c.InMemory[key]
	for _, v := range list {
		if v == value {
			return
		}
	}
	list = append(list, value)
	c.InMemory[key] = list
	// memory -> disk
	f, err := os.OpenFile(c.FileFullName, os.O_TRUNC|os.O_RDWR, 0660)
	think.IsNil(err)
	defer f.Close()
	f.Write(thinkJson.MustMarshal(c.InMemory))
}