window.onload = function() {
    userID = uuidv4()
    document.getElementById('username').innerText = "UserID: " + userID
    document.getElementById('score').innerText = "Score: 0"
    leadersRPC()
};

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
    path = '/shitlist.v1.ShitlistService/Click'
    data = {
        user_id: userID
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
    });
    document.getElementById(leaderBoardID).replaceWith(table)
}

function newUserID() {
    userID = uuidv4();
    document.getElementById('username').innerText = "UserID: " + userID
    document.getElementById('score').innerText = "Score: 0"
}

function uuidv4() {
    return ([1e7] + -1e3 + -4e3 + -8e3 + -1e11).replace(/[018]/g, c =>
        (c ^ crypto.getRandomValues(new Uint8Array(1))[0] & 15 >> c / 4).toString(16)
    );
}