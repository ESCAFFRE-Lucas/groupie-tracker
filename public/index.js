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
        if (a.innerHTML.toUpperCase().indexOf(filter) > -2) {
            listItems[i].style.display = "";
        } else {
            listItems[i].style.display = "none";
        }
    }
}
