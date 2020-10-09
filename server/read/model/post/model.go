package postModel

import (
	"database/sql"
	"fmt"
	"github.com/nvhai245/cyberblog/server/write/connection"
	"log"
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

// GetFeed func
func GetFeed(offset int32, limit int32) (success bool, foundPosts []*Post, err error) {
	queryString := "SELECT * FROM posts WHERE published = $1 ORDER BY published_at DESC OFFSET $2 LIMIT $3 "
	err = connection.DB.Select(&foundPosts, queryString, true, offset, limit)
	if err != nil {
		log.Println("Error in postModel.GetFeed(): ", err)
		return false, nil, err
	}
	return true, foundPosts, nil
}

// GetPostById func
func GetPostById(requesterId int32, postId int32) (bool, *Post, error) {
	foundPost := &Post{}
	queryString := "SELECT * FROM posts WHERE id = $1"
	rows, err := connection.DB.Queryx(queryString, postId)
	if err != nil {
		log.Println("Error in postModel.GetPostById(): ", err)
		return false, nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if rows.Err() == sql.ErrNoRows {
			log.Println("Error in model.GetPostById(): ", rows.Err())
			return false, nil, err
		}
		err = rows.StructScan(foundPost)
		if err != nil {
			log.Println("Error in postModel.GetPostById(): rows.StructScan()", err)
			return false, nil, err
		}
	}
	if foundPost == nil || (!foundPost.Published && foundPost.AuthorId != requesterId) {
		err = fmt.Errorf("this post is not published, only the creator can view its content")
		return false, nil, err
	}
	return true, foundPost, nil
}

// GetUserPublishedPosts func
func GetUserPublishedPosts(requesterId int32, userId int32, offset int32, limit int32) (success bool, foundPosts []*Post, err error) {
	queryString := "SELECT * FROM posts WHERE author_id = $1 AND published = $2 ORDER BY published_at DESC OFFSET $3 LIMIT $4 "
	err = connection.DB.Select(&foundPosts, queryString, userId, true, offset, limit)
	if err != nil {
		log.Println("Error in postModel.GetUserPublishedPosts(): ", err)
		return false, nil, err
	}
	return true, foundPosts, nil
}

// GetUserAllPosts func
func GetUserAllPosts(ownerId int32, offset int32, limit int32) (success bool, foundPosts []*Post, err error) {
	queryString := "SELECT * FROM posts WHERE author_id = $1 ORDER BY created_at DESC OFFSET $2 LIMIT $3"
	err = connection.DB.Select(&foundPosts, queryString, ownerId, offset, limit)
	if err != nil {
		log.Println("Error in postModel.GetUserAllPosts(): ", err)
		return false, nil, err
	}
	return true, foundPosts, nil
}

// GetUserUnPublishedPosts func
func GetUserUnPublishedPosts(ownerId int32, offset int32, limit int32) (success bool, foundPosts []*Post, err error) {
	queryString := "SELECT * FROM posts WHERE author_id = $1 AND published = $2 ORDER BY created_at DESC OFFSET $3 LIMIT $4"
	err = connection.DB.Select(&foundPosts, queryString, ownerId, false, offset, limit)
	if err != nil {
		log.Println("Error in postModel.GetUserUnPublishedPosts(): ", err)
		return false, nil, err
	}
	return true, foundPosts, nil
}

// GetCategoryPosts func
func GetCategoryPosts(categoryId int32, offset int32, limit int32) (success bool, foundPosts []*Post, err error) {
	queryString := "SELECT * FROM posts WHERE published = $1 AND id in (SELECT post_id FROM post_category WHERE category_id = $2) ORDER BY created_at DESC OFFSET $3 LIMIT $4"
	err = connection.DB.Select(&foundPosts, queryString, true, categoryId, offset, limit)
	if err != nil {
		log.Println("Error in postModel.GetUserUnPublishedPosts(): ", err)
		return false, nil, err
	}
	return true, foundPosts, nil
}
