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
    userID = 'ddc7a2d5-a617-4846-bc06-5bcd58c64091' // TODO: get userID from somewhere

    path = '/shitlist.v1.ShitlistService/Click'
    data = {
        user_id: userID
    }
    postData(path, data)
        .then(data => {
            console.log(data); // JSON data parsed by `data.json()` call
        });
}