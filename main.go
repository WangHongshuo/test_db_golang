package main

import (
	"fmt"

	"test_db/dao"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
Run commands in Postgre SQL Shell:
1. create database acfun;
2. \l (confirm database acfun exist)
3. \c acfun;
4. create table comments(
	comment_id 	bigint 	primary key check (comment_id > 0),
	floor_id 	int 				check (floor_id > 0),
	article_id 	bigint 				check (article_id > 0),
	user_id 	bigint 				check (user_id > 0),
	user_name 	text,
	content 	text
);
5. \d comments_test; (confirm table comments_test exist)
6. create index article_id on comments(article_id);
7. create index user_id on comments(user_id);
8. \d comments_test; (confirm index create success)
*/

func main() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=1 port=5432 dbname=acfun sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	comment1 := &dao.Comment{
		CommentId: 1,
		FloorId:   1,
		ArticleId: 1,
		UserId:    1,
		UserName:  "TestUser1",
		Content:   "TestUser1 Content on Artice 1, Floor 1",
	}

	comment2 := &dao.Comment{
		CommentId: 3,
		FloorId:   2,
		ArticleId: 1,
		UserId:    2,
		UserName:  "TestUser2",
		Content:   "TestUser2 Content on Article 1, Floor 2",
	}

	comment3 := &dao.Comment{
		CommentId: 4,
		FloorId:   1,
		ArticleId: 2,
		UserId:    1,
		UserName:  "TestUser1",
		Content:   "TestUser1 Content on Article 2, Floor 1",
	}

	comment4 := &dao.Comment{
		CommentId: 5,
		FloorId:   3,
		ArticleId: 2,
		UserId:    2,
		UserName:  "TestUser2",
		Content:   "TestUser2 Content on Article 2, Floor 3",
	}

	db.Create(comment1)
	db.Create(comment2)
	db.Create(comment3)
	db.Create(comment4)

	var result []dao.Comment
	db.Where("user_id = ?", "1").Find(&result)
	fmt.Printf("Result: %+v\n", result)

	db.Where("user_id = ?", "2").Find(&result)
	fmt.Printf("Result: %+v\n", result)

	db.Where("article_id = ?", "1").Find(&result)
	fmt.Printf("Result: %+v\n", result)

	db.Where("article_id = ?", "2").Find(&result)
	fmt.Printf("Result: %+v\n", result)

	// delete all
	db.Exec("delete from comments;")
}
