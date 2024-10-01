#include <iostream>
#include <vector>

class Solution {
public:
    bool canArrange(vector<int>& arr, int k) {
        vector<int> freq(k, 0);
        
        // Count the frequency of remainders
        for (int num : arr) {
            int rem = ((num % k) + k) % k;
            freq[rem]++;
        }
        
        // Check if the number of elements with remainder 0 is even
        if (freq[0] % 2 != 0) {
            return false; 
        }
        
        // Check pairs of remainders
        for (int i = 1; i <= k / 2; i++) {
            if (freq[i] != freq[k - i]) {
                return false; 
            }
        }
        
        // If k is even, check the middle element
        if (k % 2 == 0) {
            if (freq[k / 2] % 2 != 0) {
                return false;
            }
        }
        
        return true;
    }
};
