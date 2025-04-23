package account

import (
	"context"
	"net/http"

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
	userModel := l.svcCtx.UserModel
	_, err = userModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		l.Logger.Errorf("查询用户失败，原因：%v", err)
		return &types.RegisterResp{
			Code: http.StatusInternalServerError,
			Msg:  "注册失败,用户已存在",
		}, nil
	}

	// 2. 插入用户
	if _, err := userModel.Insert(l.ctx, &model.User{
		Username: req.Username,
		Password: req.Password,
		Account:  req.Username,
	}); err != nil {
		l.Logger.Errorf("插入用户失败，原因：%v", err)
		return &types.RegisterResp{
			Code: http.StatusInternalServerError,
			Msg:  "注册失败,请稍后重试!",
		}, nil
	}
	return &types.RegisterResp{
		Code: http.StatusOK,
		Msg:  "注册成功",
	}, nil
}
