package workflow

import (
	"workflow_http/internal/repo"
)

type Service struct {
	Repo repo.IRepository
}

func NewWorkFlow(repo repo.IRepository) IWorkflow {
	return &Service{
		Repo: repo,
	}
}
