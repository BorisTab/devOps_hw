function getVal() {
    var val = document.getElementById("get_key").value;

    fetch("/get?key=" + val)
    .then((response) => response.json())
    .then((data) => {
        document.getElementById("get_answer").innerHTML = data.value;
    });
}