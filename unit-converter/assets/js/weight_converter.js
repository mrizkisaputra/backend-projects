export default
function convertWeight({inputWeight, unitFrom, unitTo}) {
    let unit = weightUnits[unitFrom][unitTo]
    let converted = inputWeight * unit
    $("[name='to-weight']").val(converted+" "+unitTo)
}

const weightUnits = {
    "mg": {
        "mg": 1,
        "g": 0.001,
        "kg": 0.000001,
        "ounce": 0.000035274,
        "pound": 0.00000220462,
    },
    "g": {
        "mg": 1000,
        "g": 1,
        "kg": 0.001,
        "ounce": 0.035274,
        "pound": 0.00220462,
    },
    "kg": {
        "mg": 1e+6,
        "g": 1000,
        "kg": 1,
        "ounce": 35.274,
        "pound": 2.20462,
    },
    "ounce": {
        "mg": 28349.5,
        "g": 28.3495,
        "kg": 0.0283495,
        "ounce": 1,
        "pound": 0.0625,
    },
    "pound": {
        "mg": 453592,
        "g": 453.592,
        "kg": 0.453592,
        "ounce": 16,
        "pound": 1,
    }
}