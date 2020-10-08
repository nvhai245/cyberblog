package postController

import (
	"context"
	"fmt"
	authPb "github.com/nvhai245/cyberblog/server/auth/proto"
	"github.com/nvhai245/cyberblog/server/frontend/connection"
	"github.com/nvhai245/cyberblog/server/frontend/graph/model"
	"github.com/nvhai245/cyberblog/server/frontend/helper"
	readPb "github.com/nvhai245/cyberblog/server/read/proto"
	writePb "github.com/nvhai245/cyberblog/server/write/proto"
	"log"
	"time"
)

func GetFeed(ctx context.Context, offset int, limit int) (*model.GetPostsResponse, error) {
	result, err := connection.ReadClient.GetFeed(ctx, &readPb.GetFeedRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false || result.GetFoundPosts() == nil {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	var posts []*model.Post
	for _, post := range result.GetFoundPosts() {
		posts = append(posts, readPostToGraphPost(post))
	}
	return &model.GetPostsResponse{
		Message: "get feeds successful!",
		Posts:   posts,
	}, nil
}

func GetPostByID(ctx context.Context, requesterID int, postID int) (*model.GetPostByIDResponse, error) {
	//session := helper.GetSession(ctx, "auth")
	//token := fmt.Sprintf("%v", session.Values["token"])
	//checkResponse, err := connection.AuthClient.CheckToken(context.Background(), &authPb.CheckTokenRequest{Token: token})
	//if err != nil {
	//	log.Println("Error in rpc AuthClient.CheckToken(): ", err)
	//	return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	//}
	//// No need to check auth
	//_ = checkResponse
	result, err := connection.ReadClient.GetPostById(ctx, &readPb.GetPostByIdRequest{
		RequesterId: int32(requesterID),
		PostId:      int32(postID),
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false || result.GetFoundPost() == nil {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	return &model.GetPostByIDResponse{
		Message: "get post successful!",
		Post:    readPostToGraphPost(result.GetFoundPost()),
	}, nil
}

func GetUserAllPosts(ctx context.Context, ownerID int, offset int, limit int) (*model.GetPostsResponse, error) {
	session := helper.GetSession(ctx, "auth")
	token := fmt.Sprintf("%v", session.Values["token"])
	checkResponse, err := connection.AuthClient.CheckToken(context.Background(), &authPb.CheckTokenRequest{Token: token})
	if err != nil {
		log.Println("Error in rpc AuthClient.CheckToken(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	if !checkResponse.GetValid() {
		return nil, fmt.Errorf("UNAUTHORIZED!")
	}
	result, err := connection.ReadClient.GetUserAllPosts(ctx, &readPb.GetUserAllPostsRequest{
		OwnerId: int32(ownerID),
		Offset:  int32(offset),
		Limit:   int32(limit),
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false || result.GetFoundPosts() == nil {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	var posts []*model.Post
	for _, post := range result.GetFoundPosts() {
		posts = append(posts, readPostToGraphPost(post))
	}
	return &model.GetPostsResponse{
		Message: "get user all post successful!",
		Posts:   posts,
	}, nil
}

func GetUserPublishedPosts(ctx context.Context, requesterID int, userID int, offset int, limit int) (*model.GetPostsResponse, error) {
	//session := helper.GetSession(ctx, "auth")
	//token := fmt.Sprintf("%v", session.Values["token"])
	//checkResponse, err := connection.AuthClient.CheckToken(context.Background(), &authPb.CheckTokenRequest{Token: token})
	//if err != nil {
	//	log.Println("Error in rpc AuthClient.CheckToken(): ", err)
	//	return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	//}
	//_ = checkResponse
	result, err := connection.ReadClient.GetUserPublishedPosts(ctx, &readPb.GetUserPublishedPostsRequest{
		RequesterId: int32(requesterID),
		UserId:      int32(userID),
		Offset:      int32(offset),
		Limit:       int32(limit),
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false || result.GetFoundPosts() == nil {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	var posts []*model.Post
	for _, post := range result.GetFoundPosts() {
		posts = append(posts, readPostToGraphPost(post))
	}
	return &model.GetPostsResponse{
		Message: "get user published post successful!",
		Posts:   posts,
	}, nil
}

func GetUserUnPublishedPosts(ctx context.Context, ownerID int, offset int, limit int) (*model.GetPostsResponse, error) {
	session := helper.GetSession(ctx, "auth")
	token := fmt.Sprintf("%v", session.Values["token"])
	checkResponse, err := connection.AuthClient.CheckToken(context.Background(), &authPb.CheckTokenRequest{Token: token})
	if err != nil {
		log.Println("Error in rpc AuthClient.CheckToken(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	if !checkResponse.GetValid() {
		return nil, fmt.Errorf("UNAUTHORIZED!")
	}
	result, err := connection.ReadClient.GetUserUnPublishedPosts(ctx, &readPb.GetUserUnPublishedPostsRequest{
		OwnerId: int32(ownerID),
		Offset:  int32(offset),
		Limit:   int32(limit),
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false || result.GetFoundPosts() == nil {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	var posts []*model.Post
	for _, post := range result.GetFoundPosts() {
		posts = append(posts, readPostToGraphPost(post))
	}
	return &model.GetPostsResponse{
		Message: "get user unpublished post successful!",
		Posts:   posts,
	}, nil
}

func GetCategoryPosts(ctx context.Context, categoryID int, offset int, limit int) (*model.GetPostsResponse, error) {
	session := helper.GetSession(ctx, "auth")
	token := fmt.Sprintf("%v", session.Values["token"])
	log.Println(token)
	checkResponse, err := connection.AuthClient.CheckToken(context.Background(), &authPb.CheckTokenRequest{Token: token})
	if err != nil {
		log.Println("Error in rpc AuthClient.CheckToken(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	_ = checkResponse
	result, err := connection.ReadClient.GetCategoryPosts(ctx, &readPb.GetCategoryPostsRequest{
		CategoryId: int32(categoryID),
		Offset:     int32(offset),
		Limit:      int32(limit),
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false || result.GetFoundPosts() == nil {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	var posts []*model.Post
	for _, post := range result.GetFoundPosts() {
		posts = append(posts, readPostToGraphPost(post))
	}
	return &model.GetPostsResponse{
		Message: "get category post successful!",
		Posts:   posts,
	}, nil
}

func AddNewPost(ctx context.Context, newPost model.NewPost) (*model.GetPostByIDResponse, error) {
	session := helper.GetSession(ctx, "auth")
	token := fmt.Sprintf("%v", session.Values["token"])
	checkResponse, err := connection.AuthClient.CheckToken(context.Background(), &authPb.CheckTokenRequest{Token: token})
	if err != nil {
		log.Println("Error in rpc AuthClient.CheckToken(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	if !checkResponse.GetValid() {
		return nil, fmt.Errorf("UNAUTHORIZED!")
	}
	result, err := connection.WriteClient.AddNewPost(ctx, &writePb.AddNewPostRequest{NewPost: &writePb.Post{
		AuthorId: int32(newPost.AuthorID),
		ParentId: int32(newPost.ParentID),
		Title:    newPost.Title,
		Content:  newPost.Content,
	}})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false || result.GetNewPost() == nil {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	return &model.GetPostByIDResponse{
		Message: "add new post successfully!",
		Post:    writePostToGraphPost(result.GetNewPost()),
	}, nil
}

func EditPost(ctx context.Context, newPost model.NewPost) (*model.GetPostByIDResponse, error) {
	session := helper.GetSession(ctx, "auth")
	token := fmt.Sprintf("%v", session.Values["token"])
	checkResponse, err := connection.AuthClient.CheckToken(context.Background(), &authPb.CheckTokenRequest{Token: token})
	if err != nil {
		log.Println("Error in rpc AuthClient.CheckToken(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	if !checkResponse.GetValid() {
		return nil, fmt.Errorf("UNAUTHORIZED!")
	}
	result, err := connection.WriteClient.EditPost(ctx, &writePb.EditPostRequest{EditedPost: &writePb.Post{
		Id:        int32(newPost.ID),
		AuthorId:  int32(newPost.AuthorID),
		ParentId:  int32(newPost.ParentID),
		Title:     newPost.Title,
		Content:   newPost.Content,
		UpdatedAt: int32(time.Now().Unix()),
	}})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false || result.GetEditedPost() == nil {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	return &model.GetPostByIDResponse{
		Message: "add new post successfully!",
		Post:    writePostToGraphPost(result.GetEditedPost()),
	}, nil
}

func DeletePost(ctx context.Context, requesterID int, postID int) (*model.GetPostByIDResponse, error) {
	session := helper.GetSession(ctx, "auth")
	token := fmt.Sprintf("%v", session.Values["token"])
	checkResponse, err := connection.AuthClient.CheckToken(context.Background(), &authPb.CheckTokenRequest{Token: token})
	if err != nil {
		log.Println("Error in rpc AuthClient.CheckToken(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	if !checkResponse.GetValid() {
		return nil, fmt.Errorf("UNAUTHORIZED!")
	}
	result, err := connection.WriteClient.DeletePost(ctx, &writePb.DeletePostRequest{
		RequestorId: int32(requesterID),
		PostId:      int32(postID),
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false || result.GetDeletedPost() == nil {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	return &model.GetPostByIDResponse{
		Message: "add new post successfully!",
		Post:    writePostToGraphPost(result.GetDeletedPost()),
	}, nil
}

func PublishPost(ctx context.Context, requesterID int, postID int) (*model.GetPostByIDResponse, error) {
	session := helper.GetSession(ctx, "auth")
	token := fmt.Sprintf("%v", session.Values["token"])
	checkResponse, err := connection.AuthClient.CheckToken(context.Background(), &authPb.CheckTokenRequest{Token: token})
	if err != nil {
		log.Println("Error in rpc AuthClient.CheckToken(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	if !checkResponse.GetValid() {
		return nil, fmt.Errorf("UNAUTHORIZED!")
	}
	result, err := connection.WriteClient.PublishPost(ctx, &writePb.PublishPostRequest{
		RequestorId: int32(requesterID),
		PostId:      int32(postID),
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false || result.GetPublishedPost() == nil {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	return &model.GetPostByIDResponse{
		Message: "publish post successfully!",
		Post:    writePostToGraphPost(result.GetPublishedPost()),
	}, nil
}

func UnPublishPost(ctx context.Context, requesterID int, postID int) (*model.GetPostByIDResponse, error) {
	session := helper.GetSession(ctx, "auth")
	token := fmt.Sprintf("%v", session.Values["token"])
	checkResponse, err := connection.AuthClient.CheckToken(context.Background(), &authPb.CheckTokenRequest{Token: token})
	if err != nil {
		log.Println("Error in rpc AuthClient.CheckToken(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	if !checkResponse.GetValid() {
		return nil, fmt.Errorf("UNAUTHORIZED!")
	}
	result, err := connection.WriteClient.UnPublishPost(ctx, &writePb.UnPublishPostRequest{
		RequestorId: int32(requesterID),
		PostId:      int32(postID),
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false || result.GetUnPublishedPost() == nil {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	return &model.GetPostByIDResponse{
		Message: "unPublish post successfully!",
		Post:    writePostToGraphPost(result.GetUnPublishedPost()),
	}, nil
}

func UpVotePost(ctx context.Context, upVoterID int, postID int) (*model.UpVotes, error) {
	session := helper.GetSession(ctx, "auth")
	token := fmt.Sprintf("%v", session.Values["token"])
	checkResponse, err := connection.AuthClient.CheckToken(context.Background(), &authPb.CheckTokenRequest{Token: token})
	if err != nil {
		log.Println("Error in rpc AuthClient.CheckToken(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	if !checkResponse.GetValid() {
		return nil, fmt.Errorf("UNAUTHORIZED!")
	}
	result, err := connection.WriteClient.UpVotePost(ctx, &writePb.UpVotePostRequest{
		UpVoterId: int32(upVoterID),
		PostId:    int32(postID),
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	return &model.UpVotes{
		Message:    "unPublish post successfully!",
		NewUpVotes: int(result.GetNewUpVote()),
	}, nil
}

func readPostToGraphPost(readPost *readPb.Post) *model.Post {
	if readPost == nil {
		log.Println("controller.readPostToGraphPost(): readPost is nil")
		return nil
	}
	graphPost := &model.Post{
		ID:          int(readPost.GetId()),
		AuthorID:    int(readPost.GetAuthorId()),
		ParentID:    int(readPost.GetParentId()),
		Title:       readPost.GetTitle(),
		Published:   readPost.GetPublished(),
		UpVote:      int(readPost.GetUpVote()),
		Content:     readPost.GetContent(),
		CreatedAt:   int(readPost.GetCreatedAt()),
		UpdatedAt:   int(readPost.GetUpdatedAt()),
		PublishedAt: int(readPost.GetPublishedAt()),
	}
	return graphPost
}

func graphPostToReadPost(graphPost *model.Post) *readPb.Post {
	if graphPost == nil {
		log.Println("controller.graphPostToReadPost(): graphPost is nil")
		return nil
	}
	readPost := &readPb.Post{
		Id:          int32(graphPost.ID),
		AuthorId:    int32(graphPost.AuthorID),
		ParentId:    int32(graphPost.ParentID),
		Title:       graphPost.Title,
		Published:   graphPost.Published,
		UpVote:      int32(graphPost.UpVote),
		Content:     graphPost.Content,
		CreatedAt:   int32(graphPost.CreatedAt),
		UpdatedAt:   int32(graphPost.UpdatedAt),
		PublishedAt: int32(graphPost.PublishedAt),
	}
	return readPost
}

func writePostToGraphPost(writePost *writePb.Post) *model.Post {
	if writePost == nil {
		log.Println("controller.readPostToGraphPost(): readPost is nil")
		return nil
	}
	graphPost := &model.Post{
		ID:          int(writePost.GetId()),
		AuthorID:    int(writePost.GetAuthorId()),
		ParentID:    int(writePost.GetParentId()),
		Title:       writePost.GetTitle(),
		Published:   writePost.GetPublished(),
		UpVote:      int(writePost.GetUpVote()),
		Content:     writePost.GetContent(),
		CreatedAt:   int(writePost.GetCreatedAt()),
		UpdatedAt:   int(writePost.GetUpdatedAt()),
		PublishedAt: int(writePost.GetPublishedAt()),
	}
	return graphPost
}

func graphPostToWritePost(graphPost *model.Post) *writePb.Post {
	if graphPost == nil {
		log.Println("controller.graphPostToReadPost(): graphPost is nil")
		return nil
	}
	writePost := &writePb.Post{
		Id:          int32(graphPost.ID),
		AuthorId:    int32(graphPost.AuthorID),
		ParentId:    int32(graphPost.ParentID),
		Title:       graphPost.Title,
		Published:   graphPost.Published,
		UpVote:      int32(graphPost.UpVote),
		Content:     graphPost.Content,
		CreatedAt:   int32(graphPost.CreatedAt),
		UpdatedAt:   int32(graphPost.UpdatedAt),
		PublishedAt: int32(graphPost.PublishedAt),
	}
	return writePost
}
