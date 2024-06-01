package handlers

import (
	"net/http"
	"strconv"

	"github.com/SuisF/goecho-health/cmd/api/service"
	"github.com/labstack/echo"
)

func PostIndexHandler(c echo.Context) error {
	data, err := service.GetAll()

	if err !=  nil {
		c.String(http.StatusBadGateway, "Unalbe to process data")
	}

	res := make(map[string]any)
	res["status"] = "Good"
	res["data"] = data

	return c.JSON(http.StatusOK, res)
}

func PostSingleHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)

	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process data")
	}

	data, err := service.GetById(idx)
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process data")
	}

	res := make(map[string]any)
	res["status"] = "Good"
	res["data"] = data
	return c.JSON(http.StatusOK, res )
}