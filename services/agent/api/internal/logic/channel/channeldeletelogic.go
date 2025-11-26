// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package channel

import (
	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	"chihqiang/msgbox-go/services/agent/api/internal/types"
	"chihqiang/msgbox-go/services/common/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChannelDeleteLogic {
	return &ChannelDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelDeleteLogic) ChannelDelete(req *types.IDReq) error {
	agentID, err := l.ctx.Value(types.JWTAgentID).(json.Number).Int64()
	if err != nil {
		return fmt.Errorf("not find api")
	}
	var channel models.Channel
	if err := l.svcCtx.DB.Where("id = ? AND agent_id = ?", req.ID, agentID).First(&channel).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return errors.New("通道不存在")
		}
		return err
	}
	// 2. 检查是否有关联的模板
	// GORM 的 Association.Count 会自动使用 channel 的主键 (ID) 进行查询
	count := l.svcCtx.DB.Model(&channel).Association("Templates").Count()
	if count > 0 {
		return fmt.Errorf("通道仍被 %d 个模板使用，无法删除", count)
	}
	// 3. 如果没有关联，可以安全删除通道
	if err := l.svcCtx.DB.Delete(&channel).Error; err != nil {
		return err
	}
	return nil
}
