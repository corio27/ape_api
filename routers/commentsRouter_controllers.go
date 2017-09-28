package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["pae_api/controllers:EtcController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["pae_api/controllers:EtcController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["pae_api/controllers:EtcController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["pae_api/controllers:EtcController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["pae_api/controllers:EtcController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:JornadaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:JornadaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:JornadaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:JornadaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:JornadaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:JornadaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:JornadaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:JornadaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:JornadaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:JornadaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"] = append(beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"] = append(beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"] = append(beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"] = append(beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"] = append(beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoModalidadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoModalidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoModalidadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoModalidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoModalidadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoModalidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoModalidadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoModalidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoModalidadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoModalidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
