class Solution {
public:
    int countRangeSum(vector<int>& nums, int lower, int upper) {
        int n = nums.size();
        vector<long long> prefix(n + 1, 0);
        for (int i {0}; i < n; i++){
            prefix[i+1] = prefix[i] + nums[i];
        }
        return h1(prefix, 0, prefix.size(), lower, upper);
    }

    int h1(vector<long long>& prefix, int left, int right, int lower, int upper) {
        if (right - left <= 1){
            return 0;
        }
        int mid = (left + right) / 2;
        int count = h1(prefix, left, mid, lower, upper) + h1(prefix, mid, right, lower, upper);
        int l = mid, r = mid, j = mid, t {0};
        vector<long long> temp(right - left, 0);
        for (int i = left, k {0}; i < mid; i++, k++) {
            while (l < right && prefix[l] - prefix[i] < lower){
                l++;
            }
            while (r < right && prefix[r] - prefix[i] <= upper){
                r++;
            }
            count += r - l;
            while (j < right && prefix[j] < prefix[i]){
                temp[t++] = prefix[j++];
            }
            temp[t++] = prefix[i];
        }
        for (int i {0}; i < j - left; i++)
            prefix[left + i] = temp[i];
        for (; j < right; j++)
            prefix[left + (j-left)] = prefix[j];
        return count;
    }
};