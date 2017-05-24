// Copyright 2016 CodisLabs. All Rights Reserved.
// Licensed under the MIT (MIT-LICENSE.txt) license.

package models

type Topom struct {
	Token     string `json:"token"`
	StartTime string `json:"start_time"`
	AdminAddr string `json:"admin_addr"`

	ProductName string `json:"product_name"`

	Pid int    `json:"pid"`
	Pwd string `json:"pwd"`
	Sys string `json:"sys"`
}

func (t *Topom) String() string {
	return JsonString(t)
}

func (t *Topom) Encode() []byte {
	return JsonEncode(t)
}

func (t *Topom) Decode(b []byte) error {
	return JsonDecode(t, b)
}

type TopomConfig struct {
	SlotMappings []*SlotMapping
	GroupMap     map[int]*Group
	ProxyMap     map[string]*Proxy
	Sentinel     *Sentinel
}

func (c *TopomConfig) Encode() []byte {
	return GzipGobEncode(c)
}

func (c *TopomConfig) Decode(b []byte) error {
	return GzipGobDecode(c, b)
}
