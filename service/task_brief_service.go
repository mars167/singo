package service

import (
	"fmt"
	"singo/model"
	"singo/serializer"
)

type TaskBriefService struct {
	TenantUuid   string `form:"tenant_uuid" json:"tenant_uuid" binding:"required"`
	TaskPeriodId uint   `form:"task_period_id" json:"task_period_id" binding:"required"`
	UserId       uint   `from:"user_id" json:"user_id" binding:"required"`
	TaskId       uint   `from:"task_id" json:"task_id" binding:"required"`
	DepartmentId uint   `from:"department_id" json:"department_id" binding:"required"`

	Title          string `form:"title" json:"title" binding:"required"`
	ProcessContent string `from:"process_content" json:"process_content" binding:"required"`
	IssueContent   string `from:"issue_content" json:"issue_content" binding:"required"`
	PlanContent    string `from:"plan_content" json:"plan_content" binding:"required"`
}

func (service *TaskBriefService) CreateBrief() serializer.Response {
	fmt.Println("[PARAMS]")
	fmt.Println(service.TenantUuid)
	fmt.Println(service.TaskPeriodId)
	fmt.Println(service.UserId)
	fmt.Println(service.TaskId)
	fmt.Println(service.DepartmentId)
	fmt.Println(service.ProcessContent)
	fmt.Println(service.IssueContent)
	fmt.Println(service.PlanContent)

	brief := model.TaskBrief{
		TenantUuid:     service.TenantUuid,
		TaskPeriodId:   service.TaskPeriodId,
		Title:          service.Title,
		UserId:         service.UserId,
		TaskId:         service.TaskId,
		DepartmentId:   service.DepartmentId,
		ProcessContent: service.ProcessContent,
		IssueContent:   service.IssueContent,
		PlanContent:    service.PlanContent,
	}
	if err := model.DB.Create(&brief).Error; err != nil {
		return serializer.ParamErr("创建失败", err)
	}
	return serializer.BuildTaskBriefResponse(brief)
}
