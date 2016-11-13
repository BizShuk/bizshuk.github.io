/**
 * @param {number[]} g
 * @param {number[]} s
 * @return {number}
 */
var findContentChildren = function(g, s) {


    g.sort(function(a,b){ return a-b;});
    s.sort(function(a,b){ return a-b;});
    var howManyGot = 0;
    var cur_child = 0;
    for ( var i = 0 ; i < s.length ; i++ ){

        for ( var j = cur_child ; j < g.length ; j++ ) {
            if ( s[i] >= g[j] ) {
                // child got a cookie with at least minimum size
                howManyGot += 1;
                cur_child = j+1;
                break;
            }

        }
        if (cur_child == g.length) {
            break;
        }
    }
    return howManyGot;
};