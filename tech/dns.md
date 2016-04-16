## dns

A limitation in the way Linux handles DNS queries is that there can only be a maximum of three nameservers used in resolv.conf. As a workaround, you can make localhost the only nameserver in resolv.conf, and then create a separate resolv-file for your external nameservers. First, create a new resolv file for dnsmasq:

default port:53


- /etc/hosts ：這個是最早的 hostname 對應 IP 的檔案；
- /etc/resolv.conf ：這個重要！就是 ISP 的 DNS 伺服器 IP 記錄處；
- /etc/nsswitch.conf：這個檔案則是在『決定』先要使用 /etc/hosts 還是 /etc/resolv.conf 的設定



- 主機名稱.   IN  A           IPv4 的 IP 位址
- 主機名稱.   IN  AAAA        IPv6 的 IP 位址
- 領域名稱.   IN  NS          管理這個領域名稱的伺服器主機名字.
- 領域名稱.   IN  SOA         管理這個領域名稱的七個重要參數(容後說明)
    - Master DNS 伺服器主機名稱：這個領域主要是哪部 DNS 作為 master 的意思。在本例中， dns1.ksu.edu.tw 為 ksu.edu.tw 這個領域的主要 DNS 伺服器囉；
    - 管理員的 email：那麼管理員的 email 為何？發生問題可以聯絡這個管理員。要注意的是， 由於 @ 在資料庫檔案中是有特別意義的，因此這裡就將 abuse@mail.ksu.edu.tw 改寫成 abuse.mail.ksu.edu.tw ，這樣看的懂了嗎？
    - 序號 (Serial)：這個序號代表的是這個資料庫檔案的新舊，序號越大代表越新。 當 slave 要判斷是否主動下載新的資料庫時，就以序號是否比 slave 上的還要新來判斷，若是則下載，若不是則不下載。 所以當你修訂了資料庫內容時，記得要將這個數值放大才行！ 為了方便使用者記憶，通常序號都會使用日期格式『YYYYMMDDNU』來記憶，例如崑山科大的 2010080369 序號代表 2010/08/03 當天的第 69 次更新的感覺。不過，序號不可大於 2 的 32 次方，亦即必須小於 4294967296 才行喔。
    - 更新頻率 (Refresh)：那麼啥時 slave 會去向 master 要求資料更新的判斷？ 就是這個數值定義的。崑山科大的 DNS 設定每 1800 秒進行一次 slave 向 master 要求資料更新。那每次 slave 去更新時， 如果發現序號沒有比較大，那就不會下載資料庫檔案。
    - 失敗重新嘗試時間 (Retry)：如果因為某些因素，導致 slave 無法對 master 達成連線， 那麼在多久的時間內，slave 會嘗試重新連線到 master。在崑山科大的設定中，900 秒會重新嘗試一次。意思是說，每 1800 秒 slave 會主動向 master 連線，但如果該次連線沒有成功，那接下來嘗試連線的時間會變成 900 秒。若後來有成功，則又會恢復到 1800 秒才再一次連線。
    - 失效時間 (Expire)：如果一直失敗嘗試時間，持續連線到達這個設定值時限， 那麼 slave 將不再繼續嘗試連線，並且嘗試刪除這份下載的 zone file 資訊。崑山科大設定為 604800 秒。意思是說，當連線一直失敗，每 900 秒嘗試到達 604800 秒後，崑山科大的 slave 將不再更新，只能等待系統管理員的處理。
    - 快取時間 (Minumum TTL)：如果這個資料庫 zone file 中，每筆 RR 記錄都沒有寫到 TTL 快取時間的話，那麼就以這個 SOA 的設定值為主。
- 領域名稱.   IN  MX          順序數字  接收郵件的伺服器主機名字
- 主機別名.   IN  CNAME       實際代表這個主機別名的主機名字.


### dns catagory
- cache only , cache dns from lookup domain name before
- forwarding , forward to specific DNS


### bind9
[bind9 config syntax](http://www.zytrax.com/books/dns/ch7/view.html)
log : /var/log/message


### dnsmasq

/etc/dnsmasq.conf
```
    ...
    resolv-file=/etc/resolv.dnsmasq.conf
    ...
```

/etc/resolv.dnsmasq.conf
```
    nameserver 8.8.8.8
    nameserver 8.8.4.4
```

### tool
dnsutils
