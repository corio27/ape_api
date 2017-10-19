// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/corio27/pae_api/controllers"

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

		beego.NSNamespace("/componente",
			beego.NSInclude(
				&controllers.ComponenteController{},
			),
		),

		beego.NSNamespace("/componentes_menu",
			beego.NSInclude(
				&controllers.ComponentesMenuController{},
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

		beego.NSNamespace("/menu",
			beego.NSInclude(
				&controllers.MenuController{},
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

		beego.NSNamespace("/preparacion",
			beego.NSInclude(
				&controllers.PreparacionController{},
			),
		),

		beego.NSNamespace("/preparaciones_componente",
			beego.NSInclude(
				&controllers.PreparacionesComponenteController{},
			),
		),

		beego.NSNamespace("/producto",
			beego.NSInclude(
				&controllers.ProductoController{},
			),
		),

		beego.NSNamespace("/productos_preparacion",
			beego.NSInclude(
				&controllers.ProductosPreparacionController{},
			),
		),

		beego.NSNamespace("/productos_proveedor",
			beego.NSInclude(
				&controllers.ProductosProveedorController{},
			),
		),

		beego.NSNamespace("/proveedor",
			beego.NSInclude(
				&controllers.ProveedorController{},
			),
		),

		beego.NSNamespace("/rango_edad",
			beego.NSInclude(
				&controllers.RangoEdadController{},
			),
		),

		beego.NSNamespace("/rol",
			beego.NSInclude(
				&controllers.RolController{},
			),
		),

		beego.NSNamespace("/rol_persona_institucion",
			beego.NSInclude(
				&controllers.RolPersonaInstitucionController{},
			),
		),

		beego.NSNamespace("/tipo_alimento",
			beego.NSInclude(
				&controllers.TipoAlimentoController{},
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
