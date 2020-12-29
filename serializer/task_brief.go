package serializer

import (
	"singo/model"
)

type TaskBrief struct {
	ID             uint   `json:"id"`
	TenantUuid     string `json:"tenant_uuid"`
	TaskPeriodId   uint   `json:"task_period_id"`
	UserId         uint   `json:"user_id"`
	TaskId         uint   `json:"task_id"`
	Title          string `json:"title"`
	CreatedAt      int64  `json:"created_at"`
	DepartmentId   uint   `json:"department_id"`
	ProcessContent string `json:"process_content"`
	IssueContent   string `json:"issue_content"`
	PlanContent    string `json:"plan_content"`
}

func BuildTaskBrief(brief model.TaskBrief) TaskBrief {
	return TaskBrief{
		ID:             brief.ID,
		Title:          brief.Title,
		UserId:         brief.UserId,
		TenantUuid:     brief.TenantUuid,
		TaskPeriodId:   brief.TaskPeriodId,
		TaskId:         brief.TaskId,
		DepartmentId:   brief.DepartmentId,
		ProcessContent: brief.ProcessContent,
		IssueContent:   brief.IssueContent,
		PlanContent:    brief.PlanContent,

		CreatedAt: brief.CreatedAt.Unix(),
	}
}

func BuildTaskBriefResponse(brief model.TaskBrief) Response {
	return Response{
		Data: BuildTaskBrief(brief),
	}
}
