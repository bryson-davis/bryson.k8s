package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:ObjectController"] = append(beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:ObjectController"] = append(beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:ObjectController"] = append(beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:ObjectController"] = append(beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:ObjectController"] = append(beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:RBACController"] = append(beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:RBACController"],
		beego.ControllerComments{
			Method:           "RBACInClusterSelfNamespace",
			Router:           `/inclusterselfnamespace`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:RBACController"] = append(beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:RBACController"],
		beego.ControllerComments{
			Method:           "RBACInClusterOtherNamespace",
			Router:           `/inclusterothernamespace`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:RBACController"] = append(beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:RBACController"],
		beego.ControllerComments{
			Method:           "RBACInClusterNode",
			Router:           `/inclusternode`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:UserController"] = append(beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:UserController"] = append(beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:UserController"] = append(beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:UserController"] = append(beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:UserController"] = append(beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:UserController"] = append(beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:UserController"] = append(beego.GlobalControllerRouter["bryson.k8s/kubernetes-tester/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
