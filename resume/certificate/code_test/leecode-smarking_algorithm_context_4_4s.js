/**
 * @param {number[]} A
 * @param {number[]} B
 * @param {number[]} C
 * @param {number[]} D
 * @return {number}
 */
var fourSumCount = function(A, B, C, D) {

    var ret_count = 0;
    Ah = {};
    for(var i=0;i<A.length;i++){
        for(var j=0;j<B.length;j++){
            var sum = A[i] + B[j];
            if ( typeof Ah[sum] === 'undefined' ) {
                Ah[sum] = 0;
            }
            Ah[sum]++;
        }
    }

    for(var i=0;i<C.length;i++){
        for(var j=0;j<D.length;j++){
            var sum = C[i] + D[j];

            if(Ah[-sum]){
                ret_count += Ah[-sum];
            }


        }
    }




    return ret_count;

};