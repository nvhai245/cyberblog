package main

import (
	"context"
	"github.com/nvhai245/cyberblog/server/write/controller/user"
	pb "github.com/nvhai245/cyberblog/server/write/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// WriteServer for rpc
type WriteServer struct {
	pb.UnimplementedWriteServer
}

// SaveUser method
func (*WriteServer) SaveUser(ctx context.Context, req *pb.SaveUserRequest) (*pb.SaveUserResponse, error) {
	res := userController.SaveUser(req)
	return res, nil
}

// EditUser method
func (*WriteServer) EditUser(ctx context.Context, req *pb.EditUserRequest) (*pb.EditUserResponse, error) {
	res := userController.EditUser(req)
	return res, nil
}

// DeleteUser method
func (*WriteServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	res, err := userController.DeleteUser(req)
	return res, err
}

func (*WriteServer) AddNewPost(ctx context.Context, req *pb.AddNewPostRequest) (*pb.AddNewPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewPost not implemented")
}

func (*WriteServer) EditPost(ctx context.Context, req *pb.EditPostRequest) (*pb.EditPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditPost not implemented")
}

func (*WriteServer) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}

func (*WriteServer) PublishPost(ctx context.Context, req *pb.PublishPostRequest) (*pb.PublishPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishPost not implemented")
}

func (*WriteServer) UnPublishPost(ctx context.Context, req *pb.UnPublishPostRequest) (*pb.UnPublishPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnPublishPost not implemented")
}

func (*WriteServer) UpVotePost(ctx context.Context, req *pb.UpVotePostRequest) (*pb.UpVotePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpVotePost not implemented")
}

func (*WriteServer) AddNewComment(ctx context.Context, req *pb.AddNewCommentRequest) (*pb.AddNewCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewComment not implemented")
}

func (*WriteServer) EditComment(ctx context.Context, req *pb.EditCommentRequest) (*pb.EditCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditComment not implemented")
}

func (*WriteServer) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}

func (*WriteServer) UpVoteComment(ctx context.Context, req *pb.UpVoteCommentRequest) (*pb.UpVoteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpVoteComment not implemented")
}

func (*WriteServer) DownVoteComment(ctx context.Context, req *pb.DownVoteCommentRequest) (*pb.DownVoteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownVoteComment not implemented")
}

func (*WriteServer) AddNewCategory(ctx context.Context, req *pb.AddNewCategoryRequest) (*pb.AddNewCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewCategory not implemented")
}

func (*WriteServer) EditCategory(ctx context.Context, req *pb.EditCategoryRequest) (*pb.EditCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditCategory not implemented")
}

func (*WriteServer) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}

func (*WriteServer) AddPostToCategory(ctx context.Context, req *pb.AddPostToCategoryRequest) (*pb.AddPostToCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPostToCategory not implemented")
}

func (*WriteServer) RemovePostFromCategory(ctx context.Context, req *pb.RemovePostFromCategoryRequest) (*pb.RemovePostFromCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePostFromCategory not implemented")
}
