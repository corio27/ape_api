package routers

import (
	"github.com/astaxie/beego"
<<<<<<< HEAD
	"github.com/astaxie/beego/context/param"
=======
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
)

func init() {

<<<<<<< HEAD
	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoController"],
=======
	beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComplementoInstitucionController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponenteController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponenteController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ComponentesMenuController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:ComplementoInstitucionController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:DepartamentoController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EtcController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EtcController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:EtcController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["pae_api/controllers:DepartamentoController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:InstitucionController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["pae_api/controllers:EtcController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["pae_api/controllers:EtcController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["pae_api/controllers:EtcController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["pae_api/controllers:EtcController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:JornadaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:JornadaController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:EtcController"] = append(beego.GlobalControllerRouter["pae_api/controllers:EtcController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MenuController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MenuController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipioController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:MunicipioController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:InstitucionController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PersonaController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:JornadaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:JornadaController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:JornadaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:JornadaController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:JornadaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:JornadaController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:JornadaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:JornadaController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:JornadaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:JornadaController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionesComponenteController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:PreparacionesComponenteController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"] = append(beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"] = append(beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"] = append(beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"] = append(beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductoController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"] = append(beego.GlobalControllerRouter["pae_api/controllers:MunicipioController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosPreparacionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosPreparacionController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:PersonaController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:PersonaController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:PersonaController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:PersonaController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProductosProveedorController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:PersonaController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProveedorController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:ProveedorController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangoEdadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RangoEdadController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:RangoEdadController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolPersonaInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:RolPersonaInstitucionController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoInstitucionController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoAlimentoController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoAlimentoController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoInstitucionController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoInstitucionController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
<<<<<<< HEAD
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
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
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

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaTipoModalidadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaTipoModalidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaTipoModalidadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaTipoModalidadController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaTipoModalidadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaTipoModalidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaTipoModalidadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaTipoModalidadController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaTipoModalidadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoMinutaTipoModalidadController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoMinutaTipoModalidadController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoModalidadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoModalidadController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoModalidadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoModalidadController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
<<<<<<< HEAD
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
=======
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
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
<<<<<<< HEAD
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoModalidadController"] = append(beego.GlobalControllerRouter["github.com/corio27/pae_api/controllers:TipoModalidadController"],
=======
			Params: nil})

	beego.GlobalControllerRouter["pae_api/controllers:TipoModalidadController"] = append(beego.GlobalControllerRouter["pae_api/controllers:TipoModalidadController"],
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
<<<<<<< HEAD
			MethodParams: param.Make(),
=======
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
			Params: nil})

}
