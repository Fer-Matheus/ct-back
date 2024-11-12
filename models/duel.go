package models

type Result struct {
	Id           int `gorm:"primaryKey;autoIncrement"`
	ChosenOption string
	ChoiseTime   string

	DuelId int
	Duel   Duel `gorm:"foreignKey:DuelId;constrait:onDelete:CASCADE"`
}

type Duel struct {
	Id               int  `gorm:"primaryKey;autoIncrement:true"`
	CommitMessageAId int  `gorm:"uniqueIndex:user_duel"`
	CommitMessageBId int  `gorm:"uniqueIndex:user_duel"`
	UserId           int  `gorm:"uniqueIndex:user_duel"`
	User             User `gorm:"foreignKey:UserId;constrait:onDelete:CASCADE"`
	Results          []Result

	CommitMessageA CommitMessage `gorm:"foreignKey:CommitMessageAId;constrait:onDelete:CASCADE"`
	CommitMessageB CommitMessage `gorm:"foreignKey:CommitMessageBId;constrait:onDelete:CASCADE"`
}
