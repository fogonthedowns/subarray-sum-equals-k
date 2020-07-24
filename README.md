# Subarray sum equals k

Given an array of integers and an integer k, you need to find the total number of continuous subarrays whose sum equals to k.

## Use a map

This strategy has a time and space complexity of O(n). 

iterate through list, track the sum. check if the map contains (sum -k), if so increment the counter. Finally,  Map the (sum - k) to the number of times it happens. This is done, because of negative numbers, an sum can happen more than once.
