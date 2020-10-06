package postController

import (
	postModel "github.com/nvhai245/cyberblog/server/read/model/post"
	pb "github.com/nvhai245/cyberblog/server/read/proto"
	"log"
)

func GetFeed(req *pb.GetFeedRequest) (*pb.GetFeedResponse, error) {
	log.Printf("postController.GetFeed(), [Request]: %+v", req)
	success, foundPosts, err := postModel.GetFeed(req.GetOffset(), req.GetLimit())
	if success == false || err != nil {
		return nil, err
	}
	res := &pb.GetFeedResponse{
		Success:    true,
		FoundPosts: modelPostSliceToProtoPostSlice(foundPosts),
	}
	// ******************************************************************************************
	log.Printf("postController.GetFeed(), [Response]: %+v", res)
	return res, nil
}

func GetPostById(req *pb.GetPostByIdRequest) (*pb.GetPostByIdResponse, error) {
	log.Printf("postController.GetPostById(), [Request]: %+v", req)
	success, foundPost, err := postModel.GetPostById(req.GetRequesterId(), req.GetPostId())
	if success == false || err != nil {
		return nil, err
	}
	res := &pb.GetPostByIdResponse{
		Success:   true,
		FoundPost: modelPostToProtoPost(foundPost),
	}
	// ******************************************************************************************
	log.Printf("postController.GetPostById(), [Response]: %+v", res)
	return res, nil
}

func GetUserPublishedPosts(req *pb.GetUserPublishedPostsRequest) (*pb.GetUserPublishedPostsResponse, error) {
	log.Printf("postController.GetUserPublishedPosts(), [Request]: %+v", req)
	success, foundPosts, err := postModel.GetUserPublishedPosts(req.GetRequesterId(), req.GetUserId(), req.GetOffset(), req.GetLimit())
	if success == false || err != nil {
		return nil, err
	}
	res := &pb.GetUserPublishedPostsResponse{
		Success:    true,
		FoundPosts: modelPostSliceToProtoPostSlice(foundPosts),
	}
	// ******************************************************************************************
	log.Printf("postController.GetUserPublishedPosts(), [Response]: %+v", res)
	return res, nil
}
func GetUserAllPosts(req *pb.GetUserAllPostsRequest) (*pb.GetUserAllPostsResponse, error) {
	log.Printf("postController.GetUserAllPosts(), [Request]: %+v", req)
	success, foundPosts, err := postModel.GetUserAllPosts(req.GetOwnerId(), req.GetOffset(), req.GetLimit())
	if success == false || err != nil {
		return nil, err
	}
	res := &pb.GetUserAllPostsResponse{
		Success:    true,
		FoundPosts: modelPostSliceToProtoPostSlice(foundPosts),
	}
	// ******************************************************************************************
	log.Printf("postController.GetUserAllPosts(), [Response]: %+v", res)
	return res, nil
}

func GetUserUnPublishedPosts(req *pb.GetUserUnPublishedPostsRequest) (*pb.GetUserUnPublishedPostsResponse, error) {
	log.Printf("postController.GetUserUnPublishedPosts(), [Request]: %+v", req)
	success, foundPosts, err := postModel.GetUserUnPublishedPosts(req.GetOwnerId(), req.GetOffset(), req.GetLimit())
	if success == false || err != nil {
		return nil, err
	}
	res := &pb.GetUserUnPublishedPostsResponse{
		Success:    true,
		FoundPosts: modelPostSliceToProtoPostSlice(foundPosts),
	}
	// ******************************************************************************************
	log.Printf("postController.GetUserUnPublishedPosts(), [Response]: %+v", res)
	return res, nil
}

func GetCategoryPosts(req *pb.GetCategoryPostsRequest) (*pb.GetCategoryPostsResponse, error) {
	log.Printf("postController.GetCategoryPosts(), [Request]: %+v", req)
	success, foundPosts, err := postModel.GetCategoryPosts(req.GetCategoryId(), req.GetOffset(), req.GetLimit())
	if success == false || err != nil {
		return nil, err
	}
	res := &pb.GetCategoryPostsResponse{
		Success:    true,
		FoundPosts: modelPostSliceToProtoPostSlice(foundPosts),
	}
	// ******************************************************************************************
	log.Printf("postController.GetCategoryPosts(), [Response]: %+v", res)
	return res, nil
}

func protoPostToModelPost(protoPost *pb.Post) *postModel.Post {
	modelPost := &postModel.Post{
		ID:          protoPost.GetId(),
		AuthorId:    protoPost.GetAuthorId(),
		ParentId:    protoPost.GetParentId(),
		Title:       protoPost.GetTitle(),
		Published:   protoPost.GetPublished(),
		UpVote:      protoPost.GetUpVote(),
		Content:     protoPost.GetContent(),
		CreatedAt:   protoPost.GetCreatedAt(),
		UpdatedAt:   protoPost.GetUpdatedAt(),
		PublishedAt: protoPost.GetPublishedAt(),
	}
	return modelPost
}

func modelPostToProtoPost(modelPost *postModel.Post) *pb.Post {
	if modelPost == nil {
		log.Println("Error in controller.modelPostToProtoPost(): modelPost is nil")
		return nil
	}
	protoPost := &pb.Post{
		Id:          modelPost.ID,
		AuthorId:    modelPost.AuthorId,
		ParentId:    modelPost.ParentId,
		Title:       modelPost.Title,
		Published:   modelPost.Published,
		UpVote:      modelPost.UpVote,
		Content:     modelPost.Content,
		CreatedAt:   modelPost.CreatedAt,
		UpdatedAt:   modelPost.UpdatedAt,
		PublishedAt: modelPost.PublishedAt,
	}
	return protoPost
}

func modelPostSliceToProtoPostSlice(modelPostSlice []*postModel.Post) []*pb.Post {
	if len(modelPostSlice) == 0 {
		log.Println("Error in controller.modelPostToProtoPost(): modelPostSlice is empty")
		return nil
	}
	var protoPostSlice []*pb.Post
	for _, modelPost := range modelPostSlice {
		protoPost := &pb.Post{
			Id:          modelPost.ID,
			AuthorId:    modelPost.AuthorId,
			ParentId:    modelPost.ParentId,
			Title:       modelPost.Title,
			Published:   modelPost.Published,
			UpVote:      modelPost.UpVote,
			Content:     modelPost.Content,
			CreatedAt:   modelPost.CreatedAt,
			UpdatedAt:   modelPost.UpdatedAt,
			PublishedAt: modelPost.PublishedAt,
		}
		protoPostSlice = append(protoPostSlice, protoPost)
	}
	return protoPostSlice
}
