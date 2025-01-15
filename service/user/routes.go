package user

import (
	"fmt"
	"net/http"

	"github.com/Sahal-P/Go-Auth/types"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	userRouter := router.Group("/user")

	userRouter.POST("/login", h.LoginUser)
	userRouter.POST("/register", h.RegisterUser)
}

func (h *Handler) LoginUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "User loged in"})
}

func (h *Handler) RegisterUser(ctx *gin.Context) {

	var payload types.RegisterUserPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if validationErrors := payload.Validate(); validationErrors != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErrors})
		return
	}

	userExists, err := h.store.GetUserByEmail(payload.Email)

	if err != nil {
		// Handle the case where there's an error querying the database
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking email"})
		return
	}

	// If user is found (not nil), that means the user already exists
	if userExists != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("user with email %s already exists", payload.Email)})
		return
	}
	user := &types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  payload.Password, // Ensure you hash the password before storing it
	}
	createdUser, err := h.store.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	ctx.JSON(200, gin.H{
		"id":         createdUser.ID,
		"status":     "created",
		"created_at": createdUser.CreatedAt,
	})

}
