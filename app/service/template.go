package service

import "Pg_FW_Template/model"

type templateDAO interface {
	Get(id uint) (*model.TemplateTest, error)
}

type TemplateService struct {
	dao templateDAO
}

// NewTemplateService creates a new TemplateService with the given user DAO.
func NewTemplateService(dao templateDAO) *TemplateService {
	return &TemplateService{dao}
}

// Get just retrieves user using User DAO, here can be additional logic for processing data retrieved by DAOs
func (s *TemplateService) Get(id uint) (*model.TemplateTest, error) {
	return s.dao.Get(id)
}
