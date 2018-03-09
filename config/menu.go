package config

import (
	"encoding/json"

	"github.com/qjw/kelly"
)

type MenuItem struct {
	ID       int    `json:"id,omitempty" binding:"required"`
	Redirect string `json:"redirect,omitempty"`
	Name     string `json:"name,omitempty" binding:"required"`
	Icon     string `json:"icon,omitempty" binding:"required"`
	Route    string `json:"route,omitempty"`
	PID      int    `json:"pid,omitempty"`
	// -1表示不在左树显示，仅仅用于适配面包屑菜单
	MPID int `json:"mpid,omitempty"`

	Children []*MenuItem `json:"children,omitempty"`
}

var aaa = `
[
	{
		"route": "/namespaces",
		"name": "命名空间",
		"icon": "user",
		"children": [{
			"mpid": -1,
			"route": "/secrets/:namespace",
			"name": "密码",
			"icon": "user",
			"children": [{
				"mpid": -1,
				"route": "/secrets/:namespace/detail/:id",
				"name": "密码详情",
				"icon": "user"
			}]
		}]
	},
	{
		"route": "/test",
		"name": "测试",
		"icon": "api"
	}
]
`

type idGenerator func() int

func idGen(begin int) idGenerator {
	return func() int {
		begin += 1
		return begin
	}
}

func (this *MenuItem) build(idg idGenerator) {
	this.ID = idg()
	for k, v := range this.Children {
		v.PID = this.ID
		if k == 0 && v.MPID != -1 {
			this.Route = ""
			this.Redirect = v.Route
		}
		v.build(idg)
	}
}

func (this *MenuItem) ser(res []*MenuItem) []*MenuItem {
	res = append(res, this)
	for _, v := range this.Children {
		res = v.ser(res)
	}
	this.Children = make([]*MenuItem, 0)
	return res
}

func GetMenuString() []byte {
	var menus []*MenuItem
	if err := json.Unmarshal([]byte(aaa), &menus); err != nil {
		panic(err)
	}

	idg := idGen(0)
	for _, v := range menus {
		v.build(idg)
	}

	total := idg()
	var res []*MenuItem = make([]*MenuItem, 0, total-1)
	for _, v := range menus {
		res = v.ser(res)
	}

	jsonBytes, err := json.MarshalIndent(kelly.H{
		"code":    0,
		"message": "ok",
		"data":    res,
	}, "", "    ")
	if err != nil {
		panic(err)
	}

	return jsonBytes
}
