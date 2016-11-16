




// you can write to stdout for debugging purposes, e.g.
// console.log('this is a debug message');

function intewow_2(A, B, M, X, Y) {
    // write your code in JavaScript (Node.js 6.4.0)
    // A for weight , B for target floor , M for floor number of hotel
    // X for people limit , Y for weight limit
    // return stop times , include last stop on ground floor , but not first one


    var ai = 0;

    var npl = 0;    // now people limit
    var nwl = 0;    // now weight limit
    var nt_floors = {}; // now target floor
    var stops = 0;

    while (ai <= A.length) {
        if ( A[ai] > Y || X == 0 ) return false;

        // go up
        if ( npl >= X || nwl+A[ai] > Y || ai === A.length) {

            // calc floor count and +1(go back ground floor)
            stops += Object.keys(nt_floors).length + 1;

            // init count
            npl = 0;
            nwl = 0;
            nt_floors = {};
        }

        npl += 1;
        nwl += A[ai];
        nt_floors[B[ai]] = 1;

        ai++;
    }



    return stops;
}
