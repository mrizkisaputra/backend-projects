export default
function convertLength({inputLength, unitFrom, unitTo}) {
    let unit = lengthUnits[unitFrom][unitTo]
    let converted = parseFloat(inputLength) * unit

    $("[name='to-length']").val(converted+" "+unitTo)
}

const lengthUnits = {
    "mm": {
        "mm": 1,
        "cm": 0.1,
        "m": 0.001,
        "km": 0.000001,
        "inch": 0.03937,
        "foot": 0.003281,
        "yard": 0.001094,
        "mile": 6.2137e-7,
    },
    "cm": {
        "mm": 10,
        "cm": 1,
        "m": 0.01,
        "km": 0.00001,
        "inch": 0.3937,
        "foot": 0.03281,
        "yard": 0.01094,
        "mile": 6.2137e-6,
    },
    "m": {
        "mm": 1000,
        "cm": 100,
        "m": 1,
        "km": 0.001,
        "inch": 39.3701,
        "foot": 3.28084,
        "yard": 1.09361,
        "mile": 0.000621371,
    },
    "km": {
        "mm": 1e+6,
        "cm": 100000,
        "m": 1000,
        "km": 1,
        "inch": 39370.1,
        "foot": 3280.84,
        "yard": 1093.61,
        "mile": 0.621371,
    },
    "inch": {
        "mm": 25.4,
        "cm": 2.54,
        "m": 0.0254,
        "km": 0.0000254,
        "inch": 1,
        "foot": 0.08333,
        "yard": 0.02778,
        "mile": 0.000015783,
    },
    "foot": {
        "mm": 304.8,
        "cm": 30.48,
        "m": 0.3048,
        "km": 0.0003048,
        "inch": 12,
        "foot": 1,
        "yard": 0.33333,
        "mile": 0.000189394,
    },
    "yard": {
        "mm": 914.4,
        "cm": 91.44,
        "m": 0.9144,
        "km": 0.0009144,
        "inch": 36,
        "foot": 3,
        "yard": 1,
        "mile": 0.000568182,
    },
    "mile": {
        "mm": 1.609e+6,
        "cm": 160934,
        "m": 1609.34,
        "km": 1.60934,
        "inch": 63360,
        "foot": 5280,
        "yard": 1760,
        "mile": 1,
    }
}

