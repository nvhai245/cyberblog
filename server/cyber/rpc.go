package main

import (
	"context"
	"github.com/nvhai245/cyberblog/server/cyber/controller"
	pb "github.com/nvhai245/cyberblog/server/cyber/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CyberServer struct {
	pb.UnimplementedCyberServer
}

func (*CyberServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	foundUser := controller.GetUserById(req.GetRequestorId(), req.GetUserId())
	return foundUser, nil
}

func (*CyberServer) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	foundUser := controller.GetAllUsers(req.GetAdminId())
	return foundUser, nil
}

func (*CyberServer) EditUser(ctx context.Context, req *pb.EditUserRequest) (*pb.EditUserResponse, error) {
	editedUser := controller.EditUser(req.GetRequestorEmail(), req.GetRequestorIsAdmin(), req.GetUser())
	return editedUser, nil
}

func (*CyberServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	deletedUser, err := controller.DeleteUser(req.GetRequestorEmail(), req.GetRequestorIsAdmin(), req.GetUserId())
	return deletedUser, err
}

func (*CyberServer) AddNewPost(ctx context.Context, req *pb.AddNewPostRequest) (*pb.AddNewPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewPost not implemented")
}

func (*CyberServer) EditPost(ctx context.Context, req *pb.EditPostRequest) (*pb.EditPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditPost not implemented")
}

func (*CyberServer) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}

func (*CyberServer) PublishPost(ctx context.Context, req *pb.PublishPostRequest) (*pb.PublishPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishPost not implemented")
}

func (*CyberServer) UnPublishPost(ctx context.Context, req *pb.UnPublishPostRequest) (*pb.UnPublishPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnPublishPost not implemented")
}

func (*CyberServer) UpVotePost(ctx context.Context, req *pb.UpVotePostRequest) (*pb.UpVotePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpVotePost not implemented")
}

func (*CyberServer) GetPostById(ctx context.Context, req *pb.GetPostByIdRequest) (*pb.GetPostByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostById not implemented")
}

func (*CyberServer) GetUserPublishedPosts(ctx context.Context, req *pb.GetUserPublishedPostsRequest) (*pb.GetUserPublishedPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPublishedPosts not implemented")
}

func (*CyberServer) GetUserAllPosts(ctx context.Context, req *pb.GetUserAllPostsRequest) (*pb.GetUserAllPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAllPosts not implemented")
}

func (*CyberServer) GetUserUnPublishedPosts(ctx context.Context, req *pb.GetUserUnPublishedPostsRequest) (*pb.GetUserUnPublishedPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserUnPublishedPosts not implemented")
}

func (*CyberServer) GetCategoryPosts(ctx context.Context, req *pb.GetCategoryPostsRequest) (*pb.GetCategoryPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryPosts not implemented")
}

func (*CyberServer) AddNewComment(ctx context.Context, req *pb.AddNewCommentRequest) (*pb.AddNewCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewComment not implemented")
}

func (*CyberServer) EditComment(ctx context.Context, req *pb.EditCommentRequest) (*pb.EditCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditComment not implemented")
}

func (*CyberServer) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}

func (*CyberServer) UpVoteComment(ctx context.Context, req *pb.UpVoteCommentRequest) (*pb.UpVoteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpVoteComment not implemented")
}

func (*CyberServer) DownVoteComment(ctx context.Context, req *pb.DownVoteCommentRequest) (*pb.DownVoteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownVoteComment not implemented")
}

func (*CyberServer) GetPostComments(ctx context.Context, req *pb.GetPostCommentsRequest) (*pb.GetPostCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostComments not implemented")
}

func (*CyberServer) GetUserComments(ctx context.Context, req *pb.GetUserCommentsRequest) (*pb.GetUserCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserComments not implemented")
}

func (*CyberServer) AddNewCategory(ctx context.Context, req *pb.AddNewCategoryRequest) (*pb.AddNewCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewCategory not implemented")
}

func (*CyberServer) EditCategory(ctx context.Context, req *pb.EditCategoryRequest) (*pb.EditCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditCategory not implemented")
}

func (*CyberServer) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}

func (*CyberServer) AddPostToCategory(ctx context.Context, req *pb.AddPostToCategoryRequest) (*pb.AddPostToCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPostToCategory not implemented")
}

func (*CyberServer) RemovePostFromCategory(ctx context.Context, req *pb.RemovePostFromCategoryRequest) (*pb.RemovePostFromCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePostFromCategory not implemented")
}

func (*CyberServer) GetAllCategories(ctx context.Context, req *pb.GetAllCategoriesRequest) (*pb.GetAllCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategories not implemented")
}

func (*CyberServer) GetPostCategories(ctx context.Context, req *pb.GetPostCategoriesRequest) (*pb.GetPostCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostCategories not implemented")
}
