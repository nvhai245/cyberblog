package commentController

import (
	"fmt"
	commentModel "github.com/nvhai245/cyberblog/server/read/model/comment"
	pb "github.com/nvhai245/cyberblog/server/read/proto"
	"log"
)

func GetPostComments(req *pb.GetPostCommentsRequest) (*pb.GetPostCommentsResponse, error) {
	log.Printf("commentController.GetPostComments(), [Request]: %+v", req)
	success, foundComments := commentModel.GetPostComments(req.GetPostId(), req.GetOffset(), req.GetLimit())
	if success == false {
		err := fmt.Errorf("Error in commentController.GetPostComments(): failed to get comments")
		log.Println(err)
		return nil, err
	}
	res := &pb.GetPostCommentsResponse{
		Success:       true,
		FoundComments: modelCommentSliceToProtoCommentSlice(foundComments),
	}
	// ******************************************************************************************
	log.Printf("commentController.GetPostComments(), [Response]: %+v", res)
	return res, nil
}

func GetUserComments(req *pb.GetUserCommentsRequest) (*pb.GetUserCommentsResponse, error) {
	log.Printf("commentController.GetUserComments(), [Request]: %+v", req)
	success, foundComments := commentModel.GetUserComments(req.GetAuthorId(), req.GetOffset(), req.GetLimit())
	if success == false {
		err := fmt.Errorf("Error in commentController.GetUserComments(): failed to get comments")
		log.Println(err)
		return nil, err
	}
	res := &pb.GetUserCommentsResponse{
		Success:       true,
		FoundComments: modelCommentSliceToProtoCommentSlice(foundComments),
	}
	// ******************************************************************************************
	log.Printf("commentController.GetUserComments(), [Response]: %+v", res)
	return res, nil
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

func modelCommentSliceToProtoCommentSlice(modelComments []*commentModel.Comment) []*pb.Comment {
	if len(modelComments) == 0 {
		log.Println("Error in controller.modelCommentToProtoComment(): modelComments is empty slice")
		return nil
	}
	var protoComments []*pb.Comment
	for _, modelComment := range modelComments {
		protoComment := &pb.Comment{
			Id:        modelComment.ID,
			PostId:    modelComment.PostId,
			AuthorId:  modelComment.AuthorId,
			Content:   modelComment.Content,
			UpVote:    modelComment.UpVote,
			CreatedAt: modelComment.CreatedAt,
			UpdatedAt: modelComment.UpdatedAt,
		}
		protoComments = append(protoComments, protoComment)
	}
	return protoComments
}
