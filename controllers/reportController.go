package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/dtos/report"
	"server/interface/Service"
	"server/utils"
	"strconv"
)

type ReportController struct {
	reportService Service.ReportServiceInterface
}

func NewReportController(reportService Service.ReportServiceInterface) *ReportController {
	return &ReportController{reportService: reportService}
}

func (r *ReportController) CreateReport(c *gin.Context) {
	var reportCreateDto report.CreateReportDto
	if err := c.ShouldBind(&reportCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	if err := reportCreateDto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	report, err := r.reportService.CreateReport(&reportCreateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    report,
		Error:   "",
	})

}

func (r *ReportController) DeleteReport(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("reportId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
			Data:    nil,
		})
		return
	}
	if err := r.reportService.DeleteReport(id); err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    nil,
		Error:   "",
	})
	return
}
func (r *ReportController) UpdateReport(c *gin.Context) {
	var reportUpdateDto report.UpdateReportDto
	reportId, err := strconv.Atoi(c.Param("reportId"))

	if err := c.ShouldBind(&reportUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, &utils.Response{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	report, err := r.reportService.UpdateReport(reportId, reportUpdateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    report,
		Error:   "",
	})

}

func (r *ReportController) GetAllReport(c *gin.Context) {
	report, err := r.reportService.GetAllReports()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &utils.Response{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &utils.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    report,
		Error:   "",
	})
	return
}
