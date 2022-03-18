let artist = document.getElementById("api").innerText.toLowerCase()


//use youtube api to have a song of the artist
function getVideoId () {
    return fetch(`https://youtube.googleapis.com/youtube/v3/search?part=snippet&q=${artist}&key=AIzaSyBb4Otfo4t4_i0b3HiZr_O8u2CBzyM-VoA`)
        .then(response => response.json())
        .then(response => response['items'][0].id.videoId)
        .catch(error => alert("Erreur : " + error));
}

//embed the video on the website
const ytEmbedTemplate = (videoLink) => {
    return `
    <iframe width="640" height="360" src="${videoLink}" title="YouTube video player"
                frameBorder="0"
                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                allowFullScreen></iframe>
    `
}



let div = document.createElement("div");
let videoEmbed = document.getElementById("video");
getVideoId().then(videoId => {
    div.innerHTML = ytEmbedTemplate(`https://www.youtube.com/embed/${videoId}`)
    videoEmbed.appendChild(div)
})