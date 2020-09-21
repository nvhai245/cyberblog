package postController

import (
	postModel "github.com/nvhai245/cyberblog/server/read/model/post"
	pb "github.com/nvhai245/cyberblog/server/read/proto"
	"log"
)

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
