package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/LucasBelusso1/GoExpert/7-APIS/internal/dto"
	"github.com/LucasBelusso1/GoExpert/7-APIS/internal/entity"
	"github.com/LucasBelusso1/GoExpert/7-APIS/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: db}
}

// GetJWT doc
// @Summary		Get a user JWT
// @Description	Get a user JWT
// @tags		users
// @Accept		json
// @Produce		json
// @Param		request body dto.GetJWTInput true "user credentials"
// @Success		200 {object} dto.GetJWTOutput
// @Failure		404 {object} Error
// @Failure		500 {object} Error
// @Router		/users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("TokenAuth").(*jwtauth.JWTAuth)
	expires := r.Context().Value("JWTExpiresIn").(int)

	var jwtDTO dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&jwtDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	user, err := h.UserDB.FindByEmail(jwtDTO.Email)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	if !user.ValidatePassword(jwtDTO.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, token, _ := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(expires)).Unix(),
	})

	accessToken := dto.GetJWTOutput{AccessToken: token}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create		user doc
// @Summary		Create user
// @Description	Create user
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		request body dto.CreateUserInput true "user request"
// @Success		201
// @Failure		500 {object} Error
// @Router		/users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	user, err := entity.NewUser(userDTO.Name, userDTO.Email, userDTO.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.UserDB.Create(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
