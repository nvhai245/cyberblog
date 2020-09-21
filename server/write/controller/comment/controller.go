package commentController

import (
	commentModel "github.com/nvhai245/cyberblog/server/write/model/comment"
	pb "github.com/nvhai245/cyberblog/server/write/proto"
	"log"
)

func AddComment() {

}

func EditComment() {

}

func DeleteComment() {

}

func protoCommentToModelComment(protoComment *pb.Comment) *commentModel.Comment {
	modelComment := &commentModel.Comment{
		ID:        protoComment.GetId(),
		PostId:    protoComment.GetPostId(),
		AuthorId:  protoComment.GetAuthorId(),
		Content:   protoComment.GetContent(),
		UpVote:    protoComment.GetUpVote(),
		CreatedAt: protoComment.GetCreatedAt(),
		UpdatedAt: protoComment.GetUpdatedAt(),
	}
	return modelComment
}

func modelCommentToProtoComment(modelComment *commentModel.Comment) *pb.Comment {
	if modelComment == nil {
		log.Println("Error in controller.modelCommentToProtoComment(): modelComment is nil")
		return nil
	}
	protoComment := &pb.Comment{
		Id:        modelComment.ID,
		PostId:    modelComment.PostId,
		AuthorId:  modelComment.AuthorId,
		Content:   modelComment.Content,
		UpVote:    modelComment.UpVote,
		CreatedAt: modelComment.CreatedAt,
		UpdatedAt: modelComment.UpdatedAt,
	}
	return protoComment
}
