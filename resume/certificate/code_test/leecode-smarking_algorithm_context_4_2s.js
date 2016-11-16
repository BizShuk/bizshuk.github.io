/**
 * @param {number[]} nums
 * @return {boolean}
 */
var find132pattern = function(nums) {


    var i = 0;
    var f=[];
    var c=[];
    var count_m = 0;

    var start;
    var center;
    while ( i<nums.length ){
        var num = nums[i];

        //console.log(f,c,start,center,num);
        //
        for (var j = 0 ; j < count_m ; j++){
            if ( f[j] < num && num < c[j]){
                return true;
            }
        }
        if ( start < num && num < center) return true;



        if (typeof start == "undefined") {
            start = num;
            center = undefined;
        }

        if (start > num && typeof center === "undefined") {
            start = num;
        }

        if (start < num && typeof center === "undefined") {
            center = num;
        }

        if (start < num && center < num) {
            center = num;
        }

        if (start > num && center > num) {
            f.push(start);
            c.push(center);
            count_m++;
            center = undefined;
            start = num;
        }


        i++;
    }


    return false;
};
