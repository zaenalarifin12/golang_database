package respository

import (
	"context"
	"fmt"
	golang_database "golang-database"
	"golang-database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := newCommentRepositoryimpl(golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Text: "Hai aHiaiia",
	}

	result , err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := newCommentRepositoryimpl(golang_database.GetConnection())

	result , err := commentRepository.FindById(context.Background(), 25)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindAll(t *testing.T) {
	commentRepository := newCommentRepositoryimpl(golang_database.GetConnection())

	result , err := commentRepository.FindAll(context.Background())

	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
}