/*  
    Required: jquery 1.9.1+  
    ******* important ******** 
    if you use border on the elem , it delay few milliseconds for render ,so the top may not be correct.

    parameters:
        elem_id: a string for id of elem
        scroll_type: a number , only one scroll_type now


    bug:
        remove "on" class 時 top沒歸到0px ???
*/
function elem_fixed_on_top(elem_id,scroll_type){

    var last_elem_top = 0;
    var last_scroll_top = 0;
    var dmdc = 0; // diff move distance count

    var elem = $('#'+elem_id);
    elem.addClass('elem_scroll_style1');    // init class

    var elem_height = elem.outerHeight();
    var ori_elem_top = elem.position().top;
    var ori_elem_bottom = ori_elem_top + elem_height;
    var placeholder = elem_height*2;

    // simular to  android chrome mobile app's url bar , unlimited by initial elem.top==0
    function elem_fixed_1(event){
        var doc_top = $(document).scrollTop();

        var move_distance = last_scroll_top - doc_top;
        dmdc += move_distance;

        // relative時 css top:0  其position().top = 46
        // fixed時    css top:0  其position().top = 0
        var elem_top = parseInt(elem.position().top);

        // scroll 小於原本的位置bottom 
        console.log(doc_top,ori_elem_bottom);
        if ( doc_top <= ori_elem_bottom ) {
            console.log(1);
            // [有點難解釋]捲軸高度 +???  小於 elem 原位置底部
            if ( doc_top + elem_height + elem_top <= ori_elem_bottom ) {
                console.log("set default");
                elem.css({'top':'0px'});
                elem.removeClass('elem_scroll_style1_on');
            }

        } else {    // bug css top vs position top混用錯誤
            // scroll up 超過準備區 且 elem 尚未出現螢幕 , if relative => elem不會再隱藏位置 ,  only fixed and hidden will equal -elem_height
            if (doc_top <= ori_elem_bottom + placeholder){

                if( parseInt(elem_top) == parseInt(-elem_height) &&
                elem.hasClass('elem_scroll_style1_on')
                ){
                    elem.removeClass('elem_scroll_style1_on');
                    elem.css({'top': '0px'});
                    console.log("buffer area");
                }
                return;
            }

            {   
                // 變fixed前 預先移動到螢幕外
                if ( !elem.hasClass('elem_scroll_style1_on') ){
                    elem.css({'top': (-elem_height)+'px'});
                    console.log("pre_action for fixed");
                }
                //fixed預先移動完成 即加上fixed
//                console.log(elem_top , -elem_height + ori_elem_top);
                if ( parseInt(elem_top) == parseInt(-elem_height + ori_elem_top) ){ 
                    console.log("add fixed");
                    elem.addClass('elem_scroll_style1_on');
                }
            }

            // 超出高度
            if ( dmdc > 50 ){
            console.log(" show on top");
                dmdc = 50;
                elem.css({"top": "0px"});
            }
            if (dmdc < -50) {
            console.log(" hide on top");
                dmdc = -50;
                elem.css({"top": (-elem_height)+"px"});
            }
        }

        last_scroll_top = doc_top;
    }

    $(document).on("scroll", elem_fixed_1);
}

// ex: elem_fixed_on_top('input_article',1);
