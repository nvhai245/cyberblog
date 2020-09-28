package postController

import (
	"fmt"
	postModel "github.com/nvhai245/cyberblog/server/write/model/post"
	pb "github.com/nvhai245/cyberblog/server/write/proto"
	"log"
)

func AddPost(req *pb.AddNewPostRequest) (*pb.AddNewPostResponse, error) {
	log.Printf("postController.AddPost(), [Request]: %+v", req)
	newModelPost := protoPostToModelPost(req.GetNewPost())
	success, postId := postModel.Insert(newModelPost)
	if success == false {
		err := fmt.Errorf("Error in postController.AddPost(): failed to insert post")
		log.Println(err)
		return nil, err
	}
	res := &pb.AddNewPostResponse{
		Success: true,
		NewPost: &pb.Post{
			Id: postId,
		},
	}
	// ******************************************************************************************
	log.Printf("postController.AddPost(), [Response]: %+v", res)
	return res, nil
}

func PublishPost(req *pb.PublishPostRequest) (*pb.PublishPostResponse, error) {
	log.Printf("postController.PublishPost(), [Request]: %+v", req)
	success, publishedPost := postModel.Publish(req.GetRequestorId(), req.GetPostId())
	if success == false {
		err := fmt.Errorf("Error in postController.PublishPost(): failed to publish post")
		log.Println(err)
		return nil, err
	}
	res := &pb.PublishPostResponse{
		Success:       true,
		PublishedPost: modelPostToProtoPost(publishedPost),
	}
	// ******************************************************************************************
	log.Printf("postController.PublishPost(), [Response]: %+v", res)
	return res, nil
}

func EditPost(req *pb.EditPostRequest) (*pb.EditPostResponse, error) {
	log.Printf("postController.EditPost(), [Request]: %+v", req)
	postToEdit := protoPostToModelPost(req.GetEditedPost())
	success, editedPost := postModel.Update(postToEdit)
	if success == false {
		err := fmt.Errorf("Error in postController.EditPost(): failed to edit post")
		log.Println(err)
		return nil, err
	}
	res := &pb.EditPostResponse{
		Success:    true,
		EditedPost: modelPostToProtoPost(editedPost),
	}
	// ******************************************************************************************
	log.Printf("postController.EditPost(), [Response]: %+v", res)
	return res, nil
}

func DeletePost(req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	log.Printf("postController.DeletePost(), [Request]: %+v", req)
	success, deletedPost := postModel.Delete(req.GetRequestorId(), req.GetPostId())
	if success == false {
		err := fmt.Errorf("Error in postController.DeletePost(): failed to delete post")
		log.Println(err)
		return nil, err
	}
	res := &pb.DeletePostResponse{
		Success:     true,
		DeletedPost: modelPostToProtoPost(deletedPost),
	}
	// ******************************************************************************************
	log.Printf("postController.DeletePost(), [Response]: %+v", res)
	return res, nil
}

func UnPublishPost(req *pb.UnPublishPostRequest) (*pb.UnPublishPostResponse, error) {
	log.Printf("postController.UnPublishPost(), [Request]: %+v", req)
	success, unPublishedPost := postModel.UnPublish(req.GetRequestorId(), req.GetPostId())
	if success == false {
		err := fmt.Errorf("Error in postController.UnPublishPost(): failed to unpublish post")
		log.Println(err)
		return nil, err
	}
	res := &pb.UnPublishPostResponse{
		Success:         true,
		UnPublishedPost: modelPostToProtoPost(unPublishedPost),
	}
	// ******************************************************************************************
	log.Printf("postController.UnPublishPost(), [Response]: %+v", res)
	return res, nil
}

func UpVotePost(req *pb.UpVotePostRequest) (*pb.UpVotePostResponse, error) {
	log.Printf("postController.UpVotePost(), [Request]: %+v", req)
	success, newUpVote := postModel.UpVote(req.GetUpVoterId(), req.GetPostId())
	if success == false {
		err := fmt.Errorf("Error in postController.UpVotePost(): failed to upvote post")
		log.Println(err)
		return nil, err
	}
	res := &pb.UpVotePostResponse{
		Success:   true,
		NewUpVote: newUpVote,
	}
	// ******************************************************************************************
	log.Printf("postController.UpVotePost(), [Response]: %+v", res)
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
