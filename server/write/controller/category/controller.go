package categoryController

import (
	"fmt"
	categoryModel "github.com/nvhai245/cyberblog/server/write/model/category"
	pb "github.com/nvhai245/cyberblog/server/write/proto"
	"log"
)

func AddNewCategory(req *pb.AddNewCategoryRequest) (*pb.AddNewCategoryResponse, error) {
	log.Printf("commentController.AddNewCategory(), [Request]: %+v", req)
	success, ctgId := categoryModel.Insert(protoCategoryToModelCategory(req.GetNewCategory()))
	if success == false {
		err := fmt.Errorf("Error in categoryController.AddNewCategory(): failed to insert category")
		log.Println(err)
		return nil, err
	}
	res := &pb.AddNewCategoryResponse{
		Success:    true,
		CategoryId: ctgId,
	}
	// ******************************************************************************************
	log.Printf("commentController.AddNewCategory(), [Response]: %+v", res)
	return res, nil
}

func EditCategory(req *pb.EditCategoryRequest) (*pb.EditCategoryResponse, error) {
	log.Printf("commentController.EditCategory(), [Request]: %+v", req)
	success, editedCtg := categoryModel.Update(protoCategoryToModelCategory(req.GetNewCategory()))
	if success == false {
		err := fmt.Errorf("Error in categoryController.EditCategory(): failed to update category")
		log.Println(err)
		return nil, err
	}
	res := &pb.EditCategoryResponse{
		Success:        true,
		EditedCategory: modelCategoryToProtoCategory(editedCtg),
	}
	// ******************************************************************************************
	log.Printf("commentController.EditCategory(), [Response]: %+v", res)
	return res, nil
}

func DeleteCategory(req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error) {
	log.Printf("commentController.DeleteCategory(), [Request]: %+v", req)
	success, deletedCtg := categoryModel.Delete(req.GetCategoryId())
	if success == false {
		err := fmt.Errorf("Error in categoryController.DeleteCategory(): failed to update category")
		log.Println(err)
		return nil, err
	}
	res := &pb.DeleteCategoryResponse{
		Success:         true,
		DeletedCategory: modelCategoryToProtoCategory(deletedCtg),
	}
	// ******************************************************************************************
	log.Printf("commentController.DeleteCategory(), [Response]: %+v", res)
	return res, nil
}

func AddPostToCategory(req *pb.AddPostToCategoryRequest) (*pb.AddPostToCategoryResponse, error) {
	log.Printf("commentController.AddPostToCategory(), [Request]: %+v", req)
	success, postId := categoryModel.AddPostToCategory(req.GetCategoryId(), req.GetPostId())
	if success == false {
		err := fmt.Errorf("Error in categoryController.AddPostToCategory(): failed to add post to category")
		log.Println(err)
		return nil, err
	}
	res := &pb.AddPostToCategoryResponse{
		Success: true,
		PostId:  postId,
	}
	// ******************************************************************************************
	log.Printf("commentController.AddPostToCategory(), [Response]: %+v", res)
	return res, nil
}

func RemovePostFromCategory(req *pb.RemovePostFromCategoryRequest) (*pb.RemovePostFromCategoryResponse, error) {
	log.Printf("commentController.RemovePostFromCategory(), [Request]: %+v", req)
	success, postId := categoryModel.RemovePostFromCategory(req.GetCategoryId(), req.GetPostId())
	if success == false {
		err := fmt.Errorf("Error in categoryController.RemovePostFromCategory(): failed to add post to category")
		log.Println(err)
		return nil, err
	}
	res := &pb.RemovePostFromCategoryResponse{
		Success: true,
		PostId:  postId,
	}
	// ******************************************************************************************
	log.Printf("commentController.RemovePostFromCategory(), [Response]: %+v", res)
	return res, nil
}

func protoCategoryToModelCategory(protoCtg *pb.Category) *categoryModel.Category {
	modelCtg := &categoryModel.Category{
		ID:      protoCtg.GetId(),
		Title:   protoCtg.GetTitle(),
		Slug:    protoCtg.GetSlug(),
		Content: protoCtg.GetContent(),
	}
	return modelCtg
}

func modelCategoryToProtoCategory(modelCtg *categoryModel.Category) *pb.Category {
	if modelCtg == nil {
		log.Println("Error in categoryController.modelCategoryToProtoCategory(): modelCtg is nil")
	}
	protoCtg := &pb.Category{
		Id:      modelCtg.ID,
		Title:   modelCtg.Title,
		Slug:    modelCtg.Slug,
		Content: modelCtg.Content,
	}
	return protoCtg
}
