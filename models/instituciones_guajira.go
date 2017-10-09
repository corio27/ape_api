package models

type InstitucionesGuajira struct {
	CODIGODANESEDE                              int64  `orm:"column(_CODIGO_DANE_SEDE_);null"`
	NOMBREDELASEDE                              string `orm:"column(_NOMBRE_DE_LA_SEDE_);size(255);null"`
	CODIGODANE                                  int64  `orm:"column(_CODIGO_DANE_);null"`
	NOMBREDELESTABLECIMIENTOEDUCATIVO           string `orm:"column(_NOMBRE_DEL_ESTABLECIMIENTO_EDUCATIVO_);size(255);null"`
	URBANA                                      string `orm:"column(_URBANA_);size(3);null"`
	RURAL                                       string `orm:"column(_RURAL_);size(3);null"`
	ETNICORURAL                                 string `orm:"column(_ETNICO__RURAL_);size(3);null"`
	ETC                                         string `orm:"column(_ETC_);size(12);null"`
	MUNICIPIO                                   string `orm:"column(_MUNICIPIO_);size(50);null"`
	COMPLEMENTOAMActualizado                    int    `orm:"column(_COMPLEMENTO_AM_actualizado_);null"`
	COMPLEMENTOAMAntiguo                        int    `orm:"column(_COMPLEMENTO_AM_antiguo_);null"`
	COMPLEMENTOPMActulizada                     int    `orm:"column(_COMPLEMENTO_PM__Actulizada_);null"`
	COMPLEMENTOPMAntigua                        int    `orm:"column(_COMPLEMENTO_PM_Antigua__);null"`
	ALMUERZOActualizado                         int    `orm:"column(_ALMUERZO_Actualizado_);null"`
	ALMUERZOAntiguo                             int    `orm:"column(_ALMUERZO_Antiguo_);null"`
	JORNADAUNICAActualizada                     int    `orm:"column(_JORNADA_UNICA_Actualizada_);null"`
	JORNADAUNICAAntigua                         int    `orm:"column(_JORNADA_UNICA_Antigua_);null"`
	TOTALDERACIONESDIAPORINSTITUCIONActualizada int    `orm:"column(_TOTAL_DE_RACIONES__DIA_POR_INSTITUCION_Actualizada_);null"`
	NUMERODEMANIPULADORASCOMPLEMENTOAMPM        string `orm:"column(_NUMERO_DE_MANIPULADORAS_COMPLEMENTO_AM_PM_);size(1);null"`
	MODALIADADDEATENCION                        string `orm:"column(_MODALIADAD_DE_ATENCION_);size(50);null"`
	FECHADEINICIODEOPERACION                    string `orm:"column(_FECHA_DE_INICIO_DE_OPERACION___);size(22);null"`
	OBSERVACIONES                               string `orm:"column(_OBSERVACIONES__);size(255);null"`
	NOMBREDELRECTORYODIRECTOR                   string `orm:"column(_NOMBRE_DEL_RECTOR_Y_O_DIRECTOR_);size(51);null"`
	TELEFONO                                    string `orm:"column(_TELEFONO_);size(50);null"`
	CORREOELECTRONICO                           string `orm:"column(_CORREO_ELECTRONICO_);size(100);null"`
	NOMBRECOORDINADORPMAACARGOSEDE              string `orm:"column(_NOMBRE_COORDINADOR_PMA_A_CARGO__SEDE__);size(100);null"`
	NOMBRECOORDINADORFUNDALIANZAACARGODESEDE    string `orm:"column(_NOMBRE_COORDINADOR__FUNDALIANZA_A_CARGO_DE_SEDE___);size(100);null"`
	TIPODEMINUTA                                string `orm:"column(_TIPO_DE_MINUTA__);size(50);null"`
	FECHADEINICIOREAL                           string `orm:"column(_FECHA_DE_INICIO_REAL_);size(20);null"`
}
