package respository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
	"strconv"
)

type commentRepositoryimpl struct {
	DB *sql.DB
}

func newCommentRepositoryimpl(DB *sql.DB) CommentRepository {
	return &commentRepositoryimpl{DB: DB}
}

func (repository *commentRepositoryimpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(text) VALUES (?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Text)
	if err != nil {
		return comment, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(insertId)

	return comment, nil
}

func (repository *commentRepositoryimpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, text FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next(){
	//	ada
		rows.Scan(&comment.Id, &comment.Text)
		return comment, nil
	}else{
	// tidak ada
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository *commentRepositoryimpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, text FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Text)
		comments = append(comments, comment)
	}
	return comments, nil

}

