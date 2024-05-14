package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mestvv/NNBlogBackend/internal/service"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/register", h.userRegister)
		users.POST("/sign-in", h.userAuth)
		users.POST("/auth/refresh", h.userRefresh)
	}
}

type userRegisterRequest struct {
	Username  string  `json:"name" binding:"required,min=2,max=64" example:"wazzup"`
	FirstName *string `json:"first_name" binding:"omitempty,min=2,max=64" example:"John"`
	LastName  *string `json:"last_name" binding:"omitempty,min=2,max=64" example:"Doe"`
	Email     string  `json:"email" binding:"required,email,max=64" example:"mail@mail.com"`
	Password  string  `json:"password" binding:"required,min=8,max=64" example:"notasecretpassword"`
}

// @Summary User Register
// @Tags User Auth
// @Description Create user account
// @ModuleID userRegister
// @Accept  json
// @Produce  json
// @Param input body userRegisterRequest true "register info"
// @Success 201
// @Failure 400 {object} ErrorStruct
// @Failure 500
// @Router /users/register [post]
func (h *Handler) userRegister(c *gin.Context) {
	var req userRegisterRequest
	if err := c.BindJSON(&req); err != nil {
		validationErrorResponse(c, err)
		return
	}

	err := h.services.Users.Register(c.Request.Context(), service.UserRegisterInput{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	})

	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExist) {
			errorResponse(c, UserAlreadyExistsCode)
			return
		}
		h.logger.Error("failed to create user", "error", err)
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) userAuth(c *gin.Context) {

}

func (h *Handler) userRefresh(c *gin.Context) {

}
