package rest

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"segments-api/internal/model/segment"
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

func (r SegmentRoute) GetAllByUser(ctx echo.Context) error {
	segments, err := r.service.GetAllByUser(1)
	if err != nil {
		return err
	}
	_ = segment.ToDTOs(segments)

	return nil
}
