package infrastructure

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/datake914/logue/src/repository/cloudwatchlogsrepo"
	"github.com/datake914/logue/src/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	validator "gopkg.in/go-playground/validator.v9"
)

// Run starts an HTTP server.
func Run() {
	e := echo.New()
	e.Use(middleware.CORS())
	// Set body dumper.
	e.Use(middleware.BodyDump(func(c echo.Context, req, res []byte) {
		c.Logger().Info(string(req))
		c.Logger().Info(string(res))
	}))
	// Set error handler.
	e.HTTPErrorHandler = handleError
	// Set custom validator.
	e.Validator = NewCustomValidator()

	// Set Routing.
	e.GET("/api/groups", func(c echo.Context) error {
		// Convert query paramters to service request.
		req := service.SearchLogGroupRequest{}
		if err := c.Bind(&req); err != nil {
			return err
		}
		// Validate.
		if err := c.Validate(&req); err != nil {
			return err
		}
		// Execute.
		res, err := service.NewLogGroupService(
			cloudwatchlogsrepo.NewLogGroupRepository(),
		).Search(req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, res)
	})

	e.GET("/api/streams", func(c echo.Context) error {
		// Convert query paramters to service request.
		req := service.SearchLogStreamRequest{}
		if err := c.Bind(&req); err != nil {
			return err
		}
		// Validate.
		if err := c.Validate(&req); err != nil {
			return err
		}
		// Execute.
		res, err := service.NewLogStreamService(
			cloudwatchlogsrepo.NewLogStreamRepository(),
		).Search(req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, res)
	})

	e.GET("/api/events", func(c echo.Context) error {
		// Convert query paramters to service request.
		req := service.SearchLogEventRequest{}
		if err := c.Bind(&req); err != nil {
			return err
		}
		// Validate.
		if err := c.Validate(&req); err != nil {
			return err
		}
		// Execute.
		res, err := service.NewLogEventService(
			cloudwatchlogsrepo.NewLogEventRepository(),
		).Search(req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, res)
	})

	// Start an HTTP server.
	e.Logger.Debug(e.Start(":" + strconv.Itoa(*Opts.Port)))
}

// ErrorResponse represents the error response parameters.
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ValidationErrorResponse represents the error response parameters.
type ValidationErrorResponse struct {
	ErrorResponse
	Details []ValidationErrorDetail
}

// ValidationErrorDetail represents the validation error response parameters.
type ValidationErrorDetail struct {
	Field   string `json:"code"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

func handleError(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}
	// Logging.
	c.Logger().Error(err)
	// Handle validation error.
	if errs, ok := err.(validator.ValidationErrors); ok {
		res := &ValidationErrorResponse{
			ErrorResponse: ErrorResponse{
				Code:    "10400",
				Message: "Field validation error.",
			},
		}
		for _, err := range errs {
			res.Details = append(res.Details, ValidationErrorDetail{
				Field:   err.Field(),
				Tag:     err.ActualTag(),
				Message: fmt.Sprintf("Field validation for %s failed on the %s tag.", err.Field(), err.ActualTag()),
			})
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// Handle unexpected error.
	res := ErrorResponse{
		Code:    "99999",
		Message: "An unexpected error has occured.",
	}
	c.JSON(http.StatusInternalServerError, res)
}
