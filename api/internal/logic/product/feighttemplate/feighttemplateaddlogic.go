package logic

import (
	"context"
	"encoding/json"
	"micro/common/errorx"
	"micro/rpc/svs/pms/pmsclient"

	"micro/api/internal/svc"
	"micro/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FeightTemplateAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeightTemplateAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) FeightTemplateAddLogic {
	return FeightTemplateAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeightTemplateAddLogic) FeightTemplateAdd(req types.AddFeightTemplateReq) (*types.AddFeightTemplateResp, error) {
	_, err := l.svcCtx.Pms.FeightTemplateAdd(l.ctx, &pmsclient.FeightTemplateAddReq{
		Name:           req.Name,
		ChargeType:     req.ChargeType,
		FirstWeight:    int64(req.FirstWeight),
		FirstFee:       int64(req.FirstFee),
		ContinueWeight: int64(req.ContinueWeight),
		ContinmeFee:    int64(req.FirstFee),
		Dest:           req.Dest,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加运费模板信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加运费模版失败")
	}

	return &types.AddFeightTemplateResp{
		Code:    "000000",
		Message: "添加运费模版成功",
	}, nil
}
