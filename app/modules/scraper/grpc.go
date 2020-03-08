package scraper

import (
	"context"
)

type Scraper struct {}

func (s *Scraper) GetProjects(ctx context.Context, req *GetProjectsRequest) (*GetProjectsResponse, error) {
	projects, err := getProjects()
	if err != nil {
		return nil, err
	}
	return &GetProjectsResponse{Name:projects}, nil
}

func (s *Scraper) GetSpiders(ctx context.Context, req *GetSpidersRequest) (*GetSpidersResponse, error) {
	return &GetSpidersResponse{}, nil
}

func (s *Scraper) RunSpiders(ctx context.Context, req *RunSpiderRequest) (*RunSpiderResponse, error) {
	return &RunSpiderResponse{}, nil
}

func (s *Scraper) ScheduleSpider(ctx context.Context, req *ScheduleSpiderRequest) (*ScheduleSpiderResponse, error) {
	return &ScheduleSpiderResponse{}, nil
}

func (s *Scraper) CancelScheduleSpiders(ctx context.Context, req *CancelScheduleSpiderRequest) (*CancelScheduleSpiderResponse, error) {
	return &CancelScheduleSpiderResponse{}, nil
}

func (s *Scraper) GetJobs(ctx context.Context, req *GetJobsRequest) (*GetJobsResponse, error) {
	return &GetJobsResponse{}, nil
}

func (s *Scraper) GetJob(ctx context.Context, req *GetJobRequest) (*GetJobResponse, error) {
	return &GetJobResponse{}, nil
}

func (s *Scraper) CancelJob(ctx context.Context, req *CancelJobRequest) (*CancelJobResponse, error) {
	return &CancelJobResponse{}, nil
}
