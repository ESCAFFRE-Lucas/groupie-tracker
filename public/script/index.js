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
    for (let i = 0; i < listItems.length; i++) {
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

slider.oninput = function () {
    output.innerHTML = this.value;
}

let sliderAlbum = document.getElementById("inputDateAlbum");
let outputAlbum = document.getElementById("rangeValueAlbum");
outputAlbum.innerHTML = sliderAlbum.value; // Display the default slider value

// Update the current slider value (each time you drag the slider handle)
sliderAlbum.oninput = function () {
    outputAlbum.innerHTML = this.value;
}

function filterDateCreation() {
    const dateList = document.querySelectorAll(".name .dateCreation")
    const dateInput = document.getElementById("inputDateCreation").value
    const dates = [...dateList].map(x => x.innerText)

    for (let i = 0; i < dates.length; i++) {
        if (dateInput === dates[i]) {
            const listArtists = document.querySelectorAll(".name li");

            // Loop through all list items, and hide those who don't match the search query
            for (let i = 0; i < listArtists.length; i++) {
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
    const dateUserInput = formatDate()
    let arrayDateList = [...dateList]


    const mapOfArrayDate = arrayDateList.map(x => x.innerText)

    for (let i = 0; i < arrayDateList.length; i++) {
        mapOfArrayDate.push(arrayDateList[i])
        if (dateUserInput === mapOfArrayDate[i]) {
            const listOfArtists = document.querySelectorAll(".name li");
            // Loop through all list items, and hide those who don't match the search query
            for (let i = 0; i < listOfArtists.length; i++) {
                if (mapOfArrayDate[i] === dateUserInput) {
                    listOfArtists[i].style.display = "unset";
                } else {
                    listOfArtists[i].style.display = "none";
                }
            }
        }
    }
}

function formatDate() {
    let date = document.getElementById("rangeValueAlbum").innerText
    let splitDate = date.split("-")
    let dateYear = splitDate[0]
    splitDate[0] = splitDate[2]
    splitDate[2] = dateYear
    splitDate = splitDate[0] + '-' + splitDate[1] + '-' + splitDate[2]
    return splitDate
}
