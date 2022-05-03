const ws = new WebSocket("ws://localhost:5487/ws");

ws.onopen = function () {
    console.log("Connected to server");
};

ws.onmessage = function (evt) {
    var { name, artist, album, image } = JSON.parse(evt.data);

    if (name) {
        document.getElementById("box").style.visibility = "visible";
    } else {
        document.getElementById("box").style.visibility = "hidden";
    }
    document.getElementById("trackName").innerHTML = name;
    document.getElementById("trackArtist").innerHTML = artist;
    document.getElementById("trackAlbum").innerHTML = album;
    document.getElementById("image").src = `http://${image}?t=${new Date().getTime()}`;

    let i = 0;
    let setImageInterval = setInterval(function () {
        document.getElementById("image").src = `http://${image}?t=${new Date().getTime()}`;
        i++;
        if (i > 40) {
            clearInterval(setImageInterval);
        }
    }, 10);
};
