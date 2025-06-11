using ll = long long;

class Solution {
public:
    bool containsNearbyAlmostDuplicate(vector<int>& nums, int indexDiff, int valueDiff) {
        // set<ll> win;
        // for(int i{0}; i < nums.size();i++){
        //     if(i > indexDiff){
        //         win.erase(nums[i - indexDiff - 1]);
        //     }
        //     auto it = win.lower_bound((ll)nums[i] - valueDiff);
        //     if(it != win.end() && *it <= (ll)nums[i] + valueDiff){
        //         return true;
        //     }
        //     win.insert(nums[i]);
        // }
        // return true;
        int k = indexDiff, t = valueDiff;
        if (nums.size() < 2 || k == 0){
            return false;
        }
        deque<int> windows_deq;
        multiset<long> windows;
        for (int i = 0; i < nums.size(); i++) {
            if (windows.size() > k) {
                int num = windows_deq.front();
                windows_deq.pop_front();
                windows.erase(windows.find(num));
            }
            auto it = windows.lower_bound((long)nums[i] - (long)t);
            if (it == windows.end() || *it > (long)nums[i] + (long)t) {
                windows_deq.push_back(nums[i]);
                windows.insert(nums[i]);
            }
            else{
                return true;
            }
        }
        return false;
    }
};