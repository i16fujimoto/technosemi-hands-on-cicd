package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

// 仮のテーブルを作成
type user struct {
	id   int
	name string
}

type post struct {
	id        int
	userID    int
	content   string
	createdAt int
}

var (
	// 仮のデータを作成
	// Users ユーザーテーブル
	Users = map[int]*user{
		1: {id: 1, name: "user1"},
		2: {id: 2, name: "user2"},
	}
	// Posts 投稿テーブル
	Posts = map[int]*post{
		1: {id: 1, userID: 1, content: "post1", createdAt: 1},
		2: {id: 2, userID: 2, content: "post2", createdAt: 2},
		3: {id: 3, userID: 1, content: "post3", createdAt: 3},
		4: {id: 4, userID: 1, content: "post4", createdAt: 4},
	}

	// エラー定義
	ErrUserNotFound = errors.New("user not found")
)

func main() {
	// ユーザ1の投稿一覧を取得
	user, err :=    getUser(1)
	if err != nil {
		log.Println(err.Error())
	}
	posts := listPostsByUserIDs([]int{user.id})

	// 取得した投稿一覧を出力
	for _, post := range posts {
		log.Println(post)
	}
}

// 特定のユーザを取得
func getUser(id int) (*user, error) {
	if _, ok := Users[id]; !ok {
		return nil, fmt.Errorf("%w", ErrUserNotFound)
	}
	return Users[id], nil
}

// 特定のユーザの投稿一覧を取得
func listPostsByUserIDs(userIDs []int) []*post {
	if len(userIDs) == 0 {
		return []*post{}
	}
	userIDMap := make(map[int]struct{}, len(userIDs))
	for _, id := range userIDs {
		userIDMap[id] = struct{}{}
	}

	var posts []*post
	for _, post := range Posts {
		if _, ok := userIDMap[post.userID]; ok {
			posts = append(posts, post)
		}
	}

	// 投稿日時でソート
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].createdAt < posts[j].createdAt
	})

	return posts
}
