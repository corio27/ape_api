package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoInstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoInstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoInstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoInstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoInstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponenteController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponenteController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponenteController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponenteController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponenteController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuTempController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuTempController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuTempController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuTempController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuTempController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuTempController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuTempController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuTempController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuTempController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuTempController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosInstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosInstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosInstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosInstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosInstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosProveedorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosProveedorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosProveedorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosProveedorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ElementosProveedorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EstadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EstadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EstadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EstadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EstadoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EstadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EtcController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EtcController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EtcController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EtcController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EtcController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:JornadaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:JornadaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:JornadaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:JornadaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:JornadaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:JornadaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:JornadaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:JornadaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:JornadaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:JornadaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MenuController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MenuController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MenuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MenuController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MenuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MenuController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MenuController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipioController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipioController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipioController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipioController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipioController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipiosProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipiosProveedorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipiosProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipiosProveedorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipiosProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipiosProveedorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipiosProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipiosProveedorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipiosProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipiosProveedorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PeriodoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PeriodoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PeriodoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PeriodoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PeriodoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PeriodoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PeriodoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PeriodoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PeriodoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PeriodoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PlanController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PlanController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PlanController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PlanController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PlanController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PlanController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PlanController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PlanController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PlanController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PlanController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionesComponenteController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionesComponenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionesComponenteController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionesComponenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionesComponenteController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionesComponenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionesComponenteController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionesComponenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionesComponenteController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionesComponenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosPreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosPreparacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosPreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosPreparacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosPreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosPreparacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosPreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosPreparacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosPreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosPreparacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosProveedorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosProveedorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosProveedorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosProveedorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosProveedorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProveedorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProveedorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProveedorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProveedorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProveedorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangoEdadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangoEdadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangoEdadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangoEdadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangoEdadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangoEdadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangoEdadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangoEdadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangoEdadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangoEdadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangosPreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangosPreparacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangosPreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangosPreparacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangosPreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangosPreparacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangosPreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangosPreparacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangosPreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangosPreparacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolPersonaInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolPersonaInstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolPersonaInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolPersonaInstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolPersonaInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolPersonaInstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolPersonaInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolPersonaInstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolPersonaInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolPersonaInstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoAlimentoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoAlimentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoAlimentoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoAlimentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoAlimentoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoAlimentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoAlimentoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoAlimentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoAlimentoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoAlimentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoInstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoInstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoInstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoInstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoInstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoModalidadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoModalidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoModalidadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoModalidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoModalidadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoModalidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoModalidadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoModalidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoModalidadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoModalidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TotalPlaneacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TotalPlaneacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TotalPlaneacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TotalPlaneacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TotalPlaneacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TotalPlaneacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TotalPlaneacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TotalPlaneacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TotalPlaneacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TotalPlaneacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:UnidadMedidaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:UnidadMedidaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:UnidadMedidaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:UnidadMedidaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:UnidadMedidaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:UnidadMedidaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:UnidadMedidaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:UnidadMedidaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:UnidadMedidaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:UnidadMedidaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalCuatroController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalCuatroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalCuatroController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalCuatroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalCuatroController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalCuatroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalCuatroController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalCuatroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalCuatroController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalCuatroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalDosController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalDosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalDosController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalDosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalDosController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalDosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalDosController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalDosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalDosController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:WaybillTemporalDosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
