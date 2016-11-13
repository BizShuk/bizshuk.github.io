// you can write to stdout for debugging purposes, e.g.
// console.log('this is a debug message');

function intewow_1(A) {
    // write your code in JavaScript (Node.js 6.4.0)
    if (A.length ===0) return 0;

    var k = A[0]
    var count = 0;
    while (k !== -1) {
        count++;
        k = A[k];
    }
    count++;
    return count;
}