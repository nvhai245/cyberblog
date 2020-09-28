package commentModel

import (
	"database/sql"
	"github.com/fatih/structs"
	"github.com/nvhai245/cyberblog/server/write/connection"
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

func Insert(newComment *Comment) (success bool, commentId int32) {
	queryString := `INSERT INTO comments (author_id, post_id, up_vote, content) VALUES (:AuthorId, :PostId, :UpVote, :Content) RETURNING id`
	rows, err := connection.DB.NamedQuery(queryString, structs.Map(newComment))
	if err != nil {
		log.Println("Error in commentModel.Insert(): ", err)
		return false, 0
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&commentId)
		if err != nil {
			log.Println("Error in commentModel.Insert(): rows.Scan()", err)
			return false, 0
		}
	}
	return true, commentId
}

func Update(newComment *Comment) (bool, *Comment) {
	updatedComment := &Comment{}
	queryString := `UPDATE comments SET (content, updated_at) = (:Content, :UpdatedAt) WHERE id = :ID AND author_id = :AuthorId RETURNING *`
	rows, err := connection.DB.NamedQuery(queryString, structs.Map(newComment))
	if err != nil {
		log.Println("Error in commentModel.Update(): ", err)
		return false, nil
	}
	defer rows.Close()
	for rows.Next() {
		if rows.Err() == sql.ErrNoRows {
			log.Println("Error in commentModel.Update(): no rows returned!", err)
			return false, nil
		}
		err = rows.StructScan(&updatedComment)
		if err != nil {
			log.Println("Error in commentModel.Update(): rows.Scan()", err)
			return false, nil
		}
	}
	if structs.IsZero(updatedComment) {
		return false, nil
	}
	return true, updatedComment
}

func Delete(authorId int32, commentId int32) (bool, *Comment) {
	deletedComment := Comment{}
	queryString := "DELETE FROM comments WHERE id = $1 AND author_id = $2 RETURNING *"
	rows, err := connection.DB.Queryx(queryString, commentId, authorId)
	if err != nil {
		log.Println("Error in commentModel.Delete(): ", err)
		return false, &deletedComment
	}
	defer rows.Close()
	for rows.Next() {
		if rows.Err() == sql.ErrNoRows {
			log.Println("Error in commentModel.Delete(): no rows returned!", err)
			return false, &deletedComment
		}
		err = rows.StructScan(&deletedComment)
		if err != nil {
			log.Println("Error in commentModel.Delete(): rows.Scan()", err)
			return false, &deletedComment
		}
	}
	if structs.IsZero(deletedComment) {
		return false, &deletedComment
	}
	return true, &deletedComment
}

// UpVote func
func UpVote(commentId int32) (success bool, newUpVotes int32) {
	queryString := "UPDATE comments SET up_vote = up_vote + 1 WHERE id = $1 RETURNING up_vote"
	rows, err := connection.DB.Queryx(queryString, commentId)
	if err != nil {
		log.Println("Error in commentModel.UpVote(): ", err)
		return false, 0
	}
	defer rows.Close()
	for rows.Next() {
		if rows.Err() == sql.ErrNoRows {
			log.Println("Error in commentModel.UpVote(): no rows returned!", err)
			return false, 0
		}
		err = rows.Scan(&newUpVotes)
		if err != nil {
			log.Println("Error in commentModel.UpVote(): rows.Scan()", err)
			return false, 0
		}
	}
	return true, newUpVotes
}

// DownVote func
func DownVote(commentId int32) (success bool, newUpVotes int32) {
	queryString := "UPDATE comments SET up_vote = up_vote - 1 WHERE id = $1 RETURNING up_vote"
	rows, err := connection.DB.Queryx(queryString, commentId)
	if err != nil {
		log.Println("Error in commentModel.DownVote(): ", err)
		return false, 0
	}
	defer rows.Close()
	for rows.Next() {
		if rows.Err() == sql.ErrNoRows {
			log.Println("Error in commentModel.DownVote(): no rows returned!", err)
			return false, 0
		}
		err = rows.Scan(&newUpVotes)
		if err != nil {
			log.Println("Error in commentModel.DownVote(): rows.Scan()", err)
			return false, 0
		}
	}
	return true, newUpVotes
}
