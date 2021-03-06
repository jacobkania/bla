import {SERVER_URL} from "./config.js";

const baseUrl = SERVER_URL + '/post/tag/';
const pageUrl = SERVER_URL + '/page/';
const postUrl = SERVER_URL + '/post';
const postTag = window.location.pathname.split('/')[2];

const post = {
    title: "",
    contentMd: "",
    isFavorite: false,
    published: Date(),
    edited: Date(),
    tag: "",
    id: ""
};

function createDateStringFromFullDate(fullDate) {
    if (!fullDate)
        return null;

    var year = fullDate.getFullYear();

    var month = fullDate.getMonth() + 1;

    var day = fullDate.getDate().toString();

    return year + '-' + month + '-' + day;
}

function createDateFromDashSeparatedString(dateString) {
    if (!dateString)
        return null;

    let year = Number(dateString.split('-')[0]);
    let month = Number(dateString.split('-')[1]);
    let day = Number(dateString.split('-')[2]);

    return new Date(year, month - 1, day);
}

function putContentInPage() {
    document.title = "Admin | " + post.title;
    document.getElementById("title-field").value = post.title;
    document.getElementById("content-field").value = post.contentMd;
    document.getElementById("favorite-field").checked = post.isFavorite;
    document.getElementById("created-field").value = post.published ? createDateStringFromFullDate(post.published) : '';
    document.getElementById("edited-field").value = post.edited ? createDateStringFromFullDate(post.edited) : '';
    document.getElementById("tag-field").value = post.tag;
}

function updatePost(request) {
    const Http = new XMLHttpRequest();
    Http.open("PUT", postUrl + "/id/" + post.id, true);
    Http.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    Http.send(JSON.stringify(request));
    Http.onreadystatechange=()=>{
        if (Http.readyState === 4) {
            if (Http.status !== 200)
                alert("Failed to update post: " + Http.statusText);
            else
                window.location.href = pageUrl + post.tag;
        }
    }
}

function createPost(request) {
    const Http = new XMLHttpRequest();
    Http.open("POST", postUrl, true);
    Http.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    Http.send(JSON.stringify(request));
    Http.onreadystatechange=()=> {
        if (Http.readyState === 4) {
            if (Http.status !== 201)
                alert("Failed to create post: " + Http.statusText);
            else
                window.location.href = pageUrl + post.tag;
        }
    }
}

function submitForm() {
    post.title = document.getElementById("title-field").value;
    post.contentMd = document.getElementById("content-field").value;
    post.isFavorite = document.getElementById("favorite-field").checked === true;
    post.published = createDateFromDashSeparatedString(document.getElementById("created-field").value);
    post.edited = createDateFromDashSeparatedString(document.getElementById("edited-field").value);
    post.tag = document.getElementById("tag-field").value;
    post.id = post.id ? post.id : null;

    let request = {
        post: post,
        username: document.getElementById("login-field").value,
        password: document.getElementById("password-field").value
    };

    if (window.location.pathname !== "/admin")
        updatePost(request);
    else
        createPost(request);
}

function deleteForm() {
    if (window.location.pathname === '/admin') {
        let actuallyClearForm = confirm("Are you sure you'd like to clear this form?");
        if (!actuallyClearForm) {
            return
        }

       document.getElementById("title-field").value = '';
       document.getElementById("content-field").value = '';
       document.getElementById("favorite-field").checked = false;
       document.getElementById("created-field").value = '';
       document.getElementById("edited-field").value = '';
       document.getElementById("tag-field").value = '';

        return
    }

    let actuallyDelete = confirm("Are you sure you'd like to permanently delete this post?");
    if (!actuallyDelete) {
        return
    }

    let request = {
        username: document.getElementById("login-field").value,
        password: document.getElementById("password-field").value
    };

    const Http = new XMLHttpRequest();
    Http.open("DELETE", postUrl + "/id/" + post.id, true);
    Http.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    Http.send(JSON.stringify(request));
    Http.onreadystatechange=()=> {
        if (Http.readyState === 4) {
            if (Http.status !== 200)
                alert("Failed to update post: " + Http.statusText);
            else
                alert("Deleted post");
        }
    }
}

if (window.location.pathname !== "/admin") {
    fetch(baseUrl + postTag)
        .then((res) => res.json())
        .then((data) => {
            post.title = data.title;
            post.contentMd = data.contentMd;
            post.isFavorite = data.isFavorite;
            post.published = new Date(data.published);
            post.edited = data.edited ? new Date(data.edited) : null;
            post.tag = data.tag;
            post.id = data.id;
        })
        .then(() => putContentInPage());
}

document.getElementById("submit-form").onclick = () => submitForm();

document.getElementById("delete-form").onclick = () => deleteForm();
