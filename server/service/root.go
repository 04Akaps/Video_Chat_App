package service

import "github.com/04Akaps/Video_Chat_App/reposiroty"

type Service struct {
	Mysql *mysql
}

func NewService(db *reposiroty.DB) *Service {
	service := &Service{
		Mysql: newMySqlService(db),
	}

	return service
}
