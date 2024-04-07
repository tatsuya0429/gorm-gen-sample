package user

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tatsuya0429/gorm-gen-sample/internal/models"
	"github.com/tatsuya0429/gorm-gen-sample/internal/queries"
	"github.com/tatsuya0429/gorm-gen-sample/internal/sqlhandler"
	"gorm.io/gen"
)

type UserController struct {
	db    *sqlhandler.SqlHandler
	query *queries.Query
}

func New(db *sqlhandler.SqlHandler) *UserController {
	return &UserController{
		db:    db,
		query: queries.Use(db.Conn, &gen.DOConfig{}),
	}
}

func (con *UserController) FindByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(400, "invalid id")
	}
	user, err := con.query.WithContext(c.Request().Context()).User.FindByID(int64(id))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(500, "internal server error")
	}
	return c.JSON(200, user)
}

func (con *UserController) FindAll(c echo.Context) error {
	users, err := con.query.WithContext(c.Request().Context()).User.FindAll()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(500, "internal server error")
	}
	return c.JSON(200, users)
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (con *UserController) Create(c echo.Context) error {
	var user CreateUserRequest
	if err := c.Bind(&user); err != nil {
		c.Logger().Error(err)
		return c.JSON(400, "invalid request")
	}
	if err := con.query.WithContext(c.Request().Context()).User.Create(&models.User{
		Username: user.Name,
		Password: user.Password,
	}); err != nil {
		c.Logger().Error(err)
		return c.JSON(500, "internal server error")
	}
	return c.JSON(200, user)
}
