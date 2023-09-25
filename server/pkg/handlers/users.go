package handlers

import (
	"net/http"
	"tasbih/pkg/database"
	"tasbih/pkg/store"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type home struct {
	Name string
}

type Handler struct {
	us store.Store
}

func NewHandler(us store.Store) *Handler {
	return &Handler{
		us: us,
	}
}

func (h *Handler) CreateUser(c echo.Context) error {
	user := database.User{
		ID:    uuid.New().String(),
		Name:  c.FormValue("name"),
		Email: c.FormValue("email"),
		//Password: utils.HashPassword(c.FormValue("password")),
	}
	err := h.us.Create(&user)
	if err != nil {
		c.Logger().Errorf("could not create user: %s", err.Error())
		c.Response().Header().Add("HX-Retarget", "#errors")
		c.Response().Header().Add("HX-Reswap", "innerHTML")
		return c.String(http.StatusConflict, err.Error())
	}
	c.Response().Header().Add("HX-Redirect", "/home")
	return err
}

type errorMessage struct {
	Err string `json:"error"`
}

func (h *Handler) Login(c echo.Context) error {
	e := c.FormValue("email")
	u, err := h.us.GetUserByEmail(e)
	if err != nil {
		return c.JSON(http.StatusNotFound, errorMessage{
			Err: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, home{
		Name: u.Name,
	})
}
