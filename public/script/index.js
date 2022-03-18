// function initMap() {
//     const map = new google.maps.Map(document.getElementById("map"), {
//         center: { lat: 40.76, lng: -73.983 },
//         zoom: 15,
//         mapTypeId: "satellite",
//     });
//
//     map.setTilt(45);
// }


// Function that permit the user to search the artist he want by tapping his name
function searchMenu() {
    // Declare variables
    const input = document.getElementById("mySearch");
    const filter = input.value.toUpperCase();
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

//Function that prints the filter menu
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

//function that filter the artist by the date of the creation
function filterDateCreation() {
    const dateList = document.querySelectorAll(".name .dateCreation")
    const dateInput = document.getElementById("inputDateCreation").value
    const dates = [...dateList].map(x => x.innerText)

    //display the user that have the good date of creation
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

//function that filter the artist by the date of his first album
function filterDateAlbum() {
    const dateList = document.querySelectorAll(".name .dateFirstAlbum")
    const dateUserInput = formatDate()
    //take all the array one by one with ...
    let arrayDateList = [...dateList]

    //creation of a map to compare with the good input later
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

//function that split the date of the date input and put the letters in the correct order because the input (type = date)
// returns the date upside down
function formatDate() {
    let date = document.getElementById("rangeValueAlbum").innerText
    let splitDate = date.split("-")
    let dateYear = splitDate[0]
    splitDate[0] = splitDate[2]
    splitDate[2] = dateYear
    splitDate = splitDate[0] + '-' + splitDate[1] + '-' + splitDate[2]
    return splitDate
}
