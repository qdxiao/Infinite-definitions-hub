package account

import (
	"context"
	"errors"

	"infni-backend/apps/api/internal/model"
	"infni-backend/apps/api/internal/svc"
	"infni-backend/apps/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// 1. 根据用户名查询该用户是否存在
	userModel := model.NewUserModel(l.svcCtx.Mysql)
	_, err = userModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		l.Logger.Errorf("查询用户失败，原因：%v", err)
		return nil, errors.New("该用户已存在")
	}

	// 2. 插入用户
	if _, err := userModel.Insert(l.ctx, &model.User{
		Username: req.Username,
		Password: req.Password,
	}); err != nil {
		l.Logger.Errorf("插入用户失败，原因：%v", err)
		return nil, errors.New("注册失败")
	}

	return
}
