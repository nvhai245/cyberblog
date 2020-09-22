package commentController

import (
	"fmt"
	commentModel "github.com/nvhai245/cyberblog/server/write/model/comment"
	pb "github.com/nvhai245/cyberblog/server/write/proto"
	"log"
)

func AddComment(req *pb.AddNewCommentRequest) (*pb.AddNewCommentResponse, error) {
	log.Printf("commentController.AddComment(), [Request]: %+v", req)
	success, commentId := commentModel.Insert(protoCommentToModelComment(req.GetNewComment()))
	if success == false {
		err := fmt.Errorf("Error in commentController.AddComment(): failed to insert comment")
		log.Println(err)
		return nil, err
	}
	res := &pb.AddNewCommentResponse{
		Success:   true,
		CommentId: commentId,
	}
	// ******************************************************************************************
	log.Printf("commentController.AddComment(), [Response]: %+v", res)
	return res, nil
}

func EditComment(req *pb.EditCommentRequest) (*pb.EditCommentResponse, error) {
	log.Printf("commentController.EditComment(), [Request]: %+v", req)
	success, editedComment := commentModel.Update(protoCommentToModelComment(req.GetNewComment()))
	if success == false {
		err := fmt.Errorf("Error in commentController.EditComment(): failed to update comment")
		log.Println(err)
		return nil, err
	}
	res := &pb.EditCommentResponse{
		Success:       true,
		EditedComment: modelCommentToProtoComment(editedComment),
	}
	// ******************************************************************************************
	log.Printf("commentController.EditComment(), [Response]: %+v", res)
	return res, nil
}

func DeleteComment(req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	log.Printf("commentController.DeleteComment(), [Request]: %+v", req)
	success, deletedComment := commentModel.Delete(req.GetAuthorId(), req.GetCommentId())
	if success == false {
		err := fmt.Errorf("Error in commentController.DeleteComment(): failed to delete comment")
		log.Println(err)
		return nil, err
	}
	res := &pb.DeleteCommentResponse{
		Success:        true,
		DeletedComment: modelCommentToProtoComment(deletedComment),
	}
	// ******************************************************************************************
	log.Printf("commentController.DeleteComment(), [Response]: %+v", res)
	return res, nil
}

func UpVoteComment(req *pb.UpVoteCommentRequest) (*pb.UpVoteCommentResponse, error) {
	log.Printf("commentController.UpVoteComment(), [Request]: %+v", req)
	success, newUpVotes := commentModel.UpVote(req.GetCommentId())
	if success == false {
		err := fmt.Errorf("Error in commentController.UpVoteComment(): failed to upvote comment")
		log.Println(err)
		return nil, err
	}
	res := &pb.UpVoteCommentResponse{
		Success:    true,
		NewUpVotes: newUpVotes,
	}
	// ******************************************************************************************
	log.Printf("commentController.UpVoteComment(), [Response]: %+v", res)
	return res, nil
}

func DownVoteComment(req *pb.DownVoteCommentRequest) (*pb.DownVoteCommentResponse, error) {
	log.Printf("commentController.DownVoteComment(), [Request]: %+v", req)
	success, newUpVotes := commentModel.DownVote(req.GetCommentId())
	if success == false {
		err := fmt.Errorf("Error in commentController.DownVoteComment(): failed to downvote comment")
		log.Println(err)
		return nil, err
	}
	res := &pb.DownVoteCommentResponse{
		Success:    true,
		NewUpVotes: newUpVotes,
	}
	// ******************************************************************************************
	log.Printf("commentController.DownVoteComment(), [Response]: %+v", res)
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
