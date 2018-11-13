package storage

import (
	"bla/models"
	"database/sql"
	"github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

const sqlGetAllPosts string = `SELECT id, tag, title, published, is_favorite FROM posts`
const sqlGetAllFavoritePosts string = `SELECT id, tag, title, published, is_favorite FROM posts WHERE is_favorite = TRUE`
const sqlGetPostById string = `SELECT id, tag, title, content_md, content_html, published, edited, is_favorite, author FROM posts WHERE id = ?`
const sqlGetPostByTag string = `SELECT id, tag, title, content_md, content_html, published, edited, is_favorite, author FROM posts WHERE tag = ?`

const sqlCreatePost string = `INSERT INTO posts (id, tag, title, content_md, content_html, published, edited, is_favorite, author) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
const sqlUpdatePost string = `UPDATE posts SET tag = ?, title = ?, content_md = ?, content_html = ?, published = ?, edited = ?, is_favorite = ?, author = ? WHERE id = ?`

// GetAllPosts will contact the database that is passed into it and query for all posts.
// It will then return a slice of PostLite objects, containing only minimal identifying information
// about the posts, meant to be used for displaying many post titles in a list.
func GetAllPosts(db *sql.DB) (*[]models.PostLite, error) {
	rows, err := db.Query(sqlGetAllPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.PostLite

	for rows.Next() {
		post := models.PostLite{}
		err = rows.Scan(&post.Id, &post.Tag, &post.Title, &post.Published, &post.IsFavorite)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return &posts, nil
}

// GetAllFavoritePosts does the same thing as GetAllPosts, but only returns the posts marked as
// IsFavorite in the database.
func GetAllFavoritePosts(db *sql.DB) (*[]models.PostLite, error) {
	rows, err := db.Query(sqlGetAllFavoritePosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.PostLite

	for rows.Next() {
		post := models.PostLite{}
		err = rows.Scan(&post.Id, &post.Tag, &post.Title, &post.Published, &post.IsFavorite)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return &posts, nil
}

// GetPostById returns a post from the supplied database based on the given UUID.
func GetPostById(db *sql.DB, id uuid.UUID) (*models.Post, error) {
	row := db.QueryRow(sqlGetPostById, id)

	post := models.Post{}

	err := row.Scan(&post.Id, &post.Tag, &post.Title, &post.ContentMD, &post.ContentHTML, &post.Published, &post.Edited, &post.IsFavorite, &post.Author)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// GetPostByTag returns a post from the supplied database based on the given tag.
func GetPostByTag(db *sql.DB, tag string) (*models.Post, error) {
	row := db.QueryRow(sqlGetPostByTag, tag)

	post := models.Post{}

	err := row.Scan(&post.Id, &post.Tag, &post.Title, &post.ContentMD, &post.ContentHTML, &post.Published, &post.Edited, &post.IsFavorite, &post.Author)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// CreatePost will add a post to the supplied database with the content specified in the method call.
// Fields given `nil` values will be marked as `NULL` in the database.
func CreatePost(db *sql.DB, tag, title, contentMD, contentHTML string, published, edited *time.Time, isFavorite bool, author uuid.UUID) (*models.Post, error) {
	writeDB, err := db.Begin()
	if err != nil {
		writeDB.Rollback()
		return nil, err
	}

	_id, err := uuid.NewV4()
	if err != nil {
		writeDB.Rollback()
		return nil, err
	}

	_, err = writeDB.Exec(sqlCreatePost, _id, tag, title, contentMD, contentHTML, published, edited, isFavorite, author)
	if err != nil {
		writeDB.Rollback()
		return nil, err
	}

	if err = writeDB.Commit(); err != nil {
		return nil, err
	}

	return GetPostById(db, _id)
}

// UpdatePost will update a post in the supplied database, identified by the specified UUID. It will
// fully replace all contents of the post, except for the id field, which will always be ignored.
func UpdatePost(db *sql.DB, id uuid.UUID, tag, title, contentMD, contentHTML string, published, edited *time.Time, isFavorite bool, author uuid.UUID) (*models.Post, error) {
	writeDB, err := db.Begin()
	if err != nil {
		writeDB.Rollback()
		return nil, err
	}

	_, err = writeDB.Exec(sqlUpdatePost, tag, title, contentMD, contentHTML, published, edited, isFavorite, author, id)
	if err != nil {
		writeDB.Rollback()
		return nil, err
	}

	if err = writeDB.Commit(); err != nil {
		return nil, err
	}

	return GetPostById(db, id)
}
