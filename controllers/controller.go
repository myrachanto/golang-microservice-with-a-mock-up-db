package controllers

import(	
	"strconv"
	"net/http"
	"github.com/labstack/echo"
	"github.com/myrachanto/amicroservice/customermicroservice/httperors"
	"github.com/myrachanto/amicroservice/customermicroservice/model"
	"github.com/myrachanto/amicroservice/customermicroservice/service"
)

var (
	CustomerController customerController = customerController{}
)
type customerController struct{ }
/////////controllers/////////////////
func (controller customerController) Create(c echo.Context) error {
	customer := &model.Customer{}
	if err := c.Bind(customer); err != nil {
		httperror := httperors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}
	createdCustomer, err1 := service.CustomerService.Create(customer)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, createdCustomer)
}

func (controller customerController) GetAll(c echo.Context) error {
	customers := map[int]*model.Customer{}
	customers, err3 := service.CustomerService.GetAll(customers)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	if len(customers) == 0 {
		heror := httperors.NewNoResultsMessage("No results found")
		return c.JSON(heror.Code, heror)
	}
	return c.JSON(http.StatusOK, customers)
} 
func (controller customerController) GetOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil { 
		httperror := httperors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}
	customer, problem := service.CustomerService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, customer)	
}

func (controller customerController) Update(c echo.Context) error {
	customer := &model.Customer{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	if err := c.Bind(customer); err != nil {
		httperror := httperors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}
	Updatedcustomer, problem := service.CustomerService.Update(id, customer)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, Updatedcustomer)
}

func (controller customerController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	success, failure := service.CustomerService.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}