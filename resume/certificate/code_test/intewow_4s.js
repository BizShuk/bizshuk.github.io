// you can write to stdout for debugging purposes, e.g.
// console.log('this is a debug message');

function intewow_4(A, D) {
    // write your code in JavaScript (Node.js 6.4.0)

    // position -1 ?
    // get to position N
    // monkey jump 1~D , if D <= N , failed jump
    // A , N = length , index is stone position , val = after val time , the stone come out , -1 means there is no stones at that index

    A.push(0);
    var N = A.length;

    var B = [];
    B.length = N;

    for ( var i = 0 ; i < N ; i++ ) {
        var minimum_time = undefined;

        for ( var j = 1 ; j <= D ; j++) {
            if ( A[i-j] === -1 ) continue;
            var wt = (i-j >= 0) ? B[i-j] : 0;

            if ( (typeof minimum_time == 'undefined') || (minimum_time > wt)) minimum_time = wt;

        }
        if (typeof minimum_time =='undefined' ) return -1;

        B[i] = (A[i] == -1) ? undefined : Math.max(minimum_time,A[i]);
    }

    return B[N-1];

}










