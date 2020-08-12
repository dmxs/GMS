package service

import "GMS/srv/user/dao"

//Service .
type Service struct {
	Dao dao.Interface
}

//New
func New(daoIns dao.Interface) (service *Service) {
	service = new(Service)
	service.Dao = daoIns
	return
}
