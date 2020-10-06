package main

import (
	"context"
	categoryController "github.com/nvhai245/cyberblog/server/read/controller/category"
	commentController "github.com/nvhai245/cyberblog/server/read/controller/comment"
	postController "github.com/nvhai245/cyberblog/server/read/controller/post"
	"github.com/nvhai245/cyberblog/server/read/controller/user"
	pb "github.com/nvhai245/cyberblog/server/read/proto"
	"log"
	_ "log"
)

// WriteServer for rpc
type ReadServer struct {
	pb.UnimplementedReadServer
}

func (*ReadServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res := userController.GetUser(req)
	if res.GetUser() == nil {
		log.Println("controller.GetUser(): Can't find user with given email!")
	}
	return res, nil
}

func (*ReadServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	res := userController.GetUserById(req)
	if res.GetSuccess() == false {
		log.Println("controller.GetUserById(): Can't find user with given id!")
	}
	return res, nil
}

func (*ReadServer) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	res := userController.GetAllUsers(req)
	if res.GetSuccess() == false {
		log.Println("controller.GetAllUsers(): Can't get all users!")
	}
	return res, nil
}

func (*ReadServer) GetFeed(ctx context.Context, req *pb.GetFeedRequest) (*pb.GetFeedResponse, error) {
	return postController.GetFeed(req)
}

func (*ReadServer) GetPostById(ctx context.Context, req *pb.GetPostByIdRequest) (*pb.GetPostByIdResponse, error) {
	return postController.GetPostById(req)
}

func (*ReadServer) GetUserPublishedPosts(ctx context.Context, req *pb.GetUserPublishedPostsRequest) (*pb.GetUserPublishedPostsResponse, error) {
	return postController.GetUserPublishedPosts(req)
}

func (*ReadServer) GetUserAllPosts(ctx context.Context, req *pb.GetUserAllPostsRequest) (*pb.GetUserAllPostsResponse, error) {
	return postController.GetUserAllPosts(req)
}

func (*ReadServer) GetUserUnPublishedPosts(ctx context.Context, req *pb.GetUserUnPublishedPostsRequest) (*pb.GetUserUnPublishedPostsResponse, error) {
	return postController.GetUserUnPublishedPosts(req)
}

func (*ReadServer) GetCategoryPosts(ctx context.Context, req *pb.GetCategoryPostsRequest) (*pb.GetCategoryPostsResponse, error) {
	return postController.GetCategoryPosts(req)
}

func (*ReadServer) GetPostComments(ctx context.Context, req *pb.GetPostCommentsRequest) (*pb.GetPostCommentsResponse, error) {
	return commentController.GetPostComments(req)
}

func (*ReadServer) GetUserComments(ctx context.Context, req *pb.GetUserCommentsRequest) (*pb.GetUserCommentsResponse, error) {
	return commentController.GetUserComments(req)
}

func (*ReadServer) GetAllCategories(ctx context.Context, req *pb.GetAllCategoriesRequest) (*pb.GetAllCategoriesResponse, error) {
	return categoryController.GetAllCategories(req)
}

func (*ReadServer) GetPostCategories(ctx context.Context, req *pb.GetPostCategoriesRequest) (*pb.GetPostCategoriesResponse, error) {
	return categoryController.GetPostCategories(req)
}
