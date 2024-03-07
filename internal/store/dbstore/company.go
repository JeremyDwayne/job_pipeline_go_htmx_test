package dbstore

import (
	"errors"
	"job_pipeline/internal/store"
)

type CompanyStore struct {
	companies []store.Company
}

func NewCompanyStore() *CompanyStore {
	return &CompanyStore{
		companies: []store.Company{
			// set some default companys
			{
				Name: "Company1",
				Url:  "https://www.company1.com",
			},
		},
	}
}

func (s *CompanyStore) CreateCompany(name string, url string) error {
	for _, company := range s.companies {
		if company.Name == name {
			return errors.New("company already exists")
		}
	}

	s.companies = append(s.companies, store.Company{Name: name, Url: url})
	return nil
}

func (s *CompanyStore) GetCompany(name string) (*store.Company, error) {
	for _, company := range s.companies {
		if company.Name == name {
			return &company, nil
		}
	}

	return nil, errors.New("company not found")
}
