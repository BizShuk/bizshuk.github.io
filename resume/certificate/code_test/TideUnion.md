# TideUnion


### 1
How to prevent replay request? 
Ans: count for each request in client and append on header or parameters. And server verify it for specific client.

### 2
train schedule 8am~10am, between A to E with path A->B->C->D->E->D->C->B->A->... , 5 mins for each stop.

write a function with input (departure stop, arrival stop) , get departure time and arrival time

Ans1: produce a list for each day path , search matched stops with available time. 



### 3
Build 10000 concurrent transaction system , what should be aware?

Ans: 
1. connection nodes , 
2. db 
3. cache 
4. optimise db schema , api ,sql query.
5. CDN