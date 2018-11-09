package storage

import (
	"bla/models"
	"database/sql"
	"github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"
)

const sqlGetAllPosts string = `SELECT id, tag, title, published FROM posts`
const sqlGetAllFavoritePosts string = `SELECT id, tag, title, published FROM posts WHERE is_favorite = TRUE`
const sqlGetPostById string = `SELECT id, tag, title, content_md, content_html, published, edited, is_favorite, author FROM posts WHERE id = ?`
const sqlGetPostByTag string = `SELECT id, tag, title, content_md, content_html, published, edited, is_favorite, author FROM posts WHERE tag = ?`

const sqlCreatePost string = `INSERT INTO posts (id, tag, title, content_md, content_html, published, edited, is_favorite, author) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
const sqlUpdatePost string = `UPDATE posts SET tag = ?, title = ?, content_md = ?, content_html = ?, published = ?, edited = ?, is_favorite = ?, author = ? WHERE id = ?`

func GetAllPosts(db *sql.DB) (*[]models.PostLite, error) {
	//db, err := sql.Open("sqlite3", "./content/data/bla.db")
	rows, err := db.Query(sqlGetAllPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.PostLite

	for rows.Next() {
		post := models.PostLite{}
		err = rows.Scan(&post.Id, &post.Tag, &post.Title, &post.Published)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return &posts, nil
}

func GetAllFavoritePosts(db *sql.DB) (*[]models.PostLite, error) {
	rows, err := db.Query(sqlGetAllFavoritePosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.PostLite

	for rows.Next() {
		post := models.PostLite{}
		err = rows.Scan(&post.Id, &post.Tag, &post.Title, &post.Published)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return &posts, nil
}

func GetPostById(db *sql.DB, id uuid.UUID) (*models.Post, error) {
	row := db.QueryRow(sqlGetPostById, id)

	post := models.Post{}

	err := row.Scan(&post.Id, &post.Tag, &post.Title, &post.ContentMD, &post.ContentHTML, &post.Published, &post.Edited, &post.IsFavorite, &post.Author)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func GetPostByTag(db *sql.DB, tag string) (*models.Post, error) {
	row := db.QueryRow(sqlGetPostByTag, tag)

	post := models.Post{}

	err := row.Scan(&post.Id, &post.Tag, &post.Title, &post.ContentMD, &post.ContentHTML, &post.Published, &post.Edited, &post.IsFavorite, &post.Author)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func CreatePost(db *sql.DB, post *models.Post) (*models.Post, error) {
	writeDB, err := db.Begin()
	if err != nil {
		writeDB.Rollback()
		return nil, err
	}

	id, err := uuid.NewV4()
	if err != nil {
		writeDB.Rollback()
		return nil, err
	}

	_, err = writeDB.Exec(
		sqlCreatePost,
		id,
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
		return nil, err
	}

	if err = writeDB.Commit(); err != nil {
		return nil, err
	}

	return GetPostById(db, id)
}

func UpdatePost(db *sql.DB, post *models.Post, id uuid.UUID) (*models.Post, error) {
	writeDB, err := db.Begin()
	if err != nil {
		writeDB.Rollback()
		return nil, err
	}

	_, err = writeDB.Exec(
		sqlUpdatePost,
		post.Tag,
		post.Title,
		post.ContentMD,
		post.ContentHTML,
		post.Published,
		post.Edited,
		post.IsFavorite,
		post.Author,
		id)
	if err != nil {
		writeDB.Rollback()
		return nil, err
	}

	if err = writeDB.Commit(); err != nil {
		return nil, err
	}

	return GetPostById(db, id)
}
