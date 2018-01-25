package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type WaybillTemporal struct {
	Id            int    `orm:"column(id);pk"`
	InstitucionId int    `orm:"column(institucion_id);null"`
	Institucion   string `orm:"column(institucion);size(100);null"`
	Indicaciones  string `orm:"column(indicaciones);size(255);null"`
	ComplementoId int    `orm:"column(complemento_id)"`
	Complemento   string `orm:"column(complemento);size(100);null"`
	Cantidad      int    `orm:"column(cantidad);null"`
	Producto      string `orm:"column(producto);size(100);null"`
	ProductoId    int    `orm:"column(producto_id)"`
	Total         int    `orm:"column(total);null"`
	OrigenId      int    `orm:"column(origen_id)"`
	Origen        string `orm:"column(origen);size(100);null"`
	TipoProducto  string `orm:"column(tipo_producto);size(100);null"`
	UnidadId      int    `orm:"column(unidad_id)"`
	Presentacion  string `orm:"column(presentacion);size(45);null"`
	Peso          int    `orm:"column(peso);null"`
	Unidad        int64  `orm:"column(unidad);null"`
	Sobrante      int    `orm:"column(sobrante);null"`
	Resta         int    `orm:"column(resta);null"`
	Estado        int    `orm:"column(estado);null"`
	RangoMenu     string `orm:"column(rango_menu);size(50);null"`
	CodigoDane    int64  `orm:"column(codigo_dane);null"`
	Anio          int    `orm:"column(anio)"`
	Semana        int    `orm:"column(semana)"`
	Uid           string `orm:"column(uid);size(45);null"`
}

func (t *WaybillTemporal) TableName() string {
	return "waybill_temporal"
}

func init() {
	orm.RegisterModel(new(WaybillTemporal))
}

// AddWaybillTemporal insert a new WaybillTemporal into database and returns
// last inserted Id on success.
func AddWaybillTemporal(m *WaybillTemporal) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetWaybillTemporalById retrieves WaybillTemporal by Id. Returns error if
// Id doesn't exist
func GetWaybillTemporalById(id int) (v *WaybillTemporal, err error) {
	o := orm.NewOrm()
	v = &WaybillTemporal{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllWaybillTemporal retrieves all WaybillTemporal matches certain condition. Returns empty list if
// no records exist
func GetAllWaybillTemporal(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(WaybillTemporal)).RelatedSel(2)
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []WaybillTemporal
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateWaybillTemporal updates WaybillTemporal by Id and returns error if
// the record to be updated doesn't exist
func UpdateWaybillTemporalById(m *WaybillTemporal) (err error) {
	o := orm.NewOrm()
	v := WaybillTemporal{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteWaybillTemporal deletes WaybillTemporal by Id and returns error if
// the record to be deleted doesn't exist
func DeleteWaybillTemporal(id int) (err error) {
	o := orm.NewOrm()
	v := WaybillTemporal{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&WaybillTemporal{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
