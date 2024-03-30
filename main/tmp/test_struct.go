package main

import "fmt"

type LifeAuthEntrance struct {
	Show bool `form:"show,required" json:"show" query:"show,required"`

	// key是认证的二级类型
	JumpUrlMap map[int32]string `form:"jump_url_map,required" json:"jump_url_map" query:"jump_url_map,required"`
}

func main() {
	lifeAuthEntrance := LifeAuthEntrance{}

	lifeAuthEntrance.JumpUrlMap[1] = "aa"

	fmt.Println(lifeAuthEntrance)
}
