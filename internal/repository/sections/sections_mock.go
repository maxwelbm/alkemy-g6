package sectionsrp

import (
	"github.com/maxwelbm/alkemy-g6/internal/models"
	"github.com/stretchr/testify/mock"
)

type SectionsMock struct {
	mock.Mock
}

func NewSectionsMock() *SectionsMock {
	return &SectionsMock{}
}

func (s *SectionsMock) GetAll() (sections []models.Section, err error) {
	args := s.Called()
	return args.Get(0).([]models.Section), args.Error(1)
}

func (s *SectionsMock) GetByID(id int) (section models.Section, err error) {
	args := s.Called(id)
	return args.Get(0).(models.Section), args.Error(1)
}

func (s *SectionsMock) GetReportProducts(sectionID int) (reportProducts []models.ProductReport, err error) {
	args := s.Called(sectionID)
	return args.Get(0).([]models.ProductReport), args.Error(1)
}

func (s *SectionsMock) Delete(id int) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *SectionsMock) Create(section models.SectionDTO) (sectionToReturn models.Section, err error) {
	args := s.Called(section)
	return args.Get(0).(models.Section), args.Error(1)
}

func (s *SectionsMock) Update(id int, section models.SectionDTO) (sectionToReturn models.Section, err error) {
	args := s.Called(id, section)
	return args.Get(0).(models.Section), args.Error(1)
}
