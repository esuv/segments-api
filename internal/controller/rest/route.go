package rest

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"segments-api/internal/logger/sl"
	"segments-api/internal/model/segment"
	"strconv"
)

type SegmentService interface {
	Create(slug string) (segment.Segment, error)
	Delete(slug string) error
	AddUser(add []string, remove []string, userId int) error
	GetAllByUser(userID int) ([]segment.Segment, error)
}

type SegmentRoute struct {
	service SegmentService
	log     *slog.Logger
}

func New(service SegmentService, log *slog.Logger) SegmentRoute {
	return SegmentRoute{service: service, log: log}
}

type inputCreateSegment struct {
	Name string `json:"name"`
}

// Create godoc
// @Summary      Create segment
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Segment ID"
// @Router       /segments [post]
func (r SegmentRoute) Create(ctx echo.Context) error {
	var inp inputCreateSegment
	if err := ctx.Bind(&inp); err != nil {
		//TODO return error

		return err
	}

	var sgm segment.Segment
	sgm, err := r.service.Create(inp.Name)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, &sgm)
}

type inputDeleteSegment struct {
	Name string `json:"name"`
}

func (r SegmentRoute) Delete(ctx echo.Context) error {
	var inp inputDeleteSegment
	if err := ctx.Bind(&inp); err != nil {
		//TODO return correct error

		return err
	}

	if err := r.service.Delete(inp.Name); err != nil {
		//TODO return error

		return err
	}

	return nil
}

type inputAddUser struct {
	AddSegments    []string `json:"add_segments"`
	DeleteSegments []string `json:"delete_segments"`
	UserID         int      `json:"user_id"`
}

func (r SegmentRoute) AddUser(ctx echo.Context) error {
	var inp inputAddUser
	if err := ctx.Bind(&inp); err != nil {
		//TODO return error

		return err
	}

	err := r.service.AddUser(inp.AddSegments, inp.DeleteSegments, inp.UserID)
	if err != nil {
		ctx.Error(err)
		fmt.Println(err.Error())
	}

	return err
}

// GetAllByUser godoc
// @Summary      Get all segments by User ID
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  dto.SegmentDTO
// @Router       /segments/users/{id} [get]
func (r SegmentRoute) GetAllByUser(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		return err
	}

	segments, err := r.service.GetAllByUser(userID)
	if err != nil {
		return err
	}
	result := segment.ToDTOs(segments)

	err = ctx.JSON(http.StatusOK, result)
	if err != nil {
		r.log.Error("marshalling error", sl.Err(err))
		return err
	}

	return nil
}
