window.onload = function() {
    var user_id = getCookie('userID')
    document.getElementById('username').innerText = "UserID: " + user_id
    leadersRPC()
};

function getCookie(cname) {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

async function postData(url = '', data = {}) {
    // Default options are marked with *
    const response = await fetch(url, {
        method: 'POST', // *GET, POST, PUT, DELETE, etc.
        mode: 'same-origin', // no-cors, *cors, same-origin
        cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
        credentials: 'same-origin', // include, *same-origin, omit
        headers: {
            'Content-Type': 'application/json'
        },
        redirect: 'follow', // manual, *follow, error
        referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
        body: JSON.stringify(data) // body data type must match "Content-Type" header
    });
    return response.json(); // parses JSON response into native JavaScript objects
}

async function clickRPC() {
    var user_id = getCookie('userID')
    if (user_id === "") {
        alert("You must login to do that!")
    }
    path = '/shitlist.v1.ShitlistService/Click'
    data = {
        user_id: getCookie('userID')
    }
    postData(path, data)
        .then(data => {
            document.getElementById('score').innerText = "Score: " + data.clicks
        });
}

async function leadersRPC() {
    path = '/shitlist.v1.ShitlistService/Leaders'
    data = {}
    postData(path, data).then(data => { updateLeaders(data) });
}

function updateLeaders(data = {}) {
    if (data === {}) return;
    var user_id = getCookie('userID')

    const leaderBoardID = 'leaderboard'
    var table = document.createElement('table');
    table.id = leaderBoardID
    var header = table.insertRow()

    var userID = header.insertCell()
    userID.textContent = 'UserID'

    var clicks = header.insertCell()
    clicks.textContent = 'Clicks'

    data.topClickers.forEach(function(rowData) {
        var row = table.insertRow()
        var uid = row.insertCell()
        var clicks = row.insertCell()

        uid.textContent = rowData.userId
        clicks.textContent = rowData.clicks

        if (rowData.userID == user_id) {
            document.getElementById('score').innerText = "Score: " + rowData.clicks
        }
    });
    document.getElementById(leaderBoardID).replaceWith(table)
}