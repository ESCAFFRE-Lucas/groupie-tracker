const checkboxesHtmlCollection = document.querySelectorAll(".inputNumberMembers");
const checkboxes = [...checkboxesHtmlCollection];

const checkboxState = {
    "1": false,
    "2": false,
    "3": false,
}

const artists = [...document.querySelectorAll(".name li")]

checkboxes.forEach((checkbox) => {
    checkbox.addEventListener("change", (event) => {
        // console.log(event.target.checked, event.target.value);
        checkboxState[event.target.value] = event.target.checked;
        console.log(checkboxState)
        const membersByGroup = artists.map((artist) => {
            return {
                artist: artist.innerHTML,
                members: [...artist.querySelectorAll(".members p")]?.map((member) => member.innerText)
            }
        })
        console.log(checkboxState)
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

const filterByNumberOfMembers = (groupInfos) => {
    const nbMembers = groupInfos.members.length;
    console.log(groupInfos, checkboxState["1"], (checkboxState["1"] && nbMembers === 1))
    artists.forEach((artist) => {
        if (artist.innerHTML !== groupInfos.artist) return
        artist.style.display = "none"
        // console.log((checkboxState["1"] && nbMembers === 1) || (checkboxState["2"] && nbMembers === 2) || (checkboxState["3"] && nbMembers >= 3))
        if ((checkboxState["1"] && nbMembers === 1) || (checkboxState["2"] && nbMembers === 2) || (checkboxState["3"] && nbMembers >= 3)) {
            artist.style.display = "unset"
            console.log("hi")
        }
    })
}