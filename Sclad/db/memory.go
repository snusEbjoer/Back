package db

import (
	"errors"
	"sync"
)

type Memory struct {
	mu      sync.RWMutex
	Posts   map[string]Post
	Authors map[string][]string
}

func (m *Memory) InitDB() {

}

func (m *Memory) IsConnected() error {
	return nil
}

func (m *Memory) CreatePost(post Post) error {
	m.mu.Lock()
	m.Posts[post.PostId] = post
	m.Authors[post.UserId] = append(m.Authors[post.UserId], post.PostId)
	m.mu.Unlock()
	return nil
}

func (m *Memory) GetPost(id string) (Post, error) {
	m.mu.Lock()
	post, err := m.Posts[id]
	m.mu.Unlock()
	if !err {
		err := errors.New("cannot find post_id")
		return post, err
	}
	return post, nil
}

func (m *Memory) GetUserPosts(userId string) ([]Post, error) {
	var posts []Post
	m.mu.Lock()
	for _, post := range m.Posts {
		if post.UserId == userId {
			posts = append(posts, post)
		}
	}

	m.mu.Unlock()
	return posts, nil
}

type GetPostsResponse struct {
	Posts    []Post `json:"posts"`
	NextPage string `json:"nextPage,omitempty"`
}

func (m *Memory) Check(id string) bool {
	idAlreadyUsed := false
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.Posts[id]; ok {
		idAlreadyUsed = true
	}

	return idAlreadyUsed
}
