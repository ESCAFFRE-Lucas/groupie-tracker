const checkboxesHtmlCollection = document.querySelectorAll(".inputNumberMembers");
const checkboxes = [...checkboxesHtmlCollection];

const checkboxState = {
    "1": false,
    "2": false,
    "3": false,
}

const artists = [...document.querySelectorAll(".name li")]

const toggleVisibility = (artist, visible) => {
    for (let i = 0; i < artist.length; i++) {
        artist[i].style.display = visible ? "unset" : "none";
    }
}

checkboxes.forEach((checkbox) => {
    checkbox.addEventListener("change", (event) => {
        // console.log(event.target.checked, event.target.value);
        checkboxState[event.target.value] = event.target.checked;
        console.log(checkboxState)
        const members = artists.map((artist) => {
            return [...artist.querySelectorAll(".members p")]?.map((member) => member.innerText)
        })
        for (let i = 0; i < members.length; i++) {
            if (members[i].length === 1 && checkboxState[event.target.value] === true) {
                for ( let i = 0; i < artists.length; i++) {
                    if (members[i].length === 1) {
                        artists[i].style.display = "unset";
                    } else {
                        artists[i].style.display = "none";
                    }
                }
            }
        }
    })
})