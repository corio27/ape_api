package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Institucion struct {
	Id                    int       `orm:"column(id);auto"`
	Nombre                string    `orm:"column(nombre);size(100);null"`
	CodigoDane            int64     `orm:"column(codigo_dane);null"`
	TipoInstitucionId     int       `orm:"column(tipo_institucion_id);null"`
	MunicipioId           int       `orm:"column(municipio_id);null"`
	EtcId                 int64     `orm:"column(etc_id);null"`
	CantidadManipuladoras int       `orm:"column(cantidad_manipuladoras);null"`
	EsPrincipal           int8      `orm:"column(es_principal);null"`
	CodigoDanePrincipal   int64     `orm:"column(codigo_dane_principal);null"`
	Descripcion           string    `orm:"column(descripcion);size(45);null"`
	TipoMinuta            int       `orm:"column(tipo_minuta);null"`
	TipoModalidad         int       `orm:"column(tipo_modalidad);null"`
	Latitud               float64   `orm:"column(latitud);null"`
	Longitud              float64   `orm:"column(longitud);null"`
	Indicaciones          string    `orm:"column(indicaciones);size(255);null"`
	Estado                int       `orm:"column(estado);null"`
	FechaRegistro         time.Time `orm:"column(fecha_registro);type(date);null"`
	FechaActualizacion    time.Time `orm:"column(fecha_actualizacion);type(date);null"`
}

func (t *Institucion) TableName() string {
	return "institucion"
}

func init() {
	orm.RegisterModel(new(Institucion))
}

// AddInstitucion insert a new Institucion into database and returns
// last inserted Id on success.
func AddInstitucion(m *Institucion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetInstitucionById retrieves Institucion by Id. Returns error if
// Id doesn't exist
func GetInstitucionById(id int) (v *Institucion, err error) {
	o := orm.NewOrm()
	v = &Institucion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllInstitucion retrieves all Institucion matches certain condition. Returns empty list if
// no records exist
func GetAllInstitucion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Institucion)).RelatedSel(2)
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

	var l []Institucion
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

// UpdateInstitucion updates Institucion by Id and returns error if
// the record to be updated doesn't exist
func UpdateInstitucionById(m *Institucion) (err error) {
	o := orm.NewOrm()
	v := Institucion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteInstitucion deletes Institucion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInstitucion(id int) (err error) {
	o := orm.NewOrm()
	v := Institucion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Institucion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
