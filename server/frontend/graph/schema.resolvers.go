package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	commentController "github.com/nvhai245/cyberblog/server/frontend/controller/comment"

	categoryController "github.com/nvhai245/cyberblog/server/frontend/controller/category"
	postController "github.com/nvhai245/cyberblog/server/frontend/controller/post"
	userController "github.com/nvhai245/cyberblog/server/frontend/controller/user"
	"github.com/nvhai245/cyberblog/server/frontend/graph/generated"
	"github.com/nvhai245/cyberblog/server/frontend/graph/model"
)

func (r *mutationResolver) Register(ctx context.Context, email string, password string) (*model.AuthResponse, error) {
	return userController.Register(ctx, email, password)
}

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.AuthResponse, error) {
	return userController.Login(ctx, email, password)
}

func (r *mutationResolver) GetUserByID(ctx context.Context, requestorID int, userID int) (*model.GetUserByIDResponse, error) {
	return userController.GetUserByID(ctx, requestorID, userID)
}

func (r *mutationResolver) GetAllUsers(ctx context.Context, adminID int) (*model.GetAllUsersResponse, error) {
	return userController.GetAllUsers(ctx, adminID)
}

func (r *mutationResolver) EditUser(ctx context.Context, userID int, editedUser model.EditedUser) (*model.EditUserResponse, error) {
	return userController.EditUser(ctx, userID, editedUser)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, adminID int, userID int) (*model.DeleteUserResponse, error) {
	return userController.DeleteUser(ctx, adminID, userID)
}

func (r *mutationResolver) GetPostByID(ctx context.Context, requesterID int, postID int) (*model.GetPostByIDResponse, error) {
	return postController.GetPostByID(ctx, requesterID, postID)
}

func (r *mutationResolver) GetUserAllPosts(ctx context.Context, ownerID int, offset int, limit int) (*model.GetPostsResponse, error) {
	return postController.GetUserAllPosts(ctx, ownerID, offset, limit)
}

func (r *mutationResolver) GetUserPublishedPosts(ctx context.Context, requesterID int, userID int, offset int, limit int) (*model.GetPostsResponse, error) {
	return postController.GetUserPublishedPosts(ctx, requesterID, userID, offset, limit)
}

func (r *mutationResolver) GetUserUnPublishedPosts(ctx context.Context, ownerID int, offset int, limit int) (*model.GetPostsResponse, error) {
	return postController.GetUserUnPublishedPosts(ctx, ownerID, offset, limit)
}

func (r *mutationResolver) GetCategoryPosts(ctx context.Context, categoryID int, offset int, limit int) (*model.GetPostsResponse, error) {
	return postController.GetCategoryPosts(ctx, categoryID, offset, limit)
}

func (r *mutationResolver) AddNewPost(ctx context.Context, newPost model.NewPost) (*model.GetPostByIDResponse, error) {
	return postController.AddNewPost(ctx, newPost)
}

func (r *mutationResolver) EditPost(ctx context.Context, newPost model.NewPost) (*model.GetPostByIDResponse, error) {
	return postController.EditPost(ctx, newPost)
}

func (r *mutationResolver) DeletePost(ctx context.Context, requesterID int, postID int) (*model.GetPostByIDResponse, error) {
	return postController.DeletePost(ctx, requesterID, postID)
}

func (r *mutationResolver) PublishPost(ctx context.Context, requesterID int, postID int) (*model.GetPostByIDResponse, error) {
	return postController.PublishPost(ctx, requesterID, postID)
}

func (r *mutationResolver) UnPublishPost(ctx context.Context, requesterID int, postID int) (*model.GetPostByIDResponse, error) {
	return postController.UnPublishPost(ctx, requesterID, postID)
}

func (r *mutationResolver) UpVotePost(ctx context.Context, upVoterID int, postID int) (*model.UpVotes, error) {
	return postController.UpVotePost(ctx, upVoterID, postID)
}

func (r *mutationResolver) GetPostComments(ctx context.Context, postID int, offset int, limit int) (*model.GetCommentsResponse, error) {
	return commentController.GetPostComments(ctx, postID, offset, limit)
}

func (r *mutationResolver) GetUserComments(ctx context.Context, authorID int, offset int, limit int) (*model.GetCommentsResponse, error) {
	return commentController.GetUserComments(ctx, authorID, offset, limit)
}

func (r *mutationResolver) AddComment(ctx context.Context, newComment model.NewComment) (*model.AddNewCommentResponse, error) {
	return commentController.AddComment(ctx, newComment)
}

func (r *mutationResolver) EditComment(ctx context.Context, newComment model.NewComment) (*model.AddNewCommentResponse, error) {
	return commentController.EditComment(ctx, newComment)
}

func (r *mutationResolver) DeleteComment(ctx context.Context, authorID int, commentID int) (*model.AddNewCommentResponse, error) {
	return commentController.DeleteComment(ctx, authorID, commentID)
}

func (r *mutationResolver) UpVoteComment(ctx context.Context, commentID int) (*model.UpVotes, error) {
	return commentController.UpVoteComment(ctx, commentID)
}

func (r *mutationResolver) DownVoteComment(ctx context.Context, commentID int) (*model.UpVotes, error) {
	return commentController.DownVoteComment(ctx, commentID)
}

func (r *mutationResolver) GetAllCategories(ctx context.Context, requesterID int) (*model.GetCategoriesResponse, error) {
	return categoryController.GetAllCategories(ctx, requesterID)
}

func (r *mutationResolver) GetPostCategories(ctx context.Context, postID int) (*model.GetCategoriesResponse, error) {
	return categoryController.GetPostCategories(ctx, postID)
}

func (r *mutationResolver) AddNewCategory(ctx context.Context, newCategory model.NewCategory) (*model.GetCategoryResponse, error) {
	return categoryController.AddNewCategory(ctx, newCategory)
}

func (r *mutationResolver) EditCategory(ctx context.Context, newCategory model.NewCategory) (*model.GetCategoryResponse, error) {
	return categoryController.EditCategory(ctx, newCategory)
}

func (r *mutationResolver) DeleteCategory(ctx context.Context, categoryID int) (*model.GetCategoryResponse, error) {
	return categoryController.DeleteCategory(ctx, categoryID)
}

func (r *mutationResolver) AddPostToCategory(ctx context.Context, categoryID int, postID int) (*model.PostCategoryResponse, error) {
	return categoryController.AddPostToCategory(ctx, categoryID, postID)
}

func (r *mutationResolver) RemovePostFromCategory(ctx context.Context, categoryID int, postID int) (*model.PostCategoryResponse, error) {
	return categoryController.RemovePostFromCategory(ctx, categoryID, postID)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return []*model.User{}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
