package acl

import (
	"strings"
	//"fmt"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	//"github.com/Rivalz-ai/framework-be/framework/log"
)

// white list route
var whiteList = []string{
	"/user/sign-in",
}

/*
role id 1: admin
role id 10: user
*/
var acl = map[string]string{
	//<METHOD>_<ROUTE_URL>:<ROLE_ID>
	//user
	"GET_/api/v2/user/me": "user",
	//reward
	"GET_/api/v2/reward/rivalz-testnet": "user",
	//pag
	"POST_/api/v2/pag/create": "user",
}

func CheckACL(roles []string, method, route string) bool {
	//temp return true
	return true
	route = strings.TrimSuffix(route, "/")
	//if role is admin then allow all
	for i := 0; i < len(roles); i++ {
		if roles[i] == "admin" {
			return true
		}
	}
	//if route is in white list then allow
	if utils.Contains(whiteList, route) {
		return true
	}
	//
	route_key := method + "_" + route
	val, ok := acl[route_key]
	if ok == false {
		return false
	}
	//check if role is in acl
	if utils.Contains(roles, val) {
		return true
	}
	return false
}
