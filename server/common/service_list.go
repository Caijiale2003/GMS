package common

import (
	"GameManageSystem/global"
	"GameManageSystem/models"
	"fmt"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo          // 分页查询
	Likes           []string // 需要模糊匹配的字段列表
	Debug           bool     // 是否打印sql
	Where           *gorm.DB // 额外的查询
	Preload         []string // 预加载的字段列表
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	//Where 添加查询条件 查询model
	query := global.DB.Where(model)
	//是否开启debug模式
	if option.Debug {
		//开启后执行查询时会输出相关的调试信息，例如生成的 SQL 查询语句以及执行的参数等
		query = query.Debug()
	}
	// 默认按照时间往后排
	// 默认一页显示10条
	if option.Limit == 0 {
		option.Limit = 10
	}
	// 如果有高级查询就加上
	if option.Where != nil {
		query = query.Where(option.Where)
	}
	// 模糊匹配
	if option.Key != "" {
		likeQuery := global.DB.Where("")
		for index, column := range option.Likes {
			// 第一个模糊匹配和前面的关系是and关系，后面的和前面的模糊匹配是or的关系
			// %%%s%% 匹配包含%s的字段
			if index == 0 {
				likeQuery.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			} else {
				likeQuery.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			}
		}
		// 整个模糊匹配它是一个整体
		query = query.Where(likeQuery)
	}
	// 查列表，获取总数
	count = query.Find(&list).RowsAffected
	// 预加载
	for _, preload := range option.Preload {
		query = query.Preload(preload)
	}
	// 计算偏移
	offset := (option.Page - 1) * option.Limit
	//Limit 一页显示几个 Offset从哪开始 Order按什么排
	err = query.Limit(option.Limit).
		Offset(offset).
		Order(option.Sort).Find(&list).Error

	return
}
