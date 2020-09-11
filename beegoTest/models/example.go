// @BeeOverwrite YES
// @BeeGenerateTime 20200911_220506
package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"time"
)

type Example struct {
	Id   int       ` orm:"auto"json:"id" form:"id"`               // ID
	Row1 int       ` json:"row1" form:"row1"`                     // 字段一
	Row2 string    ` orm:"size(255)"json:"row2" form:"row2"`      // 字段二
	Row3 string    ` orm:"type(longtext)"json:"row3" form:"row3"` // 字段三
	Row4 time.Time ` orm:"type(datetime)"json:"row4" form:"row4"` // 字段四

}

func (t *Example) TableName() string {
	return "example"
}

func init() {
	orm.RegisterModel(new(Example))
}

// AddExample insert a new Example into database and returns
// last inserted Id on success.
func AddExample(m *Example) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(m)
	return
}

// GetExampleById retrieves Example by Id. Returns error if
// Id doesn't exist
func GetExampleById(id int) (v *Example, err error) {
	o := orm.NewOrm()
	v = &Example{
		Id: id,
	}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllExampleretrieves all Example matches certain condition. Returns empty list if
// no records exist
func GetAllExample(query map[string]string, fields []string, sortby []string, order []string,
	currentPage int, pageSize int) (ml ListData, err error) {
	page := NewPagination(currentPage, pageSize)
	ml = ListData{
		List: make([]interface{}, 0),
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(Example))
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
					err = errors.New("Error: Invalid order. Must be either [asc|desc]")
					return
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
					err = errors.New("Error: Invalid order. Must be either [asc|desc]")
					return
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			err = errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
			return
		}
	} else {
		if len(order) != 0 {
			err = errors.New("Error: unused 'order' fields")
			return
		}
	}

	cnt, err := qs.Count()
	if err != nil {
		return
	}
	page.Total = int(cnt)

	var l []Example
	qs = qs.OrderBy(sortFields...)
	_, err = qs.Limit(page.PageSize, (page.Current-1)*page.PageSize).All(&l, fields...)
	if err != nil {
		return
	}

	if len(fields) == 0 {
		for _, v := range l {
			ml.List = append(ml.List, v)
		}
	} else {
		// trim unused fields
		for _, v := range l {
			m := make(map[string]interface{})
			val := reflect.ValueOf(v)
			for _, fname := range fields {
				m[fname] = val.FieldByName(fname).Interface()
			}
			ml.List = append(ml.List, m)
		}
	}
	ml.Pagination = *page
	return
}

// UpdateExample updates Example by Id and returns error if
// the record to be updated doesn't exist
func UpdateExampleById(m *Example) (err error) {
	o := orm.NewOrm()
	v := Example{
		Id: m.Id,
	}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteExample deletes Example by Id and returns error if
// the record to be deleted doesn't exist
func DeleteExample(id int) (err error) {
	o := orm.NewOrm()
	v := Example{
		Id: id,
	}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Example{
			Id: id,
		}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
