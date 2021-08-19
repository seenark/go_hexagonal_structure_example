package handler

import (
	"bank/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type customerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(cusService service.CustomerService) customerHandler {
	return customerHandler{
		customerService: cusService,
	}
}

func (h customerHandler) GetAllCustomer(c echo.Context) error {
	customers, err := h.customerService.GetCustomers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, customers)
}

func (h customerHandler) GetCustomerById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "Invalid Id")
	}
	customer, err := h.customerService.GetCustomer(id)
	if err != nil {
		if err.Error() == "not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"message": fmt.Sprintf("Not Found id %s", id)})
		}
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, customer)

}
