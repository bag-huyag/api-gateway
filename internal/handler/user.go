package handler

import (
	"context"

	pb "github.com/bag-huyag/api-gateway/proto/gen"
	"github.com/google/uuid"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	users map[string]*pb.User
}

func (h *UserHandler) ensureUsers() {
	if h.users == nil {
		h.users = make(map[string]*pb.User)
	}
}

func (h *UserHandler) GetUsers(ctx context.Context, in *pb.Empty) (*pb.UserList, error) {
	h.ensureUsers()

	users := make([]*pb.User, 0, len(h.users))
	for _, user := range h.users {
		users = append(users, user)
	}

	return &pb.UserList{Users: users}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, in *pb.UserId) (*pb.User, error) {
	h.ensureUsers()
	return h.users[in.Id], nil
}

func (h *UserHandler) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	h.ensureUsers()
	id := uuid.New().String()
	user := &pb.User{
		Id:    id,
		Name:  in.Name,
		Email: in.Email,
	}
	h.users[id] = user
	return user, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	h.ensureUsers()
	h.users[in.Id] = in
	return in, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, in *pb.UserId) (*pb.Empty, error) {
	h.ensureUsers()
	delete(h.users, in.Id)
	return &pb.Empty{}, nil
}
