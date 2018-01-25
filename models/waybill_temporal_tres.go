package models

type WaybillTemporalTres struct {
	Id_RENAME     int    `orm:"column(id)"`
	InstitucionId int    `orm:"column(institucion_id);null"`
	ComplementoId int    `orm:"column(complemento_id)"`
	ProductoId    int    `orm:"column(producto_id)"`
	CodigoDane    int64  `orm:"column(codigo_dane);null"`
	Anio          int    `orm:"column(anio)"`
	Semana        int    `orm:"column(semana)"`
	Uid           string `orm:"column(uid);size(45);null"`
	Unidad        int64  `orm:"column(unidad);null"`
	Sobrante      int    `orm:"column(sobrante);null"`
	UnidadDos     int64  `orm:"column(unidad_dos);null"`
	SobranteDos   int    `orm:"column(sobrante_dos);null"`
	Peso          int    `orm:"column(peso);null"`
	PesoDos       int    `orm:"column(peso_dos);null"`
	IdDos         int    `orm:"column(id_dos)"`
}
