package model

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/transition"
)

type TaskBrief struct {
	gorm.Model
	TenantUuid     string
	TaskPeriodId   uint
	Title          string
	UserId         uint
	DepartmentId   uint
	ProcessContent string
	IssueContent   string
	PlanContent    string
	transition.Transition
}

// 状态变化记录
func GetStateChangeLogs() []transition.StateChangeLog {
	var stateChangeLogs = transition.GetStateChangeLogs(&TaskBrief{}, DB)
	return stateChangeLogs
}

// 获取状态机
// 使用方法：
// recordStateMachine := getStateMachine(&brief)
// recordStateMachine.Trigger("send_brief", &brief, db)
func getStateMachine(brief *TaskBrief) *transition.StateMachine {
	var recordStateMachine = transition.New(&brief)

	recordStateMachine.Initial("draft")
	recordStateMachine.State("sent")
	recordStateMachine.State("finished")

	// 发送简报
	recordStateMachine.Event("send_brief").To("sent").From("draft").After(
		func(brief interface{}, tx *gorm.DB) error {
			// 发送通知
			return nil
		})

	return recordStateMachine
}

func GetBriefById(ID interface{}) (TaskBrief, error) {
	var brief TaskBrief
	result := DB.First(&brief, ID)
	return brief, result.Error
}
