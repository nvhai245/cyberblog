package commentModel

import (
	"github.com/nvhai245/cyberblog/server/read/connection"
	"log"
)

type Comment struct {
	ID        int32  `db:"id"`
	PostId    int32  `db:"post_id"`
	AuthorId  int32  `db:"author_id"`
	Content   string `db:"content"`
	UpVote    int32  `db:"up_vote"`
	CreatedAt int32  `db:"created_at"`
	UpdatedAt int32  `db:"updated_at"`
}

// GetPostComments func
func GetPostComments(postId int32, offset int32, limit int32) (success bool, foundComments []*Comment) {
	queryString := `SELECT * FROM comments WHERE post_id = $1 ORDER BY created_at OFFSET $2 LIMIT $3`
	err := connection.DB.Select(&foundComments, queryString, postId, offset, limit)
	if err != nil {
		log.Println("Error in commentModel.GetPostComments(): ", err)
		return false, nil
	}
	return true, foundComments
}

// GetUserComments func
func GetUserComments(authorId int32, offset int32, limit int32) (success bool, foundComments []*Comment) {
	queryString := `SELECT * FROM comments WHERE author_id = $1 ORDER BY created_at DESC OFFSET $2 LIMIT $3`
	err := connection.DB.Select(&foundComments, queryString, authorId, offset, limit)
	if err != nil {
		log.Println("Error in commentModel.GetUserComments(): ", err)
		return false, nil
	}
	return true, foundComments
}
