package service

import (
	"SnowBrick-Backend/common/encrypt"
	"SnowBrick-Backend/common/oss"
	"SnowBrick-Backend/internal/brick/model/entity"
	"SnowBrick-Backend/internal/brick/model/req"
	"SnowBrick-Backend/internal/brick/model/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"

	"SnowBrick-Backend/common/log"
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
		eachResp.ID, err = encrypt.EncodeID(s.c.Encrypt.HashID.Salt, int(set.ID))
		if err != nil {
			log.Error("ListSets EncodeID error(%v)", err)
			return nil, err
		}
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
			typeID, e := encrypt.EncodeID(s.c.Encrypt.HashID.Salt, int(set.TypeID))
			if e != nil {
				log.Error("ListSets EncodeID error(%v)", e)
				return nil, e
			}
			eachResp.Type = &resp.ListSetTypeResp{
				ID:          typeID,
				ChineseName: set.Type.ChineseName,
			}
		}
		if set.Theme != nil {
			themeID, e := encrypt.EncodeID(s.c.Encrypt.HashID.Salt, int(set.ThemeID))
			if e != nil {
				log.Error("ListSets EncodeID error(%v)", e)
				return nil, e
			}
			eachResp.Theme = &resp.ListSetThemeResp{
				ID:          themeID,
				ChineseName: set.Theme.ChineseName,
			}
		}
		if set.SubTheme != nil {
			subThemeID, e := encrypt.EncodeID(s.c.Encrypt.HashID.Salt, int(set.SubThemeID))
			if e != nil {
				log.Error("ListSets EncodeID error(%v)", e)
				return nil, e
			}
			eachResp.SubTheme = &resp.ListSetSubThemeResp{
				ID:          subThemeID,
				ChineseName: set.SubTheme.ChineseName,
			}
		}
		if set.Brand != nil {
			brandID, e := encrypt.EncodeID(s.c.Encrypt.HashID.Salt, int(set.BrandID))
			if e != nil {
				log.Error("ListSets EncodeID error(%v)", e)
				return nil, e
			}
			eachResp.Brand = &resp.ListSetBrandResp{
				ID:          brandID,
				ChineseName: set.Brand.ChineseName,
			}
		}
		if set.Packaging != nil {
			packagingID, e := encrypt.EncodeID(s.c.Encrypt.HashID.Salt, int(set.PackagingID))
			if e != nil {
				log.Error("ListSets EncodeID error(%v)", e)
				return nil, e
			}
			eachResp.Packaging = &resp.ListSetPackagingResp{
				ID:          packagingID,
				ChineseName: set.Packaging.ChineseName,
			}
		}
		if set.Tags != nil {
			tags, e := s.getSetLimitTagResponse(set.Tags, 3)
			if e != nil {
				log.Error("ListSets getSetLimitTagResponse error(%v)", e)
				return nil, e
			}
			eachResp.Tags = tags
		}
		if set.Media != nil && len(set.Media) > 0 {
			cover, e := s.getSetCoverMediaResponse(set.Media, bucket)
			if e != nil {
				log.Error("ListSets getSetCoverMediaResponse error(%v)", e)
				return nil, e
			}
			eachResp.Cover = cover
		}

		res = append(res, eachResp)
	}

	return res, err
}

func (s *Service) getSetLimitTagResponse(tags []*entity.BrickSetTag, limit int) (tagsResp []*resp.ListSetTagResp, e error) {
	for _, tag := range tags[0:int(math.Min(float64(len(tags)), float64(limit)))] {
		eachTagResp := new(resp.ListSetTagResp)
		tagID, e := encrypt.EncodeID(s.c.Encrypt.HashID.Salt, int(tag.ID))
		if e != nil {
			log.Error("getSetLimitTagResponse error(%v)", e)
			return nil, e
		}
		eachTagResp.ChineseName = tag.ChineseName
		eachTagResp.ID = tagID
		tagsResp = append(tagsResp, eachTagResp)
	}
	return
}

func (s *Service) getSetCoverMediaResponse(mediaList []*entity.BrickMedia, bucket *oss.Bucket) (mediaResp *resp.ListSetMediaResp, e error) {
	media := mediaList[0]
	mediaID, e := encrypt.EncodeID(s.c.Encrypt.HashID.Salt, int(media.ID))
	if e != nil {
		log.Error("getSetCoverMediaResponse error(%v)", e)
		return nil, e
	}
	mediaResp = new(resp.ListSetMediaResp)
	mediaResp.ID = mediaID
	mediaResp.Type = media.Type
	mediaResp.Duration = media.Duration
	mediaResp.High = media.High
	mediaResp.Width = media.Width
	signSrc, e := bucket.SignUrl(media.Src, "image", "/quality,Q_20")
	if e != nil {
		log.Error("getSetCoverMediaResponse error(%v)", e)
		return nil, e
	}
	mediaResp.Src = signSrc

	return mediaResp, nil
}
