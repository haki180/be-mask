package domain

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Datas struct {
	Id        uint64
	UUID      string
	MaskType  int
	ImageUri  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataTransformer struct {
	UUID      string `json:"uuid"`
	MaskType  string `json:"mask_type"`
	ImageUri  string `json:"image_uri"`
	CreatedAt string `json:"created_at"`
}

func (d Datas) ToTransform() *DataTransformer {
	return &DataTransformer{
		UUID:      d.UUID,
		MaskType:  d.GetMaskType(),
		ImageUri:  d.ImageUri,
		CreatedAt: fmt.Sprintf("%s %s %s %s:%s:%s", strconv.Itoa(d.CreatedAt.Day()), d.CreatedAt.Month().String(), strconv.Itoa(d.CreatedAt.Year()), strconv.Itoa(d.CreatedAt.Hour()), strconv.Itoa(d.CreatedAt.Minute()), strconv.Itoa(d.CreatedAt.Second())),
	}
}

func (d Datas) GetMaskType() string {
	switch d.MaskType {
	case 1:
		return "Proper Facemask"
	case 2:
		return "Improper Facemask"
	case 3:
		return "No Mask"
	}
	return ""
}

func (d *Datas) IsEmpty() bool {
	return d == nil
}

type GetDataChartParams struct {
	Start string `query:"start"`
	End   string `query:"end"`
}

func (g *GetDataChartParams) SetStart(start string) {
	g.Start = start
}

func (g *GetDataChartParams) SetEnd(end string) {
	g.End = end
}

type DataChartTransfmer struct {
	Proper   int `json:"proper"`
	Improper int `json:"improper"`
	No       int `json:"no"`
}

func NewDataChartTransfmer(proper int, improper int, no int) *DataChartTransfmer {
	return &DataChartTransfmer{
		Proper:   proper,
		Improper: improper,
		No:       no,
	}
}

type CreateDataRequest struct {
	MaskType int    `json:"mask_type" binding:"required"`
	ImageUri string `json:"image_uri" binding:"required"`
}

func (c CreateDataRequest) ToData() *Datas {
	return &Datas{
		UUID:     uuid.New().String(),
		MaskType: c.MaskType,
		ImageUri: c.ImageUri,
	}
}

func NewCreateDataResponse(c *gin.Context, code int, message interface{}) {
	c.JSON(code, map[string]interface{}{
		"status_code": code,
		"message":     message,
	})
}

type GetDatasPaginateParams struct {
	Page  int     `query:"page"`
	Limit int     `query:"limit"`
	Type  *int    `query:"type"`
	Start *string `query:"start"`
	End   *string `query:"end"`
}

func (g GetDatasPaginateParams) IsStartEmpty() bool {
	return g.Start == nil
}

func (g GetDatasPaginateParams) IsEndEmpty() bool {
	return g.End == nil
}

func (g GetDatasPaginateParams) IsTypeEmpty() bool {
	return g.Type == nil
}

func (g *GetDatasPaginateParams) setDefaultPage() {
	g.Page = 1
}

func (g *GetDatasPaginateParams) setDefaultLimit() {
	g.Limit = 1000
}

func (g *GetDatasPaginateParams) SetPage(page string) {
	if page != "" {
		pageInt, _ := strconv.Atoi(page)
		g.Page = pageInt
	} else {
		g.setDefaultPage()
	}
}

func (g *GetDatasPaginateParams) SetLimit(limit string) {
	if limit != "" {
		limitInt, _ := strconv.Atoi(limit)
		g.Limit = limitInt
	} else {
		g.setDefaultLimit()
	}
}

func getMaskType(maskType string) int {
	switch maskType {
	case "proper":
		return 1
	case "improper":
		return 2
	case "no":
		return 3
	default:
		return 10
	}
}

func (g *GetDatasPaginateParams) SetType(maskType string) {
	if maskType != "" {
		typeInt := getMaskType(maskType)
		if typeInt != 10 {
			g.Type = &typeInt
		}
	}
}

func (g *GetDatasPaginateParams) SetStart(start string) {
	if start != "" {
		g.Start = &start
	}
}

func (g *GetDatasPaginateParams) SetEnd(end string) {
	if end != "" {
		g.End = &end
	}
}

type GetDataPaginateResponse struct {
	Data []DataTransformer `json:"data"`
	Meta PaginateMeta      `json:"meta"`
}

func NewGetDataPaginateResponse(data []DataTransformer, meta PaginateMeta) *GetDataPaginateResponse {
	return &GetDataPaginateResponse{
		Data: data,
		Meta: meta,
	}
}

type PaginateMeta struct {
	Page      int `json:"page"`
	TotalPage int `json:"total_pages"`
	TotalData int `json:"total_datas"`
}

func NewPaginateMeta(page int, totalPage int, totalDatas int) *PaginateMeta {
	return &PaginateMeta{
		Page:      page,
		TotalPage: totalPage,
		TotalData: totalDatas,
	}
}
