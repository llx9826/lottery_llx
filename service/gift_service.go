package service

import (
	"lottery_llx/dao"
	"lottery_llx/datasource"
	"lottery_llx/model"
)

type GiftService interface {
	GetAll(useCache bool) []model.LtGift
	CountAll() int64
	Get(id int) *model.LtGift
	Delete(id int) error
	Update(data *model.LtGift, columns []string) error
	Create(data *model.LtGift) error
}

type giftService struct {
	dao *dao.GiftDao
}

func (s *giftService) GetAll(useCache bool) []model.LtGift {
	return s.dao.GetAll()
}

func (s *giftService) CountAll() int64 {
	//TODO implement me
	panic("implement me")
}

func (s *giftService) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func (s *giftService) Create(data *model.LtGift) error {
	//TODO implement me
	panic("implement me")
}

func NewGiftServie() GiftService {
	return &giftService{
		dao: dao.NewGiftDao(datasource.NewDbMaster()),
	}
}

func (s *giftService) Get(id int) *model.LtGift {
	return s.dao.Get(id)
}

func (s *giftService) Update(data *model.LtGift, columns []string) error {
	return s.dao.Update(data, columns)
}
