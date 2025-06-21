package handler

import (
	// "context"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	pb "github.com/bag-huyag/api-gateway/proto/gen"
	// "github.com/google/uuid"
)

type Handler struct {
	UserClient pb.UserServiceClient
}

// POST /users
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.UserClient.CreateUser(r.Context(), &pb.NewUser{
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		http.Error(w, "gRPC error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// PUT /users/{id}
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user := &pb.User{
		Id:    id,
		Name:  req.Name,
		Email: req.Email,
	}

	updated, err := h.UserClient.UpdateUser(r.Context(), user)
	if err != nil {
		http.Error(w, "gRPC error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updated)
}

// DELETE /users/{id}
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := h.UserClient.DeleteUser(r.Context(), &pb.UserId{Id: id})
	if err != nil {
		http.Error(w, "gRPC error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// type UserHandler struct {
// 	pb.UnimplementedUserServiceServer
// 	users map[string]*pb.User
// }

// func (h *UserHandler) ensureUsers() {
// 	if h.users == nil {
// 		h.users = make(map[string]*pb.User)
// 	}
// }

// func (h *UserHandler) GetUsers(ctx context.Context, in *pb.Empty) (*pb.UserList, error) {
// 	h.ensureUsers()

// 	users := make([]*pb.User, 0, len(h.users))
// 	for _, user := range h.users {
// 		users = append(users, user)
// 	}

// 	return &pb.UserList{Users: users}, nil
// }

// func (h *UserHandler) GetUser(ctx context.Context, in *pb.UserId) (*pb.User, error) {
// 	h.ensureUsers()
// 	return h.users[in.Id], nil
// }

// func (h *UserHandler) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
// 	h.ensureUsers()
// 	id := uuid.New().String()
// 	user := &pb.User{
// 		Id:    id,
// 		Name:  in.Name,
// 		Email: in.Email,
// 	}
// 	h.users[id] = user
// 	return user, nil
// }

// func (h *UserHandler) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
// 	h.ensureUsers()
// 	h.users[in.Id] = in
// 	return in, nil
// }

// func (h *UserHandler) DeleteUser(ctx context.Context, in *pb.UserId) (*pb.Empty, error) {
// 	h.ensureUsers()
// 	delete(h.users, in.Id)
// 	return &pb.Empty{}, nil
// }
