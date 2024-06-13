package services

import (
	"github.com/IraIvanishak/stack_app/internal/analyzer"
	"github.com/IraIvanishak/stack_app/internal/repo"
)

type JobService struct {
	repo *repo.JobRepo
}

func NewJobService(repo *repo.JobRepo) *JobService {
	return &JobService{repo: repo}
}

func (service *JobService) FilRaw() error {
	raw, err := service.repo.GetRaw()
	if err != nil {
		return err
	}

	if len(raw) == 0 {
		return nil
	}

	// analyze text
	updated := analyzer.Analyze(raw)

	err = service.repo.SaveRaw(updated)
	if err != nil {
		return err
	}

	return nil
}

func (service *JobService) GetAnalytic(category string) (map[string]int, error) {
	raw, err := service.repo.GetByCategory(category)
	if err != nil {
		return nil, err
	}

	if len(raw) == 0 {
		return nil, nil
	}

	m := make(map[string]int)

	for _, v := range raw {
		for _, stack := range v.RequiredStack {
			m[stack]++
		}
		for _, stack := range v.WelcomeStack {
			m[stack]++
		}
	}
	m = analyzer.Merge(m)
	return m, nil

}
