package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Plan struct {
	Id                         int    `orm:"column(id);auto"`
	Uid                        string `orm:"column(uid);size(45);null"`
	InstitucionId              int    `orm:"column(institucion_id)"`
	CodigoDane                 int64  `orm:"column(codigo_dane);null"`
	Institucion                string `orm:"column(institucion);size(500);null"`
	RangoMenu                  string `orm:"column(rango_menu);size(45);null"`
	MenuInicial                int    `orm:"column(menu_inicial);null"`
	MenuFinal                  int    `orm:"column(menu_final);null"`
	Anio                       int    `orm:"column(anio)"`
	Semana                     int    `orm:"column(semana)"`
	Direccion                  string `orm:"column(direccion);size(500);null"`
	FechaEntrega               string `orm:"column(fecha_entrega);size(10);null"`
	TipoMinutaId               int    `orm:"column(tipo_minuta_id);null"`
	Abreviacion                string `orm:"column(abreviacion);size(45);null"`
	TipoMinuta                 string `orm:"column(tipo_minuta);size(45);null"`
	DiasAtencion               int    `orm:"column(dias_atencion);null"`
	FechasGeneradas            string `orm:"column(fechas_generadas);size(45);null"`
	MunicipioId                int    `orm:"column(municipio_id);null"`
	MunicipioCodigoDivipola    string `orm:"column(municipio_codigo_divipola);size(45);null"`
	Municipio                  string `orm:"column(municipio);size(45);null"`
	DepartamentoId             int    `orm:"column(departamento_id);null"`
	Departamento               string `orm:"column(departamento);size(45);null"`
	DepartamentoCodigoDivipola string `orm:"column(departamento_codigo_divipola);size(45);null"`
	Estado                     string `orm:"column(estado);size(45);null"`
}

func (t *Plan) TableName() string {
	return "plan"
}

func init() {
	orm.RegisterModel(new(Plan))
}

// AddPlan insert a new Plan into database and returns
// last inserted Id on success.
func AddPlan(m *Plan) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPlanById retrieves Plan by Id. Returns error if
// Id doesn't exist
func GetPlanById(id int) (v *Plan, err error) {
	o := orm.NewOrm()
	v = &Plan{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPlan retrieves all Plan matches certain condition. Returns empty list if
// no records exist
func GetAllPlan(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Plan)).RelatedSel(2)
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

	var l []Plan
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

// UpdatePlan updates Plan by Id and returns error if
// the record to be updated doesn't exist
func UpdatePlanById(m *Plan) (err error) {
	o := orm.NewOrm()
	v := Plan{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePlan deletes Plan by Id and returns error if
// the record to be deleted doesn't exist
func DeletePlan(id int) (err error) {
	o := orm.NewOrm()
	v := Plan{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Plan{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
