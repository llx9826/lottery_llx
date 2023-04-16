package dao

import (
	"gorm.io/gorm"
	"log"
	"lottery_llx/model"
)

type GiftDao struct {
	db *gorm.DB
}

func NewGiftDao(db *gorm.DB) *GiftDao {
	return &GiftDao{
		db: db,
	}
}

func (d *GiftDao) Get(id int) *model.LtGift {
	data := &model.LtGift{
		ID: int32(id),
	}
	res := d.db.First(data)
	if res.Error != nil {
		return data
	} else {
		data.ID = 0
		return data
	}
}

func (d *GiftDao) GetAll() []model.LtGift {
	dataList := make([]model.LtGift, 0)
	res := d.db.Order("sys_status ASC").Order("displayorder ASC").Find(&dataList)
	if res == nil {
		log.Println("giftDao GetAll error = ", res.Error)
		return dataList
	}
	return dataList
}

func (d *GiftDao) CountAll() int64 {
	var count int64
	res := d.db.Model(&model.LtGift{}).Count(&count)
	if res != nil {
		log.Println("giftDao CountAll error = ", res.Error)
		return 0
	}
	return count
}

func (d *GiftDao) Delete(id int) error {
	res := d.db.Model(&model.LtGift{}).Where("id = ?", id).Update("sys_status", 1)
	return res.Error
}
func (d *GiftDao) Update(data *model.LtGift, columns []string) error {
	res := d.db.Model(&model.LtGift{}).Where("id = ?", data.ID).Select(columns).Updates(data)
	return res.Error
}

func (d *GiftDao) Create(data *model.LtGift) error {
	res := d.db.Create(data)
	return res.Error
}
