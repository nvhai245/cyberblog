package main

import (
	"context"
	"github.com/nvhai245/cyberblog/server/read/controller"
	pb "github.com/nvhai245/cyberblog/server/read/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	_ "log"
)

// WriteServer for rpc
type ReadServer struct {
	pb.UnimplementedReadServer
}

func (*ReadServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res := controller.GetUser(req)
	if res.GetUser() == nil {
		log.Println("controller.GetUser(): Can't find user with given email!")
	}
	return res, nil
}

func (*ReadServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	res := controller.GetUserById(req)
	if res.GetSuccess() == false {
		log.Println("controller.GetUserById(): Can't find user with given id!")
	}
	return res, nil
}

func (*ReadServer) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	res := controller.GetAllUsers(req)
	if res.GetSuccess() == false {
		log.Println("controller.GetAllUsers(): Can't get all users!")
	}
	return res, nil
}

func (*ReadServer) GetPostById(ctx context.Context, req *pb.GetPostByIdRequest) (*pb.GetPostByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostById not implemented")
}

func (*ReadServer) GetUserPublishedPosts(ctx context.Context, req *pb.GetUserPublishedPostsRequest) (*pb.GetUserPublishedPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPublishedPosts not implemented")
}

func (*ReadServer) GetUserAllPosts(ctx context.Context, req *pb.GetUserAllPostsRequest) (*pb.GetUserAllPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAllPosts not implemented")
}

func (*ReadServer) GetUserUnPublishedPosts(ctx context.Context, req *pb.GetUserUnPublishedPostsRequest) (*pb.GetUserUnPublishedPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserUnPublishedPosts not implemented")
}

func (*ReadServer) GetCategoryPosts(ctx context.Context, req *pb.GetCategoryPostsRequest) (*pb.GetCategoryPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryPosts not implemented")
}

func (*ReadServer) GetPostComments(ctx context.Context, req *pb.GetPostCommentsRequest) (*pb.GetPostCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostComments not implemented")
}

func (*ReadServer) GetUserComments(ctx context.Context, req *pb.GetUserCommentsRequest) (*pb.GetUserCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserComments not implemented")
}

func (*ReadServer) GetAllCategories(ctx context.Context, req *pb.GetAllCategoriesRequest) (*pb.GetAllCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategories not implemented")
}

func (*ReadServer) GetPostCategories(ctx context.Context, req *pb.GetPostCategoriesRequest) (*pb.GetPostCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostCategories not implemented")
}
