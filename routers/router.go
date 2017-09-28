// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"pae_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/complemento",
			beego.NSInclude(
				&controllers.ComplementoController{},
			),
		),

		beego.NSNamespace("/complemento_institucion",
			beego.NSInclude(
				&controllers.ComplementoInstitucionController{},
			),
		),

		beego.NSNamespace("/departamento",
			beego.NSInclude(
				&controllers.DepartamentoController{},
			),
		),

		beego.NSNamespace("/etc",
			beego.NSInclude(
				&controllers.EtcController{},
			),
		),

		beego.NSNamespace("/institucion",
			beego.NSInclude(
				&controllers.InstitucionController{},
			),
		),

		beego.NSNamespace("/jornada",
			beego.NSInclude(
				&controllers.JornadaController{},
			),
		),

		beego.NSNamespace("/municipio",
			beego.NSInclude(
				&controllers.MunicipioController{},
			),
		),

		beego.NSNamespace("/persona",
			beego.NSInclude(
				&controllers.PersonaController{},
			),
		),

		beego.NSNamespace("/rango_edad",
			beego.NSInclude(
				&controllers.RangoEdadController{},
			),
		),

		beego.NSNamespace("/tipo_institucion",
			beego.NSInclude(
				&controllers.TipoInstitucionController{},
			),
		),

		beego.NSNamespace("/tipo_minuta",
			beego.NSInclude(
				&controllers.TipoMinutaController{},
			),
		),

		beego.NSNamespace("/tipo_minuta_tipo_modalidad",
			beego.NSInclude(
				&controllers.TipoMinutaTipoModalidadController{},
			),
		),

		beego.NSNamespace("/tipo_modalidad",
			beego.NSInclude(
				&controllers.TipoModalidadController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
