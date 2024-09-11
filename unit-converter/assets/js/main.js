import convertLength from "./length_converter.js";
import convertWeight from "./weight_converter.js"
import convertTemperature from "./temperature_converter.js";

/**
 * if document is loaded
 * load the length form by default
 */
$(document).on("DOMContentLoaded", function () {
    loadForm("_length.html", $("#btn-length"))
})

/**
 * if unitconverter weight click
 */
$("#btn-weight").on("click", function () {
    loadForm("_weight.html", this)
})

/**
 * if unitconverter length click
 */
$("#btn-length").on("click", function () {
    loadForm("_length.html", this)
})

/**
 * if unitconverter temperature click
 */
$("#btn-temperature").on("click", function () {
    loadForm("_temperature.html", this)
})

function loadForm(file, $self) {
    $("#container-form-converter").load(file)
    $(".tab-button").removeClass("active")
    $($self).addClass("active")
}

/**
 * submit length converter
 */
$(document).on("submit", "#form-length", function (event) {
    event.preventDefault(); // mencegah refresh halaman

    const formData = $(this).serializeArray()
    let data = {}
    formData.forEach(item => {
        data[item.name] = item.value
    })
    convertLength(data)
});

/**
 * submit weight converter
 */
$(document).on("submit", "#form-weight", function (event) {
    event.preventDefault(); // mencegah refresh halaman

    const formData = $(this).serializeArray()
    let data = {}
    formData.forEach(item => {
        data[item.name] = item.value
    })
    convertWeight(data)
});

/**
 * submit temperature converter
 */
$(document).on("submit", "#form-temperature", function (event) {
    event.preventDefault(); // mencegah refresh halaman

    const formData = $(this).serializeArray()
    let data = {}
    formData.forEach(item => {
        data[item.name] = item.value
    })
    convertTemperature(data)
});
