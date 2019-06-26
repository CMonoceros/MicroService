package http

import (
	"SnowBrick-Backend/common/log"
	"SnowBrick-Backend/common/response"
	"SnowBrick-Backend/internal/model/req"
	"github.com/gin-gonic/gin"
)

func listSets(ctx *gin.Context) {
	var setsReq req.ListSetsReq
	if err := ctx.ShouldBindQuery(&setsReq); err != nil {
		log.Error("listSets ctx.ShouldBindQuery error(%v)", err)
		response.JSON(ctx, nil, err)
		return
	}

	sets, err := svc.ListSets(ctx, &setsReq)

	response.JSON(ctx, sets, err)
}
