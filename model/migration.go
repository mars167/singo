package model

import "github.com/qor/transition"

//执行数据迁移

func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&TaskBrief{})
	// 状态机change log,记录状态变化
	DB.AutoMigrate(&transition.StateChangeLog{})

}
