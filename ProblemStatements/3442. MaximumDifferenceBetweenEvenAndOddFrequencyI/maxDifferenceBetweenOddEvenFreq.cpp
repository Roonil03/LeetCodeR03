class Solution {
public:
    int maxDifference(string s) {
        vector<int> freq(26, 0);
         for(char c : s) {
            freq[c - 'a']++;
        }
        int a {INT_MIN}, b = INT_MAX;
        for(int n : freq){
            if (n % 2 == 1){
                a = max(a, n);
            } else if (n != 0){
                b = min(b, n);
            }
        }
        //cout << a << b;
        return a - b;
    }
};