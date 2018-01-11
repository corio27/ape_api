package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Waybill struct {
	Id               int       `orm:"column(id);pk"`
	Alimento         string    `orm:"column(alimento);size(45);null"`
	ProductoId       int       `orm:"column(producto_id);null"`
	Peso             int       `orm:"column(peso);null"`
	Presentacion     string    `orm:"column(presentacion);size(45);null"`
	PesoPresentacion int       `orm:"column(peso_presentacion);null"`
	Unidades         int       `orm:"column(unidades);null"`
	Origen           string    `orm:"column(origen);size(45);null"`
	Institucion      string    `orm:"column(institucion);size(255);null"`
	Direccion        string    `orm:"column(direccion);size(45);null"`
	CuposManana      int       `orm:"column(cupos_manana);null"`
	CuposTarde       int       `orm:"column(cupos_tarde);null"`
	Almuerzo         int       `orm:"column(almuerzo);null"`
	Menu             string    `orm:"column(menu);size(50);null"`
	TipoAlimento     string    `orm:"column(tipo_alimento);size(45);null"`
	Observacion      string    `orm:"column(observacion);size(255);null"`
	Semana           int       `orm:"column(semana);null"`
	Fecha            time.Time `orm:"column(fecha);type(date);null"`
	EntregaReal      int       `orm:"column(entrega_real);null"`
	Anio             int       `orm:"column(anio);null"`
	Uid              string    `orm:"column(uid);size(45);null"`
}

func (t *Waybill) TableName() string {
	return "waybill"
}

func init() {
	orm.RegisterModel(new(Waybill))
}

// AddWaybill insert a new Waybill into database and returns
// last inserted Id on success.
func AddWaybill(m *Waybill) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetWaybillById retrieves Waybill by Id. Returns error if
// Id doesn't exist
func GetWaybillById(id int) (v *Waybill, err error) {
	o := orm.NewOrm()
	v = &Waybill{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllWaybill retrieves all Waybill matches certain condition. Returns empty list if
// no records exist
func GetAllWaybill(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Waybill))
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

	var l []Waybill
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

// UpdateWaybill updates Waybill by Id and returns error if
// the record to be updated doesn't exist
func UpdateWaybillById(m *Waybill) (err error) {
	o := orm.NewOrm()
	v := Waybill{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteWaybill deletes Waybill by Id and returns error if
// the record to be deleted doesn't exist
func DeleteWaybill(id int) (err error) {
	o := orm.NewOrm()
	v := Waybill{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Waybill{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
