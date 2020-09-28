package commentController

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

func GetPostComments(ctx context.Context, postID int, offset int, limit int) (*model.GetCommentsResponse, error) {
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

	result, err := connection.ReadClient.GetPostComments(ctx, &readPb.GetPostCommentsRequest{
		PostId: int32(postID),
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}

	return &model.GetCommentsResponse{
		Message:  "get post comments successful!",
		Comments: readCommentsToGraphComments(result.GetFoundComments()),
	}, nil
}

func GetUserComments(ctx context.Context, authorID int, offset int, limit int) (*model.GetCommentsResponse, error) {
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
	result, err := connection.ReadClient.GetUserComments(ctx, &readPb.GetUserCommentsRequest{
		AuthorId: int32(authorID),
		Offset:   int32(offset),
		Limit:    int32(limit),
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}

	return &model.GetCommentsResponse{
		Message:  "get user comments successful!",
		Comments: readCommentsToGraphComments(result.GetFoundComments()),
	}, nil
}

func AddComment(ctx context.Context, newComment model.NewComment) (*model.AddNewCommentResponse, error) {
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

	result, err := connection.WriteClient.AddNewComment(ctx, &writePb.AddNewCommentRequest{
		NewComment: &writePb.Comment{
			PostId:   int32(newComment.PostID),
			AuthorId: int32(newComment.AuthorID),
			Content:  newComment.Content,
		},
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	return &model.AddNewCommentResponse{
		Message:   "add comment successful!",
		CommentID: int(result.GetCommentId()),
	}, nil
}

func EditComment(ctx context.Context, newComment model.NewComment) (*model.AddNewCommentResponse, error) {
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

	result, err := connection.WriteClient.EditComment(ctx, &writePb.EditCommentRequest{
		NewComment: &writePb.Comment{
			Id:        int32(newComment.ID),
			PostId:    int32(newComment.PostID),
			AuthorId:  int32(newComment.AuthorID),
			Content:   newComment.Content,
			UpdatedAt: int32(time.Now().Unix()),
		},
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	return &model.AddNewCommentResponse{
		Message:   "edit comment successful!",
		CommentID: int(result.GetEditedComment().GetId()),
	}, nil
}

func DeleteComment(ctx context.Context, authorID int, commentID int) (*model.AddNewCommentResponse, error) {
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
	result, err := connection.WriteClient.DeleteComment(ctx, &writePb.DeleteCommentRequest{
		AuthorId:  int32(authorID),
		CommentId: int32(commentID),
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	return &model.AddNewCommentResponse{
		Message:   "delete comment successful!",
		CommentID: int(result.GetDeletedComment().GetId()),
	}, nil
}

func UpVoteComment(ctx context.Context, commentID int) (*model.UpVotes, error) {
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
	if err != nil {
		log.Println("Error in rpc AuthClient.CheckToken(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	if !checkResponse.GetValid() {
		return nil, fmt.Errorf("UNAUTHORIZED!")
	}
	result, err := connection.WriteClient.UpVoteComment(ctx, &writePb.UpVoteCommentRequest{
		CommentId: int32(commentID),
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	return &model.UpVotes{
		Message:    "up vote comment successful!",
		NewUpVotes: int(result.GetNewUpVotes()),
	}, nil
}

func DownVoteComment(ctx context.Context, commentID int) (*model.UpVotes, error) {
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
	if err != nil {
		log.Println("Error in rpc AuthClient.CheckToken(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	if !checkResponse.GetValid() {
		return nil, fmt.Errorf("UNAUTHORIZED!")
	}
	result, err := connection.WriteClient.DownVoteComment(ctx, &writePb.DownVoteCommentRequest{
		CommentId: int32(commentID),
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	return &model.UpVotes{
		Message:    "down vote comment successful!",
		NewUpVotes: int(result.GetNewUpVotes()),
	}, nil
}

func readCommentsToGraphComments(readComments []*readPb.Comment) []*model.Comment {
	if readComments == nil || len(readComments) == 0 {
		log.Println("CommentController.readCommentsToGraphComments(): readComments is empty")
		return nil
	}
	var Cmts []*model.Comment
	for _, readCmt := range readComments {
		cmt := &model.Comment{
			ID:        int(readCmt.GetId()),
			PostID:    int(readCmt.GetPostId()),
			AuthorID:  int(readCmt.GetAuthorId()),
			Content:   readCmt.GetContent(),
			UpVote:    int(readCmt.GetUpVote()),
			CreatedAt: int(readCmt.GetCreatedAt()),
			UpdatedAt: int(readCmt.GetUpdatedAt()),
		}
		Cmts = append(Cmts, cmt)
	}
	return Cmts
}

func writeCommentToGraphComment(writeComment *writePb.Comment) *model.Comment {
	if writeComment == nil {
		log.Println("CommentController.writeCommentToGraphComment(): writeComment is nil")
		return nil
	}
	cmt := &model.Comment{
		ID:        int(writeComment.GetId()),
		PostID:    int(writeComment.GetPostId()),
		AuthorID:  int(writeComment.GetAuthorId()),
		Content:   writeComment.GetContent(),
		UpVote:    int(writeComment.GetUpVote()),
		CreatedAt: int(writeComment.GetCreatedAt()),
		UpdatedAt: int(writeComment.GetUpdatedAt()),
	}
	return cmt
}
