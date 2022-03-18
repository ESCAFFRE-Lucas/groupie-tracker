const checkboxesHtmlCollection = document.querySelectorAll(".inputNumberMembers");
const checkboxes = [...checkboxesHtmlCollection];

const checkboxState = {
    "1": false,
    "2": false,
    "3": false,
}

const artists = [...document.querySelectorAll(".name li")]

//loop on all checkbox to see the numbers of the input
checkboxes.forEach((checkbox) => {
    checkbox.addEventListener("change", (event) => {
        checkboxState[event.target.value] = event.target.checked;
        //creation of a map to have all the members of an artist
        const membersByGroup = artists.map((artist) => {
            return {
                artist: artist.innerHTML,
                members: [...artist.querySelectorAll(".members p")]?.map((member) => member.innerText)
            }
        })
        //condition to have all the artist after the user uncheck all input
        if (checkboxState["1"] === false  && checkboxState["2"] === false && checkboxState["3"] === false) {
            artists.forEach((artist) => {
                artist.style.display = "unset"
            })
        } else {
            for (let i = 0; i < membersByGroup.length; i++) {
                filterByNumberOfMembers(membersByGroup[i])
            }
        }
    })
})

//function that return the good artist for the good checkbox with condition
const filterByNumberOfMembers = (groupInfos) => {
    const nbMembers = groupInfos.members.length;

    artists.forEach((artist) => {
        if (artist.innerHTML !== groupInfos.artist) return
        artist.style.display = "none"
        if ((checkboxState["1"] && nbMembers === 1) || (checkboxState["2"] && nbMembers === 2) || (checkboxState["3"] && nbMembers >= 3)) {
            artist.style.display = "unset"
            console.log("hi")
        }
    })
}