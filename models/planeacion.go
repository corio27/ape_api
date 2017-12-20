package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Planeacion struct {
	Id            int       `orm:"column(id);auto"`
	InstitucionId int       `orm:"column(institucion_id);null"`
	Institucion   string    `orm:"column(institucion);size(100);null"`
	RangoId       int       `orm:"column(rango_id)"`
	Rango         string    `orm:"column(rango);size(100);null"`
	MenuId        int       `orm:"column(menu_id)"`
	Menu          string    `orm:"column(menu);size(45);null"`
	ComplementoId int       `orm:"column(complemento_id)"`
	Complemento   string    `orm:"column(complemento);size(100);null"`
	PreparacionId int       `orm:"column(preparacion_id)"`
	Preparacion   string    `orm:"column(preparacion);size(100);null"`
	ProductoId    int       `orm:"column(producto_id)"`
	Producto      string    `orm:"column(producto);size(100);null"`
	Cantidad      int       `orm:"column(cantidad);null"`
	PesoBruto     float64   `orm:"column(peso_bruto);null"`
	Total         float64   `orm:"column(total);null"`
	RangoMenu     string    `orm:"column(rango_menu);size(50);null"`
	Fecha         time.Time `orm:"column(fecha);type(datetime);null"`
}

func (t *Planeacion) TableName() string {
	return "planeacion"
}

func init() {
	orm.RegisterModel(new(Planeacion))
}

// AddPlaneacion insert a new Planeacion into database and returns
// last inserted Id on success.
func AddPlaneacion(m *Planeacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPlaneacionById retrieves Planeacion by Id. Returns error if
// Id doesn't exist
func GetPlaneacionById(id int) (v *Planeacion, err error) {
	o := orm.NewOrm()
	v = &Planeacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPlaneacion retrieves all Planeacion matches certain condition. Returns empty list if
// no records exist
func GetAllPlaneacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Planeacion))
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

	var l []Planeacion
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

// UpdatePlaneacion updates Planeacion by Id and returns error if
// the record to be updated doesn't exist
func UpdatePlaneacionById(m *Planeacion) (err error) {
	o := orm.NewOrm()
	v := Planeacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePlaneacion deletes Planeacion by Id and returns error if
// the record to be deleted doesn't exist
func DeletePlaneacion(id int) (err error) {
	o := orm.NewOrm()
	v := Planeacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Planeacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
