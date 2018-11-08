package storage

import "bla/models"

const sqlGetAllPosts string = `SELECT tag, title, published FROM posts`
const sqlGetAllFavoritePosts string = `SELECT tag, title, published FROM posts WHERE is_favorite = TRUE`
const sqlGetPostById string = `SELECT id, tag, title, content_md, content_html, published, edited, is_favorite, author FROM posts WHERE id = ?`
const sqlGetPostByTag string = `SELECT id, tag, title, content_md, content_html, published, edited, is_favorite, author FROM posts WHERE tag = ?`

const sqlCreatePost string = `INSERT INTO posts (id, tag, title, content_md, content_html, published, edited, is_favorite, author) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

func GetAllPosts() (*[]models.PostLite, error) {

	return nil, nil
}

func GetAllFavoritePosts() (*[]models.PostLite, error) {

	return nil, nil
}

func GetPostById(id string) (*models.Post, error) {

	return nil, nil
}

func GetPostByTag(tag string) (*models.Post, error) {

	return nil, nil
}

func CreatePost(post *models.Post) error {
	writeDB, err := readDB.Begin()
	if err != nil {
		writeDB.Rollback()
		return err
	}

	_, err = writeDB.Exec(
		sqlCreatePost,
		post.Id,
		post.Tag,
		post.Title,
		post.ContentMD,
		post.ContentHTML,
		post.Published,
		post.Edited,
		post.IsFavorite,
		post.Author)
	if err != nil {
		writeDB.Rollback()
		return err
	}

	return writeDB.Commit()
}
