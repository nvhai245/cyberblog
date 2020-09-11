package categoryModel

import (
	"github.com/nvhai245/cyberblog/server/read/connection"
	"log"
)

type Category struct {
	ID        int32  `db:"id"`
	Title     string `db:"title"`
	Slug      string `db:"slug"`
	Content   string `db:"content"`
	CreatedAt int32  `db:"created_at"`
	UpdatedAt int32  `db:"updated_at"`
}

type PostCategory struct {
	PostId     int32 `db:"post_id"`
	CategoryId int32 `db:"category_id"`
}

// GetAllCategories func
func GetAllCategories(requestorId int32) (success bool, allCategories []*Category) {
	queryString := "SELECT * FROM categories"
	err := connection.DB.Select(&allCategories, queryString)
	if err != nil {
		log.Println("Error in categoryModel.GetAllCategories(): ", err)
		return false, nil
	}
	return true, allCategories
}

// GetPostCategories func
func GetPostCategories(postId int32) (success bool, foundCategory []*Category) {
	queryString := "SELECT * FROM categories WHERE id in (SELECT category_id FROM post_category WHERE post_id = $1)"
	err := connection.DB.Select(&foundCategory, queryString, postId)
	if err != nil {
		log.Println("Error in categoryModel.GetPostCategory(): ", err)
		return false, nil
	}
	return true, foundCategory
}
