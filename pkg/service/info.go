package service

import (
	"fmt"
	"time"
)

type Info struct {
	Id       string
	Name     string
	Version  string
	Metadata map[string]string
}

func NewServiceInfo(name, version, id string) Info {
	if id == "" {
		id = fmt.Sprint(time.Now().Unix())
	}
	return Info{
		Name:     name,
		Version:  version,
		Id:       id,
		Metadata: map[string]string{},
	}
}

func (this *Info) GetInstanceId() string {
	return this.Name + "." + this.Id
}

func (this *Info) SetMataData(k, v string) {
	this.Metadata[k] = v
}
