package account

import (
	"context"
	"net/http"

	"infni-backend/apps/api/internal/svc"
	"infni-backend/apps/api/internal/types"
	"infni-backend/apps/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 1. 通过用户名查看用户是否存在
	userModel := l.svcCtx.UserModel
	user, err := userModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil || user == nil {
		l.Logger.Errorf("查询用户失败，原因：%v", err)
		return &types.LoginResp{
			Code: http.StatusInternalServerError,
			Msg:  "登录失败,用户不存在",
		}, nil
	}

	// 2. 判断用户名和密码是否一致
	if user.Password != req.Password {
		l.Logger.Errorf("用户名或密码错误")
		return &types.LoginResp{
			Code: http.StatusInternalServerError,
			Msg:  "登录失败,用户名或密码错误",
		}, nil
	}

	// 3. 登录成功，生成token并返回
	token, err := utils.GenerateToken(user.Id, user.Username, l.svcCtx.Config)
	if err != nil {
		l.Logger.Errorf("生成token失败，原因：%v", err)
		return &types.LoginResp{
			Code: http.StatusInternalServerError,
			Msg:  "登录失败,生成token失败",
		}, nil
	}

	return &types.LoginResp{
		Code: http.StatusOK,
		Msg:  "登录成功",
		Data: types.Resp{
			Token:    token,
			UserId:   user.Id,
			Username: user.Username,
			ExpireAt: int64(l.svcCtx.Config.Jwt.Expire),
		},
	}, nil
}
