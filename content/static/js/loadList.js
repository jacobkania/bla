const baseUrl = 'http://localhost:8081/post';
const favoriteUrl = 'http://localhost:8081/favorite';
const monthNames = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];

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

function generateListItem(item) {
    var wrapper = document.createElement('div');
    wrapper.setAttribute('class', 'item-wrapper');

    var link = document.createElement('a');
    link.setAttribute('href', '/page/' + item.tag);

    var dateItem = document.createElement('span');
    dateItem.setAttribute('class', 'item-date');

    var dateText = document.createTextNode(formatDateFromString(item.published));

    dateItem.appendChild(dateText);
    link.appendChild(dateItem);

    var fav = document.createElement('span');
    fav.setAttribute('class', 'item-fav');
    if (item.isFavorite) {
        var favImg = document.createElement('img');
        favImg.setAttribute('src', '/img/star.png');
        favImg.setAttribute('width', '16');
        favImg.setAttribute('height', '16');
        favImg.setAttribute('alt', 'favorite');

        fav.appendChild(favImg);
    }
    link.appendChild(fav);

    var title = document.createElement('span');
    title.setAttribute('class', 'item-title');

    var titleText = document.createTextNode(item.title);

    title.appendChild(titleText);
    link.appendChild(title);

    wrapper.appendChild(link);

    return wrapper;
}

function putContentInPage(data) {
    data.sort((a, b) => Date.parse(a.published) < Date.parse(b.published));
    data.forEach(item => {
        document.getElementById('items').appendChild(generateListItem(item));
    });
}

fetch(baseUrl)
    .then((res) => res.json())
    .then((data) => putContentInPage(data));