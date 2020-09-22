package categoryController

import (
	"fmt"
	categoryModel "github.com/nvhai245/cyberblog/server/read/model/category"
	pb "github.com/nvhai245/cyberblog/server/read/proto"
	"log"
)

func GetAllCategories(req *pb.GetAllCategoriesRequest) (*pb.GetAllCategoriesResponse, error) {
	log.Printf("categoryController.GetAllCategories(), [Request]: %+v", req)
	success, allCtgs := categoryModel.GetAllCategories(req.GetRequesterId())
	if success == false {
		err := fmt.Errorf("Error in categoryController.GetAllCategories(): failed to find category")
		log.Println(err)
		return nil, err
	}
	res := &pb.GetAllCategoriesResponse{
		Success:       true,
		AllCategories: modelCategoriesToProtoCategories(allCtgs),
	}
	// ******************************************************************************************
	log.Printf("categoryController.GetAllCategories(), [Response]: %+v", res)
	return res, nil
}

func GetPostCategories(req *pb.GetPostCategoriesRequest) (*pb.GetPostCategoriesResponse, error) {
	log.Printf("categoryController.GetPostCategories(), [Request]: %+v", req)
	success, allCtgs := categoryModel.GetPostCategories(req.GetPostId())
	if success == false {
		err := fmt.Errorf("Error in categoryController.GetPostCategories(): failed to find category")
		log.Println(err)
		return nil, err
	}
	res := &pb.GetPostCategoriesResponse{
		Success:        true,
		PostCategories: modelCategoriesToProtoCategories(allCtgs),
	}
	// ******************************************************************************************
	log.Printf("categoryController.GetPostCategories(), [Response]: %+v", res)
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

func modelCategoriesToProtoCategories(modelCtgs []*categoryModel.Category) []*pb.Category {
	if len(modelCtgs) == 0 {
		log.Println("Error in categoryController.modelCategoryToProtoCategory(): modelCtgs is empty slice")
	}
	var protoCtgs []*pb.Category
	for _, modelCtg := range modelCtgs {
		protoCtg := &pb.Category{
			Id:      modelCtg.ID,
			Title:   modelCtg.Title,
			Slug:    modelCtg.Slug,
			Content: modelCtg.Content,
		}
		protoCtgs = append(protoCtgs, protoCtg)
	}
	return protoCtgs
}
