package geocoder

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var OrmDB *gorm.DB

func init() {
	var err error
	OrmDB, err = gorm.Open("mysql", "root:root@/go?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Printf("connetion err %s", err)
	}
	defer OrmDB.Close()
}
func (d *Dao) HospitalCount(ctx context.Context) (count int64, err error) {
	//and lng = 0
	err = OrmDB.Model(&Hospital{}).Where(nullAddress).Count(&count).Error
	if err != nil {
		log.Errorc(ctx, "[dao|doctor] HospitalCount err(%v) ", err)
	}
	return
}

// 清理一次
func (d *Dao) GetHospitalIdNeedGeocoder(ctx context.Context) (data []int32, err error) {
	hospitals := make([]*Hospital, 0)
	err = OrmDB.Model(&Hospital{}).Select("id").Where(" lng = 0 ").Where(nullAddress).Find(&hospitals).Error
	if err != nil {
		log.Errorc(ctx, "[dao|doctor] Hospital2 err(%v) ", err)
	}
	// 避免多次扩容
	res := make([]int32,len(hospitals))
	for i := 0; i < len(hospitals); i++ {
		res[i] = hospitals[i].Id
	}
	return res,nil
}

// 找到需要再次更新的对象
func (d *Dao) Hospital2(ctx context.Context,ids []int32) (hospitals []*Hospital, err error) {
	hospitals = make([]*Hospital, 0)
	err = OrmDB.Model(&Hospital{}).Where(" id in (?) ",ids).Find(&hospitals).Error
	if err != nil {
		log.Errorc(ctx, "[dao|doctor] Hospital2 err(%v) ", err)
	}
	return
}

func (d *Dao) HospitalCount2(ctx context.Context) (count int64, err error) {
	err = OrmDB.Model(&Hospital{}).Where(" lng = 0 ").Where( nullAddress).Count(&count).Error
	if err != nil {
		log.Errorc(ctx, "[dao|doctor] HospitalCount2 err(%v) ", err)
	}
	return
}



/*
	UPDATE categories
	SET display_order = CASE id
	WHEN 1 THEN 3
	WHEN 2 THEN 4
	WHEN 3 THEN 5
	END,
	title = CASE id
	WHEN 1 THEN 'New Title 1'
	WHEN 2 THEN 'New Title 2'
	WHEN 3 THEN 'New Title 3'
	END
	WHERE id IN (1,2,3)

	db.Exec("UPDATE orders SET shipped_at=? WHERE id IN (?)", time.Now, []int64{11,22,33})
*/
// 改为批量，否则数据库会报错
func (d *Dao) HospitalUpdateGeocoder(ctx context.Context, hospitals []*Hospital) (err error) {
	sql := " UPDATE hospital SET lat = CASE id "
	l := len(hospitals)
	for i := 0; i < l; i++ {
		sql += fmt.Sprintf(" WHEN %d THEN %f",hospitals[i].Id, hospitals[i].Lat)
	}
	sql += " END, lng = CASE id "
	for i := 0; i < l; i++ {
		sql += fmt.Sprintf(" WHEN %d THEN %f",hospitals[i].Id, hospitals[i].Lng)
	}
	sql += " END WHERE id IN ("
	for i := 0; i < l - 1; i++ {
		sql += fmt.Sprintf("%d,",hospitals[i].Id)
	}
	// 接上最后一个
	sql += fmt.Sprintf("%d)",hospitals[l-1].Id)
	err = OrmDB.Model(&Hospital{}).Exec(sql).Error
	//log.Info("sql is: %s", sql)
	if err != nil {
		log.Errorc(ctx, "[dao|doctor] HospitalUpdateGeocoder err(%v) ", err)
	}
	return
}

