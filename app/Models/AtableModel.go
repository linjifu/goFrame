package Models

type AtableModel struct {
	ID   uint   `gorm:"primaryKey;autoIncrement:true;column:id;type:int(11) unsigned AUTO_INCREMENT;not null" json:"id"`
	Name string `gorm:"column:name;type:varchar(255);not null;default:'';comment:名称" json:"name"`
}

func (a *AtableModel) TableName() string {
	return "a"
}
