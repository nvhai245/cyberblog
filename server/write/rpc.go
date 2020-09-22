package main

import (
	"context"
	categoryController "github.com/nvhai245/cyberblog/server/write/controller/category"
	commentController "github.com/nvhai245/cyberblog/server/write/controller/comment"
	postController "github.com/nvhai245/cyberblog/server/write/controller/post"
	"github.com/nvhai245/cyberblog/server/write/controller/user"
	pb "github.com/nvhai245/cyberblog/server/write/proto"
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
	return postController.AddPost(req)
}

func (*WriteServer) EditPost(ctx context.Context, req *pb.EditPostRequest) (*pb.EditPostResponse, error) {
	return postController.EditPost(req)
}

func (*WriteServer) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	return postController.DeletePost(req)
}

func (*WriteServer) PublishPost(ctx context.Context, req *pb.PublishPostRequest) (*pb.PublishPostResponse, error) {
	return postController.PublishPost(req)
}

func (*WriteServer) UnPublishPost(ctx context.Context, req *pb.UnPublishPostRequest) (*pb.UnPublishPostResponse, error) {
	return postController.UnPublishPost(req)
}

func (*WriteServer) UpVotePost(ctx context.Context, req *pb.UpVotePostRequest) (*pb.UpVotePostResponse, error) {
	return postController.UpVotePost(req)
}

func (*WriteServer) AddNewComment(ctx context.Context, req *pb.AddNewCommentRequest) (*pb.AddNewCommentResponse, error) {
	return commentController.AddComment(req)
}

func (*WriteServer) EditComment(ctx context.Context, req *pb.EditCommentRequest) (*pb.EditCommentResponse, error) {
	return commentController.EditComment(req)
}

func (*WriteServer) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	return commentController.DeleteComment(req)
}

func (*WriteServer) UpVoteComment(ctx context.Context, req *pb.UpVoteCommentRequest) (*pb.UpVoteCommentResponse, error) {
	return commentController.UpVoteComment(req)
}

func (*WriteServer) DownVoteComment(ctx context.Context, req *pb.DownVoteCommentRequest) (*pb.DownVoteCommentResponse, error) {
	return commentController.DownVoteComment(req)
}

func (*WriteServer) AddNewCategory(ctx context.Context, req *pb.AddNewCategoryRequest) (*pb.AddNewCategoryResponse, error) {
	return categoryController.AddNewCategory(req)
}

func (*WriteServer) EditCategory(ctx context.Context, req *pb.EditCategoryRequest) (*pb.EditCategoryResponse, error) {
	return categoryController.EditCategory(req)
}

func (*WriteServer) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error) {
	return categoryController.DeleteCategory(req)
}

func (*WriteServer) AddPostToCategory(ctx context.Context, req *pb.AddPostToCategoryRequest) (*pb.AddPostToCategoryResponse, error) {
	return categoryController.AddPostToCategory(req)
}

func (*WriteServer) RemovePostFromCategory(ctx context.Context, req *pb.RemovePostFromCategoryRequest) (*pb.RemovePostFromCategoryResponse, error) {
	return categoryController.RemovePostFromCategory(req)
}
