/**
 * @param {string} str
 * @return {boolean}
 */
var repeatedSubstringPattern = function(str) {
    var index_mc = [];
    var index = 0;
    console.log("1");
    for (var i = 0 ; i < str.length ; i++){
        if ( index_mc[0] === str[i]) {
            // check all sub
            var mc_len = index_mc.length;
            var loopindex = 0;
            var suc = true;
            for(var j = i;j <str.length ; j++ ,loopindex++){
                //console.log(str,j,str[j],index_mc,loopindex%mc_len,index_mc[loopindex%mc_len]);
                if (str[j] !== index_mc[ loopindex % mc_len]){
                    suc = false;
                    break;
                }
            }
            loopindex--;
            j--;
            // console.log(suc,index_mc,loopindex,j);
            if (suc && (loopindex% mc_len) == index_mc.length-1) return true;
        }


        index_mc.push(str[i]);
    }
    return false;

};
