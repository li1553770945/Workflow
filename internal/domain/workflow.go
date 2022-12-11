package domain

type WorkflowTemplateEntity struct {
	ID           uint
	Name         string
	SolverUserId string
}

type WorkflowEntity struct {
	ID         uint
	TemplateID uint
	Template   *WorkflowTemplateEntity
	Content    string
	OpenUserId string
}
