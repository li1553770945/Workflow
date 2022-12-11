package dao

type GroupDO struct {
	BaseModel BaseModel `gorm:"embedded"`

	Name        string
	GroupNumber string
	WorkflowID  uint
	Workflow    *WorkflowDO
}

type GroupMessageDO struct {
	BaseModel BaseModel `gorm:"embedded"`

	Message        string
	SenderNickName string
	SenderId       string
	GroupID        uint
	Group          *GroupDO
}

func (dao *DAO) CreateGroup(do *GroupDO) error {
	err := dao.DB.Create(&do).Error
	return err
}
func (dao *DAO) CreateGroupMessage(do *GroupMessageDO) error {
	err := dao.DB.Create(&do).Error
	return err

}
