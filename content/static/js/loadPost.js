import {SERVER_URL} from "./config.js";

const postTagUrl = SERVER_URL + '/post/tag/';
const authorUrl = SERVER_URL + '/user/id/';
const postId = window.location.pathname.split('/')[2];
const monthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];

const post = {
    title: "",
    contentHtml: "",
    published: Date(),
    edited: Date(),
    tag: "",
    id: "",
    author: ""
};

function formatDateFromString(dateString) {
    var date = new Date(dateString);

    var year = date.getFullYear();

    // var month = (date.getMonth() + 1).toString();
    // month = month.length > 1 ? month : '0' + month;
    var month = monthNames[date.getMonth()];

    var day = date.getDate().toString();
    day = day.length > 1 ? day : '0' + day;

    return year + '-' + month + '-' + day;
}

function putContentInPage() {
    document.getElementById("post").innerHTML = post.contentHtml;
    document.title = "Blog | " + post.title;
    document.getElementById("title").innerText = post.title;
    document.getElementById("created-date").innerText = post.published ? formatDateFromString(post.published) : '';
    document.getElementsByClassName("info-edited").item(0).innerHTML = post.edited
        ? 'Edited <span id="edited-date">' + formatDateFromString(post.edited) + '</span>'
        : '';
}

function getAuthorById(id) {
    fetch(authorUrl + id)
        .then((res) => res.json())
        .then((data) => {
            document.getElementById("author").innerText = data.firstName + " " + data.lastName;
        })
}

fetch(postTagUrl + postId)
    .then((res) => res.json())
    .then((data) => {
        post.title = data.title;
        post.contentHtml = data.contentHtml;
        post.published = new Date(data.published);
        post.edited = data.edited ? new Date(data.edited) : null;
        post.tag = data.tag;
        post.id = data.id;
        post.author = data.author;
    })
    .then(() => putContentInPage())
    .then(() => Prism.highlightAll())
    .then(() => getAuthorById(post.author));