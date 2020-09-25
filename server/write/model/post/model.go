package postModel

import (
	"database/sql"
	"github.com/fatih/structs"
	"github.com/nvhai245/cyberblog/server/write/connection"
	"log"
	"time"
)

type Post struct {
	ID          int32  `db:"id"`
	AuthorId    int32  `db:"author_id"`
	ParentId    int32  `db:"parent_id"`
	Title       string `db:"title"`
	Published   bool   `db:"published"`
	UpVote      int32  `db:"up_vote"`
	Content     string `db:"content"`
	CreatedAt   int32  `db:"created_at"`
	UpdatedAt   int32  `db:"updated_at"`
	PublishedAt int32  `db:"published_at"`
}

type Category struct {
	ID      int32  `db:"id"`
	Title   string `db:"title"`
	Slug    string `db:"slug"`
	Content string `db:"content"`
}

type PostCategory struct {
	PostId     int32 `db:"post_id"`
	CategoryId int32 `db:"category_id"`
}

// Insert func
func Insert(newPost *Post) (success bool, postId int32) {
	queryString := `INSERT INTO posts (author_id, parent_id, title, published, up_vote, content) VALUES (:AuthorId, :ParentId, :Title, :Published, :UpVote, :Content) RETURNING id`
	rows, err := connection.DB.NamedQuery(queryString, structs.Map(newPost))
	if err != nil {
		log.Println("Error in postModel.Insert(): ", err)
		return false, 0
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&postId)
		if err != nil {
			log.Println("Error in postModel.Insert(): rows.Scan()", err)
			return false, 0
		}
	}
	return true, postId
}

// Update func
func Update(newPost *Post) (bool, *Post) {
	updatedPost := &Post{}
	queryString := `UPDATE posts SET (title, content, updated_at) = (:Title, :Content, :UpdatedAt) WHERE id = :ID RETURNING *`
	rows, err := connection.DB.NamedQuery(queryString, structs.Map(newPost))
	if err != nil {
		log.Println("Error in postModel.Update(): ", err)
		return false, nil
	}
	defer rows.Close()
	for rows.Next() {
		if rows.Err() == sql.ErrNoRows {
			log.Println("Error in postModel.Update(): no rows returned!", err)
			return false, nil
		}
		err = rows.StructScan(&updatedPost)
		if err != nil {
			log.Println("Error in postModel.Update(): rows.Scan()", err)
			return false, nil
		}
	}
	if structs.IsZero(updatedPost) {
		return false, nil
	}
	return true, updatedPost
}

// Delete func
func Delete(requestorId int32, postId int32) (bool, *Post) {
	deletedPost := Post{}
	queryString := "DELETE FROM posts WHERE id = $1 AND author_id = $2 RETURNING *"
	rows, err := connection.DB.Queryx(queryString, postId, requestorId)
	if err != nil {
		log.Println("Error in postModel.Delete(): ", err)
		return false, &deletedPost
	}
	defer rows.Close()
	for rows.Next() {
		if rows.Err() == sql.ErrNoRows {
			log.Println("Error in postModel.Delete(): no rows returned!", err)
			return false, &deletedPost
		}
		err = rows.StructScan(&deletedPost)
		if err != nil {
			log.Println("Error in postModel.Delete(): rows.Scan()", err)
			return false, &deletedPost
		}
	}
	if structs.IsZero(deletedPost) {
		return false, &deletedPost
	}
	return false, nil
}

// Publish func
func Publish(requestorId int32, postId int32) (success bool, publishedPost *Post) {
	queryString := "UPDATE posts SET published = $1 AND published_at = $2 WHERE id = $3 AND author_id = $4 RETURNING *"
	rows, err := connection.DB.Queryx(queryString, true, time.Now().Unix(), postId, requestorId)
	if err != nil {
		log.Println("Error in postModel.Publish(): ", err)
		return false, nil
	}
	defer rows.Close()
	for rows.Next() {
		if rows.Err() == sql.ErrNoRows {
			log.Println("Error in postModel.Publish(): no rows returned!", err)
			return false, nil
		}
		err = rows.StructScan(&publishedPost)
		if err != nil {
			log.Println("Error in postModel.Publish(): rows.Scan()", err)
			return false, nil
		}
	}
	if structs.IsZero(publishedPost) {
		return false, nil
	}
	return true, publishedPost
}

// UnPublish func
func UnPublish(requestorId int32, postId int32) (success bool, unPublishedPost *Post) {
	queryString := "UPDATE posts SET published = $1 AND published_at = $2 WHERE id = $3 AND author_id = $4 RETURNING *"
	rows, err := connection.DB.Queryx(queryString, false, nil, postId, requestorId)
	if err != nil {
		log.Println("Error in postModel.UnPublish(): ", err)
		return false, nil
	}
	defer rows.Close()
	for rows.Next() {
		if rows.Err() == sql.ErrNoRows {
			log.Println("Error in postModel.UnPublish(): no rows returned!", err)
			return false, nil
		}
		err = rows.StructScan(&unPublishedPost)
		if err != nil {
			log.Println("Error in postModel.UnPublish(): rows.Scan()", err)
			return false, nil
		}
	}
	if structs.IsZero(unPublishedPost) {
		return false, nil
	}
	return true, unPublishedPost
}

// UpVote func
func UpVote(upVoterId int32, postId int32) (success bool, newUpVotes int32) {
	queryString := "UPDATE posts SET up_vote = up_vote + 1 WHERE id = $1 RETURNING up_vote"
	rows, err := connection.DB.Queryx(queryString, postId)
	if err != nil {
		log.Println("Error in postModel.UpVote(): ", err)
		return false, 0
	}
	defer rows.Close()
	for rows.Next() {
		if rows.Err() == sql.ErrNoRows {
			log.Println("Error in postModel.UpVote(): no rows returned!", err)
			return false, 0
		}
		err = rows.Scan(&newUpVotes)
		if err != nil {
			log.Println("Error in postModel.UpVote(): rows.Scan()", err)
			return false, 0
		}
	}
	if structs.IsZero(newUpVotes) {
		return false, 0
	}
	return true, newUpVotes
}
