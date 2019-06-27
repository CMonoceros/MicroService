package service

import (
	"SnowBrick-Backend/common/oss"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"SnowBrick-Backend/common/log"
	"SnowBrick-Backend/internal/model/req"
	"SnowBrick-Backend/internal/model/resp"
)

func (s *Service) getAgeRange(minAge, maxAge int8) string {
	if maxAge == -1 && minAge == -1 {
		return fmt.Sprintf("-")
	} else if maxAge == -1 {
		return fmt.Sprintf("%s+", strconv.Itoa(int(minAge)))
	}
	return fmt.Sprintf("%s-%s", strconv.Itoa(int(minAge)), strconv.Itoa(int(maxAge)))
}

func (s *Service) ListSets(ctx *gin.Context, setReq *req.ListSetsReq) (res []*resp.ListSetResp, err error) {
	sets, err := s.dao.RawBrickSet(ctx, &setReq.BasePage)
	if err != nil {
		log.Error("ListSets RawBrickSet error(%v)", err)
		return nil, err
	}

	bucket, err := oss.New(s.c.Oss, oss.BUCKET_BRICK_SET)
	if err != nil {
		log.Error("ListSets oss.New error(%v)", err)
		return nil, err
	}

	for _, set := range sets {
		eachResp := new(resp.ListSetResp)
		eachResp.UpdatedAt = set.UpdatedAt.Format("2006-01-02 15:04:05")
		eachResp.CreatedAt = set.CreatedAt.Format("2006-01-02 15:04:05")
		eachResp.EnglishName = set.EnglishName
		eachResp.ChineseName = set.ChineseName
		eachResp.ChineseDescription = set.ChineseDescription
		eachResp.MinifigCount = int(set.MinifigCount)
		eachResp.PartCount = int(set.PartCount)
		eachResp.ReleaseDate = set.ReleaseDate.Format("2006")
		eachResp.Number = set.Number
		eachResp.AgeRange = s.getAgeRange(set.AgeMin, set.AgeMax)

		if set.Type != nil {
			eachResp.Type = &resp.ListSetTypeResp{
				ChineseName: set.Type.ChineseName,
			}
		}
		if set.Theme != nil {
			eachResp.Theme = &resp.ListSetThemeResp{
				ChineseName: set.Theme.ChineseName,
			}
		}
		if set.SubTheme != nil {
			eachResp.SubTheme = &resp.ListSetSubThemeResp{
				ChineseName: set.SubTheme.ChineseName,
			}
		}
		if set.Brand != nil {
			eachResp.Brand = &resp.ListSetBrandResp{
				ChineseName: set.Brand.ChineseName,
			}
		}
		if set.Packaging != nil {
			eachResp.Packaging = &resp.ListSetPackagingResp{
				ChineseName: set.Packaging.ChineseName,
			}
		}
		if set.Tags != nil {
			var tags []*resp.ListSetTagResp
			for _, tag := range set.Tags {
				eachTagResp := new(resp.ListSetTagResp)
				eachTagResp.ChineseName = tag.ChineseName
				tags = append(tags, eachTagResp)
			}
			eachResp.Tags = tags
		}
		if set.Media != nil {
			var mediaList []*resp.ListSetMediaResp
			for _, media := range set.Media {
				eachMediaResp := new(resp.ListSetMediaResp)
				eachMediaResp.Type = media.Type
				eachMediaResp.Duration = media.Duration
				eachMediaResp.High = media.High
				eachMediaResp.Width = media.Width

				objectName, e := oss.ParseObjectName(oss.BUCKET_BRICK_SET, media.Src)
				if e != nil {
					log.Error("ListSets oss.ParseObjectName error(%v)", e)
					return nil, e
				}
				signSrc, e := bucket.SignURL(objectName, http.MethodGet,
					int64(time.Duration(s.c.Oss.SignExpireTime).Seconds()))
				if e != nil {
					log.Error("ListSets bucket.SignURL error(%v)", e)
					return nil, e
				}
				eachMediaResp.Src = signSrc

				mediaList = append(mediaList, eachMediaResp)
			}
			eachResp.Media = mediaList
		}

		res = append(res, eachResp)
	}

	return res, err
}
