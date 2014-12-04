// Project Euler 1 (PE1)
function multof5() {
    var arr = [],
        z = 0;
    for (var i = 3; i < 1000; i++) {
        // prevent repeated multiples of 3 and 5
        if (i % 3 === 0 && i % 5 != 0) {
            arr.push(i)
        }
        if (i % 5 === 0) {
            arr.push(i)
        }
    }
    // add all elements in array
    for (var x = 0; x < arr.length; x++) {
        z += arr[x];
    }
    console.log(z);
}

// PE2
function fibo() {
    var z, x = 0,
        y = 1,
        a = 0,
        arr = [];
    for (var i = 0; i < 32; i++) {
        z = x + y;
        x = y;
        y = z;
        if (z % 2 === 0) {
            arr.push(z);
        }
    }
    //console.log(z);

    // add em up:
    for (var g = 0; g < arr.length; g++) {
        a += arr[g];
    }
    console.log(a);
}