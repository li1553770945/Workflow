package domain

type GroupEntity struct {
	ID          uint
	WorkflowID  uint
	Workflow    *WorkflowEntity
	Name        string
	GroupNumber string
}

type GroupMessageEntity struct {
	ID             uint
	Message        string
	SenderNickName string
	SenderId       string
	GroupID        uint
	Group          *GroupEntity
}
