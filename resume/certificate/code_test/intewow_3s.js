// you can write to stdout for debugging purposes, e.g.
// console.log('this is a debug message');

function intewow_3(A) {
    // write your code in JavaScript (Node.js 6.4.0)

    // A , two dimension array , value is color type
    // return with how many countries is in the matrixs
    var row = A.length;
    var column = A[0].length;
    var count = row * column;

    for (var i = 0; i < row ; i++) {
        for (var j = 0 ; j < column ; j++) {
            if ( i<row-1 && A[i][j] === A[i+1][j] ) count--;
            if ( j<column-1 && A[i][j] === A[i][j+1] ) count--;
        }
    }
    return count;

}
