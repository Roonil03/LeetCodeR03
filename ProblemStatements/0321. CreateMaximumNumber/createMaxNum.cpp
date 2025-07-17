class Solution {
public:
    vector<int> maxNumber(vector<int>& nums1, vector<int>& nums2, int k) {
        int m = nums1.size(), n = nums2.size();
        vector<int> res(k, 0);
        int l = max(0, k - n), r = min(k, m)
        ;
        for (int i = l; i <= r; i++){
            vector<int>s1 = h1(nums1, i);
            vector<int>s2 = h1(nums2, k - i);
            vector<int> can = h2(s1, s2);
            if(can > res){
                res = move(can);
            }
        }
        return res;
    }

    vector<int>h1(vector<int>& nums, int k){
        int a = nums.size() - k;
        vector<int>st;
        for (int x : nums){
            while(!st.empty() && a > 0 && st.back() < x){
                st.pop_back();
                a--;
            }
            st.push_back(x);
        }
        st.resize(k);
        return st;
    }

    vector<int>h2(vector<int>& a, vector<int>& b){
        int i {0}, j {0}, m = a.size(), n = b.size();
        vector<int>res;
        res.reserve(m + n);
        while(i < m || j < n){
            if(h3(a, i, b, j)){
                res.push_back(a[i++]);
            } else{
                res.push_back(b[j++]);
            }
        }
        return res;
    }

    bool h3(vector<int>& a, int i, vector<int>& b, int j){
        int m = a.size(), n = b.size();
        while(i < m && j < n){
            if(a[i] != b[j]){
                return a[i] > b[j];
            }
            i++;
            j++;
        }
        return (m - i) > (n - j);
    }
};