// function initMap() {
//     const map = new google.maps.Map(document.getElementById("map"), {
//         center: { lat: 40.76, lng: -73.983 },
//         zoom: 15,
//         mapTypeId: "satellite",
//     });
//
//     map.setTilt(45);
// }

function searchMenu() {
    // Declare variables
    const input = document.getElementById("mySearch");
    const filter = input.value.toUpperCase();
    // ul = document.getElementsByClassName("name");
    // li = ul.getElementsByTagName("li")
    const listItems = document.querySelectorAll(".name li")

    // Loop through all list items, and hide those who don't match the search query
    for ( let i = 0; i < listItems.length; i++) {
        let a = listItems[i].getElementsByTagName("a")[0];
        if (a.innerText.toUpperCase().indexOf(filter) > -1) {
            listItems[i].style.display = "unset";
        } else {
            listItems[i].style.display = "none";
        }
    }
}

function Menu() {
    let x = document.getElementById("myMenu");
    if (x.style.display === "none") {
        x.style.display = "flex";
    } else {
        x.style.display = "none";
    }
}

let slider = document.getElementById("inputDateCreation");
let output = document.getElementById("rangeValue");
output.innerHTML = slider.value;

slider.oninput = function() {
    output.innerHTML = this.value;
}

let sliderAlbum = document.getElementById("inputDateAlbum");
let outputAlbum = document.getElementById("rangeValueAlbum");
outputAlbum.innerHTML = sliderAlbum.value; // Display the default slider value

// Update the current slider value (each time you drag the slider handle)
sliderAlbum.oninput = function() {
    outputAlbum.innerHTML = this.value;
}

function filterDateCreation() {
    const dateList = document.querySelectorAll(".name .dateCreation")
    // console.log(dateList)
    const dateInput = document.getElementById("inputDateCreation").value
    // console.log(dateInput)

    const dates = [...dateList].map( x => x.innerText)
    console.log(dates)

    for (let i = 0; i < dates.length; i++) {
        if (dateInput === dates[i]) {
            const listArtists = document.querySelectorAll(".name li");

            // Loop through all list items, and hide those who don't match the search query
            for ( let i = 0; i < listArtists.length; i++) {
                if (dates[i] === dateInput) {
                    listArtists[i].style.display = "unset";
                } else {
                    listArtists[i].style.display = "none";
                }
            }
        }
    }
}

function filterDateAlbum() {
    const dateList = document.querySelectorAll(".name .dateFirstAlbum")
    console.log(dateList)
    const dateInput = document.getElementById("inputDateAlbum").value
    console.log(dateInput)
    let arrayDateList = [...dateList]


    const test = arrayDateList.map( x => x.innerText)
    console.log(test)

    for (let i = 0; i < arrayDateList.length; i++) {
        test.push(arrayDateList[i])
        if (dateInput === test[i]) {
            const listArtists = document.querySelectorAll(".name li");

            // Loop through all list items, and hide those who don't match the search query
            for ( let i = 0; i < listArtists.length; i++) {
                if (test[i] === dateInput) {
                    listArtists[i].style.display = "unset";
                } else {
                    listArtists[i].style.display = "none";
                }
            }
        }
    }
}

function filterNumberMembers(dateInput) {
    const dateList = document.querySelectorAll(".name .members")
    console.log(dateInput)

    let arrayDateList = [...dateList];
    if (!arrayDateList[dateInput-1].checked){
        // TODO clear filter
        console.log(arrayDateList[dateInput-1])
        return;
    } else {
        console.log(arrayDateList[dateInput-1])
    }


    for (let i = 0; i < dateList.length; i++) {
        arrayDateList.push(dateList[i].innerHTML)
        console.log(typeof dateList[i].innerHTML)
        console.log(typeof arrayDateList)
        if (dateInput === dateList[i].innerHTML.length) {
            const listArtists = document.querySelectorAll(".name li");

            // Loop through all list items, and hide those who don't match the search query
            for ( let i = 0; i < listArtists.length; i++) {
                if (dateList[i].innerHTML.length === dateInput) {
                    listArtists[i].style.display = "unset";
                } else {
                    listArtists[i].style.display = "none";
                }
                console.log("here")
            }
        }
    }
}

function formatDate() {
    let date = document.getElementById("rangeValueAlbum")
    let splitDate = date.split("-")
    console.log(splitDate)
}