package categoryModel

import (
	"database/sql"
	"github.com/fatih/structs"
	"github.com/nvhai245/cyberblog/server/write/connection"
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

// Insert func
func Insert(newCategory *Category) (success bool, categoryId int32) {
	queryString := `INSERT INTO categories (title, slug, content) VALUES (:Title, :Slug, :Content) RETURNING id`
	rows, err := connection.DB.NamedQuery(queryString, structs.Map(newCategory))
	if err != nil {
		log.Println("Error in categoryModel.Insert(): ", err)
		return false, 0
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&categoryId)
		if err != nil {
			log.Println("Error in categoryModel.Insert(): rows.Scan()", err)
			return false, 0
		}
	}
	return true, categoryId
}

// Update func
func Update(newCategory *Category) (bool, *Category) {
	updatedCategory := &Category{}
	queryString := `UPDATE categories SET (title, slug, content) = (:Title, :Slug, :Content) WHERE id = :ID RETURNING *`
	rows, err := connection.DB.NamedQuery(queryString, structs.Map(newCategory))
	if err != nil {
		log.Println("Error in categoryModel.Update(): ", err)
		return false, nil
	}
	defer rows.Close()
	for rows.Next() {
		if rows.Err() == sql.ErrNoRows {
			log.Println("Error in categoryModel.Update(): no rows returned!", err)
			return false, nil
		}
		err = rows.StructScan(&updatedCategory)
		if err != nil {
			log.Println("Error in categoryModel.Update(): rows.Scan()", err)
			return false, nil
		}
	}
	if structs.IsZero(updatedCategory) {
		return false, nil
	}
	return true, updatedCategory
}

// Delete func
func Delete(categoryId int32) (bool, *Category) {
	deletedCategory := Category{}
	queryString := "DELETE FROM categories WHERE id = $1 RETURNING *"
	rows, err := connection.DB.Queryx(queryString, categoryId)
	if err != nil {
		log.Println("Error in categoryModel.Delete(): ", err)
		return false, &deletedCategory
	}
	defer rows.Close()
	for rows.Next() {
		if rows.Err() == sql.ErrNoRows {
			log.Println("Error in categoryModel.Delete(): no rows returned!", err)
			return false, &deletedCategory
		}
		err = rows.StructScan(&deletedCategory)
		if err != nil {
			log.Println("Error in categoryModel.Delete(): rows.Scan()", err)
			return false, &deletedCategory
		}
	}
	if structs.IsZero(deletedCategory) {
		return false, &deletedCategory
	}
	return true, &deletedCategory
}

// AddPostToCategory func
func AddPostToCategory(categoryId int32, postId int32) (bool, int32) {
	queryString := "INSERT INTO post_category (post_id, category_id) VALUES ($1, $2) RETURNING post_id"
	var insertedPostId int32
	rows, err := connection.DB.Query(queryString, postId, categoryId)
	if err != nil {
		log.Println("Error in categoryModel.AddPostToCategory(): ", err)
		return false, 0
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&insertedPostId)
		if err != nil {
			log.Println("Error in categoryModel.AddPostToCategory(): rows.Scan()", err)
			return false, 0
		}
	}
	return true, insertedPostId
}

// RemovePostFromCategory func
func RemovePostFromCategory(categoryId int32, postId int32) (bool, int32) {
	queryString := "DELETE FROM post_category WHERE category_id = $1 AND post_id = $2 RETURNING post_id"
	var removedPostId int32
	rows, err := connection.DB.Query(queryString, categoryId, postId)
	if err != nil {
		log.Println("Error in categoryModel.RemovePostFromCategory(): ", err)
		return false, 0
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&removedPostId)
		if err != nil {
			log.Println("Error in categoryModel.RemovePostFromCategory(): rows.Scan()", err)
			return false, 0
		}
	}
	return true, removedPostId
}
