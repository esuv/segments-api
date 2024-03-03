package rest

import "github.com/labstack/echo/v4"

type SegmentService interface {
	Create(slug string) (int, error)
	Delete(slug string) error
	AddUser(add []string, remove []string, userId int) error
	GetAllByUser(userID int)
}

type SegmentRoute struct {
	service *SegmentService
}

func New(service SegmentService) SegmentRoute {
	return SegmentRoute{service: &service}
}

type inputCreateSegment struct {
	Name string `json:"name"`
}

func (r *SegmentRoute) Create(ctx echo.Context) error {
	var inp inputCreateSegment
	if err := ctx.Bind(&inp); err != nil {

		return err
	}

	return nil
}

func (r *SegmentRoute) Delete(ctx echo.Context) error {
	return nil
}

func (r *SegmentRoute) AddUser(ctx echo.Context) error {
	return nil
}

func (r *SegmentRoute) GetAllByUser(ctx echo.Context) error {
	return nil
}
