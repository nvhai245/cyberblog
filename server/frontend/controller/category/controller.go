package categoryController

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
)

func GetAllCategories(ctx context.Context, requesterID int) (*model.GetCategoriesResponse, error) {
	session := helper.GetSession(ctx, "auth")
	token := fmt.Sprintf("%v", session.Values["token"])
	checkResponse, err := connection.AuthClient.CheckToken(context.Background(), &authPb.CheckTokenRequest{Token: token})
	if err != nil {
		log.Println("Error in rpc AuthClient.CheckToken(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	// No need to check auth
	_ = checkResponse
	result, err := connection.ReadClient.GetAllCategories(ctx, &readPb.GetAllCategoriesRequest{RequesterId: int32(requesterID)})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	var ctgs []*model.Category
	for _, ctg := range result.GetAllCategories() {
		ctgs = append(ctgs, readCtgToGraphCtg(ctg))
	}

	return &model.GetCategoriesResponse{
		Message:    "get all categories successful!",
		Categories: ctgs,
	}, nil
}

func GetPostCategories(ctx context.Context, postId int) (*model.GetCategoriesResponse, error) {
	session := helper.GetSession(ctx, "auth")
	token := fmt.Sprintf("%v", session.Values["token"])
	checkResponse, err := connection.AuthClient.CheckToken(context.Background(), &authPb.CheckTokenRequest{Token: token})
	if err != nil {
		log.Println("Error in rpc AuthClient.CheckToken(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	// No need to check auth
	_ = checkResponse
	result, err := connection.ReadClient.GetPostCategories(ctx, &readPb.GetPostCategoriesRequest{PostId: int32(postId)})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	var ctgs []*model.Category
	for _, ctg := range result.GetPostCategories() {
		ctgs = append(ctgs, readCtgToGraphCtg(ctg))
	}

	return &model.GetCategoriesResponse{
		Message:    "get post categories successful!",
		Categories: ctgs,
	}, nil
}

func AddNewCategory(ctx context.Context, newCategory model.NewCategory) (*model.GetCategoryResponse, error) {
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
	if !checkResponse.GetIsAdmin() {
		return nil, fmt.Errorf("YOU ARE NOT AN ADMIN!")
	}

	result, err := connection.WriteClient.AddNewCategory(ctx, &writePb.AddNewCategoryRequest{NewCategory: &writePb.Category{
		Title:   newCategory.Title,
		Slug:    newCategory.Slug,
		Content: newCategory.Content,
	}})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	return &model.GetCategoryResponse{
		Message: "Add new Category successful!",
		Category: &model.Category{
			ID: int(result.GetCategoryId()),
		},
	}, nil
}

func EditCategory(ctx context.Context, newCategory model.NewCategory) (*model.GetCategoryResponse, error) {
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
	if !checkResponse.GetIsAdmin() {
		return nil, fmt.Errorf("YOU ARE NOT AN ADMIN!")
	}
	result, err := connection.WriteClient.EditCategory(ctx, &writePb.EditCategoryRequest{NewCategory: &writePb.Category{
		Id:      int32(newCategory.ID),
		Title:   newCategory.Title,
		Slug:    newCategory.Slug,
		Content: newCategory.Content,
	}})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}

	return &model.GetCategoryResponse{
		Message:  "Edit Category successful!",
		Category: writeCtgToGraphCtg(result.GetEditedCategory()),
	}, nil
}

func DeleteCategory(ctx context.Context, categoryID int) (*model.GetCategoryResponse, error) {
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
	if !checkResponse.GetIsAdmin() {
		return nil, fmt.Errorf("YOU ARE NOT AN ADMIN!")
	}
	result, err := connection.WriteClient.DeleteCategory(ctx, &writePb.DeleteCategoryRequest{CategoryId: int32(categoryID)})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}

	return &model.GetCategoryResponse{
		Message:  "Delete category successful!",
		Category: writeCtgToGraphCtg(result.GetDeletedCategory()),
	}, nil
}

func AddPostToCategory(ctx context.Context, categoryID int, postID int) (*model.PostCategoryResponse, error) {
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
	result, err := connection.WriteClient.AddPostToCategory(ctx, &writePb.AddPostToCategoryRequest{
		CategoryId: int32(categoryID),
		PostId:     int32(postID),
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}

	return &model.PostCategoryResponse{
		Message: "add post to category successful!",
		PostID:  int(result.GetPostId()),
	}, nil
}

func RemovePostFromCategory(ctx context.Context, categoryID int, postID int) (*model.PostCategoryResponse, error) {
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
	result, err := connection.WriteClient.RemovePostFromCategory(ctx, &writePb.RemovePostFromCategoryRequest{
		CategoryId: int32(categoryID),
		PostId:     int32(postID),
	})
	if err != nil {
		return nil, err
	}
	if result.GetSuccess() == false {
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}

	return &model.PostCategoryResponse{
		Message: "remove post from category successful!",
		PostID:  int(result.GetPostId()),
	}, nil
}

func writeCtgToGraphCtg(writeCtg *writePb.Category) *model.Category {
	if writeCtg == nil {
		log.Println("CategoryController.writeCtgToGraphCtg(): writeCtg is nil")
		return nil
	}
	graphCtg := &model.Category{
		ID:      int(writeCtg.GetId()),
		Title:   writeCtg.GetTitle(),
		Slug:    writeCtg.GetSlug(),
		Content: writeCtg.GetContent(),
	}
	return graphCtg
}

func graphCtgToReadCtg(graphCtg *model.Category) *readPb.Category {
	if graphCtg == nil {
		log.Println("CategoryController.graphCtgToWriteCtg(): graphCtg is nil")
		return nil
	}
	readCtg := &readPb.Category{
		Id:      int32(graphCtg.ID),
		Title:   graphCtg.Title,
		Slug:    graphCtg.Slug,
		Content: graphCtg.Content,
	}
	return readCtg
}

func readCtgToGraphCtg(readCtg *readPb.Category) *model.Category {
	if readCtg == nil {
		log.Println("CategoryController.writeCtgToGraphCtg(): writeCtg is nil")
		return nil
	}
	graphCtg := &model.Category{
		ID:      int(readCtg.GetId()),
		Title:   readCtg.GetTitle(),
		Slug:    readCtg.GetSlug(),
		Content: readCtg.GetContent(),
	}
	return graphCtg
}
