問卷doc

db.event 搜尋survay%
ps:survay_model_matrix_all 用create view 建起來的 , 如果更新請參閱survay.inc.php::get_survay_detail的 $matrix_sql without where
單純方便填寫頁面使用

後台
資料大致model
survay
     page
          matrixs
               questions
                    options
                         options2 [order]


index.php

     editor.php(angularjs 寫的)     => survay_dump.pl,survay_dump_10.pl,api.php
          survay_setup.js     (dump)
          survay.inc.php
          ps:f12 console.log(questions);
          **********************
          單選與複選 其他欄是掛在 matrix.otheroption上
          但分類單選為了兼容每個選項都可以自填 在db.survay_model_option多了op_type(0:default,1:自填框)

          分類單選 db.survay_model_option的op_category為分類的類別 其類別名稱掛在op_name "-"的前面

          在table query時,須注意op_id 與op_category的順序 ,op_id順序要對 進array資料才會對 ,但op_category順序要對 顯示才會是正常的分類順序

     analyze.php                    => analyze_answer.dump.pl,show_answers_list.php
          survay.inc.php
          analyze.inc.js(filter and cross operations)
          analyze.inc2.js(show_tables function)
          ps:f12 console.log(survay_model,formated_data,answers);

     generate_table_schema.php 產生table schema的php



survay_dump[_10].pl
     @h_ds = ({"userid":""},{"extra_data":[]},{"matrixs":[]})     => 改問卷的model

     1. 組問卷model
     2. 欄位名稱(excel第一行)
     3. 拉user answers
     4. output user line




玩家顯示

index.php
     dump出頁數與題型 傳給下面php產生html(並會拉出有題目過濾選項過濾的答案做判定)

     survay.inc.php
          option.inc.php
               option_view.inc.php

     survay.inc.js
          1. 特殊設定時使用的js, ex:排序不能重複,選擇陰影框 ,分類單選的+號, 數字意見欄的填寫限定
          2. 拉整頁的資料 (getallanswer) 送到api.php


api.php ( 儲存填寫資料的地方 )
     api.inc.php

     starttime 進入php時間
     client_startime 第一頁讀取完畢時間
     end_time 做答完成時間












 
==================================================================

# 變數設定那邊 有分兩種, 一種是產品那邊帶(EX: uid), 一種是我們這邊會帶(ex: 國家); 前者需要從edit那邊建變數(存到db裡), 後者在index,php時會自己帶
  
# 網址問題(http & https):  理論上我們會直接把網址複製給產品 請他們加(看是嵌在遊戲內, 還是跳網頁出來)
     (1) 如果是玩家一開始就連錯連不到, 應該就是產品那邊跳轉錯誤(因為我們給的網址是正確的)
     (2) 如果是玩家填一填跳頁連不到時, 就有可能是我們這邊寫錯了(問卷內的問題歸我們管XDDD)

# 


#靜態server  init.inc.php => getStaticFileDomain