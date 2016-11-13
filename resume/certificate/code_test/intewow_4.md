A monkey wants to get to the other side of a river. The monkey is initially located on one bank of the river (position −1) and wants to get to the opposite bank (position N). The monkey can jump any (integer) distance between 1 and D. If D is less than or equal to N then the monkey cannot jump right across the river. Luckily, there are many stones hidden under the water. The water level is constantly decreasing, and soon some of the stones will be out of the water. The monkey can jump to and from the stones, but only when the particular stone is already out of the water.
The stones in the river are described in array A consisting of N integers. A[K] represents a time when the stone at position K will be out of the water (A[K] = −1 means that there is no stone at position K). You can assume that no two stones will surface simultaneously. The goal is to find the earliest time when the monkey can get to the other side of the river.
For example, consider integer D = 3 and the following array A consisting of N = 6 integers:
  A[0] = 1
  A[1] = -1
  A[2] = 0
  A[3] = 2
  A[4] = 3
  A[5] = 5
Initially, the monkey cannot jump across the river in a single jump. However, at time 2, there will be three stones out of the water.

Time 2 is the earliest moment when the monkey can jump across the river (for example, by jumps of length 1, 3 and 3, as marked on the picture above).
Write a function:
function solution(A, D);
that, given a zero-indexed array A consisting of N integers, and integer D, returns the earliest time when the monkey can jump to the other side of the river. If the monkey can leap across the river in just one jump, the function should return 0. If the monkey is never able to jump to the other side of the river, the function should return −1.
For example, given array A and integer D as defined above, the function should return 2 as explained above.
Assume that:
N is an integer within the range [0..100,000];
D is an integer within the range [1..100,001];
each element of array A is an integer within the range [−1..100,000];
no two stones will surface simultaneously.

Complexity:
expected worst-case time complexity is O(N+max(A));
expected worst-case space complexity is O(N+max(A)), beyond input storage (not counting the storage required for input arguments).

Elements of input arrays can be modified.


Copyright 2009–2016 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.



        
      