# bla
A blog engine written in Go which emphasizes simplicity.

[![GoDoc](https://godoc.org/github.com/jacobkania/bla?status.svg)](https://godoc.org/github.com/jacobkania/bla)

Bla is a standalone blog engine which allows for simple pages and content editing.

On the first run, Bla will create an SQLite database with the required tables and prompt the owner to create an account. This account will be used to log in from the website for creating and editing blog posts.

Blog posts are meant to be created in markdown. Upon submitting a new post or updating an existing post, the system will compile the markdown into HTML. When a viewer makes a request to the blog post endpoint, the HTML is returned and can be rendered directly by the browser after it is placed into the page with javascript.

A simple HTML page structure is used. Give a few empty div's an appropriate id, and the javascript will handle the rest.

When loading the website's index page, the javascript will make a request to the server for a list of the titles/dates of all posts. It'll then render this list into the page.

When loading an individual blog post, the javascript will make a request to the server for the specific post, and will then inject the response data into the page including the post title, content, and author.

# Creating a blog post

For an admin user, creating a blog post is simple. Just go to https://{{domain}}**/admin** and input the information. Upon clicking submit, the post will be created, and you'll be redirected to the page containing that post.

Blog posts are meant to be written in markdown, and Bla will automatically convert it to HTML when you submit the post. That HTML will then be served whenever the page is requested.

# Editing a blog post

For an admin user, editing a blog post is just as simple as creating one. Go to https://{{domain}}/page/{{tag}}**/admin** and the existing information will be loaded into the form. You can then make any changes and submit the form again, which will overwrite the previous post with the new information.

Flexibility is simple here, so you will have to manually set an edited date, if you would like the post to show that it was edited. Bla will not assume anything for you.