package usecase

import (
	"time"
)

// UseCase is an interface for notice
type UseCase interface {
	CreateTable(userID string, tableID string)
	DeleteTable(userID string, tableID string)
	CopyTable(userID string, tableID string, sourceID string)

	AppendNotice(userID string, tableID string, contents string, time time.Time)
	RemoveNotice(userID string, tableID string, noticeID int)

	Repeat(userID string, tableID string, noticeID int, interval time.Duration)

	UserOn(userID string)
	TableOn(userID string, tableID string)
	NoticeOn(userID string, tableID string, noticeID int)

	UserOf(userID string)
	TableOff(userID string, tableID string)
	NoticeOff(userID string, tableID string, noticeID int)

	TableList(userID string)
	NoticeList(userID string, tableID string)
	Info(userID string, tableID string, noticeID int)
}
