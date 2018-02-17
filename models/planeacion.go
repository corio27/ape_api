package models

type Planeacion struct {
	InstitucionId  int     `orm:"column(institucion_id)"`
	RangoId        int     `orm:"column(rango_id)"`
	Rango          string  `orm:"column(rango);size(100);null"`
	MenuId         int     `orm:"column(menu_id)"`
	Menu           string  `orm:"column(menu);size(45);null"`
	ComplementoId  int     `orm:"column(complemento_id)"`
	Complemento    string  `orm:"column(complemento);size(100);null"`
	PreparacionId  int     `orm:"column(preparacion_id)"`
	Preparacion    string  `orm:"column(preparacion);size(100);null"`
	ProductoId     int     `orm:"column(producto_id)"`
	Codigo         int     `orm:"column(codigo);null"`
	Producto       string  `orm:"column(producto);size(100);null"`
	Cantidad       int     `orm:"column(cantidad);null"`
	PesoBruto      float64 `orm:"column(peso_bruto);null"`
	Total          float64 `orm:"column(total);null"`
	Anio           int64   `orm:"column(anio)"`
	Semana         int64   `orm:"column(semana)"`
	Proveedor      string  `orm:"column(proveedor);size(100);null"`
	Valor          int     `orm:"column(valor);null"`
	TipoAlimentoId int     `orm:"column(tipo_alimento_id)"`
	TipoAlimento   string  `orm:"column(tipo_alimento);size(100);null"`
	Uid            string  `orm:"column(uid);size(100)"`
}
