package api

import (
	"Sclad/db"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"time"
)

func (h *Handlers) AddItem(rw http.ResponseWriter, r *http.Request) {
	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte(`{ "message": "` + "can't read body" + `"}`))
		return
	}
	userId, _ := (*r).Header["Storage-User-Id"]
	category, _ := (*r).Header["Storage-Category"]
	title, _ := (*r).Header["Storage-Title"]
	count, _ := (*r).Header["Storage-Count"]
	text, _ := (*r).Header["Storage-Text"]

	var body CreateTextRequest

	json.Unmarshal(rawBody, &body)

	postId, _ := h.generatePostId()

	var post = db.Post{
		PostId:       postId,
		ICategory:    category[0],
		ITitle:       title[0],
		ICount:       count[0],
		Text:         text[0],
		UserId:       userId[0],
		ISOTimestamp: time.Now().Format(time.RFC3339),
	}

	h.Database.CreatePost(post)

	response, _ := json.Marshal(post)
	writeJsonToResponse(&rw, http.StatusOK, response)
}

func (h *Handlers) GetItem(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	params := mux.Vars(r)
	id, _ := params["itemId"]

	post, _ := h.Database.GetPost(id)

	response, _ := json.Marshal(post)
	writeJsonToResponse(&rw, http.StatusOK, response)
}

func (h *Handlers) GetItems(rw http.ResponseWriter, r *http.Request) {
	
    w.Header().Set("Access-Control-Allow-Origin", "*")

    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")



	params := mux.Vars(r)
	userId := params["userId"]

	//получаем все данные по userId
	posts, _ := h.Database.GetUserPosts(userId)

	response, _ := json.Marshal(posts)
	writeJsonToResponse(&rw, http.StatusOK, response)
}
