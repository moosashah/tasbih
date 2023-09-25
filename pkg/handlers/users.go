package handlers

import (
	"net/http"
	"tasbih/pkg/database"
	"tasbih/pkg/store"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type home struct {
	Header header
	Name   string
}

type signup_page struct {
	Header header
	Errors string
}

func Signup_page(c echo.Context) error {
	return c.Render(http.StatusOK, "signup.html", signup_page{
		Header: header{
			Title: "Sign up",
		},
	})
}

type Handler struct {
	us store.Store
}

type errorBlock struct {
	Err string
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

func (h *Handler) ValidateEmail(c echo.Context) error {
	e := c.FormValue("email")
	_, err := h.us.GetUserByEmail(e)
	if err != nil {
		c.Logger().Errorf("Email already taken")
		return c.String(http.StatusConflict, "Email is already taken")
	}
	return err
}

func Login_page(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", page{
		header{
			Title: "Login",
		},
	})
}

func (h *Handler) Login(c echo.Context) error {
	e := c.FormValue("email")
	u, err := h.us.GetUserByEmail(e)
	if err != nil {
		return c.Render(http.StatusNotFound, "error", errorBlock{
			Err: err.Error(),
		})
	}

	return c.Render(http.StatusOK, "home.html", home{
		Header: header{
			Title: "Home",
		},
		Name: u.Name,
	})
}

func Home_page(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", home{
		Header: header{
			Title: "Home",
		},
	})
}
