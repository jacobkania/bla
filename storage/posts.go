package storage

import "bla/models"

const sqlGetAllPosts string = `SELECT tag, title, published FROM posts`
const sqlGetAllFavoritePosts string = `SELECT tag, title, published FROM posts WHERE is_favorite = TRUE`
const sqlGetPostById string = `SELECT id, tag, title, content_md, content_html, published, edited, is_favorite, author FROM posts WHERE id = ?`
const sqlGetPostByTag string = `SELECT id, tag, title, content_md, content_html, published, edited, is_favorite, author FROM posts WHERE tag = ?`

const sqlCreatePost string = `INSERT INTO posts (id, tag, title, content_md, content_html, published, edited, is_favorite, author) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
const sqlUpdatePost string = `UPDATE posts SET tag = ?, title = ?, content_md = ?, content_html = ?, published = ?, edited = ?, is_favorite = ?, author = ?`

func GetAllPosts() (*[]models.PostLite, error) {
	rows, err := readDB.Query(sqlGetAllPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.PostLite

	for rows.Next() {
		post := models.PostLite{}
		err = rows.Scan(&post.Tag, &post.Title, &post.Published)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return &posts, nil
}

func GetAllFavoritePosts() (*[]models.PostLite, error) {
	rows, err := readDB.Query(sqlGetAllFavoritePosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.PostLite

	for rows.Next() {
		post := models.PostLite{}
		err = rows.Scan(&post.Tag, &post.Title, &post.Published)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return &posts, nil
}

func GetPostById(id string) (*models.Post, error) {
	row := readDB.QueryRow(sqlGetPostById, id)

	post := models.Post{}

	err := row.Scan(&post.Id, &post.Tag, &post.Title, &post.ContentMD, &post.ContentHTML, &post.Published, &post.Edited, &post.IsFavorite, &post.Author)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func GetPostByTag(tag string) (*models.Post, error) {
	row := readDB.QueryRow(sqlGetPostByTag, tag)

	post := models.Post{}

	err := row.Scan(&post.Id, &post.Tag, &post.Title, &post.ContentMD, &post.ContentHTML, &post.Published, &post.Edited, &post.IsFavorite, &post.Author)
	if err != nil {
		return nil, err
	}

	return &post, nil
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

func UpdatePost(post *models.Post) error {
	writeDB, err := readDB.Begin()
	if err != nil {
		writeDB.Rollback()
		return err
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
		post.Author)
	if err != nil {
		writeDB.Rollback()
		return err
	}

	return writeDB.Commit()
}
