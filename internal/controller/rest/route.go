package rest

import "github.com/labstack/echo/v4"

type SegmentService interface {
	Create(slug string) (int, error)
	Delete(slug string) error
	AddUser(add []string, remove []string, userId int) error
	GetAllByUser(userID int)
}

type SegmentRoute struct {
	service SegmentService
}

func New(service SegmentService) SegmentRoute {
	return SegmentRoute{service: service}
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

	if _, err := r.service.Create(inp.Name); err != nil {
		//TODO return error

		return err
	}

	return nil
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

	if err := r.service.AddUser(inp.AddSegments, inp.DeleteSegments, inp.UserID); err != nil {
		//TODO return error

		return err
	}

	return nil
}

func (r SegmentRoute) GetAllByUser(ctx echo.Context) error {
	return nil
}
