const baseUrl = 'https://localhost:8081/post/tag/';
const postId = window.location.pathname.split('/')[2];
const monthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];

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

function putContentInPage(data) {
    document.getElementById("post").innerHTML = data.contentHtml;
    document.title = "Blog | " + data.title;
    document.getElementById("title").innerText = data.title;
    document.getElementById("author").innerText = data.author || '';
    document.getElementById("created-date").innerText = data.published ? formatDateFromString(data.published) : '';
    document.getElementsByClassName("info-edited").item(0).innerHTML = data.edited
        ? 'Edited <span id="edited-date">' + formatDateFromString(data.edited) + '</span>'
        : '';
}

fetch(baseUrl + postId)
    .then((res) => res.json())
    .then((data) => putContentInPage(data))
    .then(() => Prism.highlightAll());