package entity

import "gorm.io/gorm"

type StageWord struct {
	gorm.Model
	Tag           string `json:"tag" binding:"required" gorm:"type:VARCHAR(20);uniqueIndex:unique_tag_lang;"`
	Lang          string `json:"lang" binding:"required" gorm:"type:VARCHAR(10);uniqueIndex:unique_tag_lang;"`
	Word          string `json:"word" binding:"required" gorm:"type:text"`
	AdviseUser_ID uint   `json:"adviseUserID"  gorm:"type:uint;comment:由 iadmin 傳來的id, 可能沒人 advice"`
	Organize_ID   uint   `json:"organizeID"`
}

type CommittedWord struct {
	StageWord
	CommitUser_ID uint `json:"commitUserID" gorm:"type:uint;comment:由 iadmin 傳來的id, 只要有人 create 或是 確認 commit 就算是這邊"`
}

type SearchCondition_StageWord struct {
	StageWord
	Pagination
}

type SearchCondition_CommittedWord struct {
	CommittedWord
	Pagination
}
