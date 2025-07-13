class Solution {
public:
    vector<int> countSmaller(vector<int>& nums) {
        int n = nums.size();
        vector<int> res(n, 0);
        vector<pair<int, int>> id;
        for (int i {0}; i < n; i++){
            id.push_back({nums[i], i});
        }
        mergeSort(id, res, 0, n - 1);
        return res;
    }

    void mergeSort(vector<pair<int, int>>& id, vector<int>& res, int l, int r){
        if (l >= r){
            return;
        }
        int m = l + (r - l) / 2;
        mergeSort(id, res, l, m);
        mergeSort(id, res, m + 1, r);
        merge(id, res, l, m, r);
    }

    void merge(vector<pair<int, int>>& id, vector<int>& res, int l, int m, int r){
        vector<pair<int, int>> temp(r - l + 1);
        int i = l, j = m + 1, k = 0;
        while(i <= m && j <= r){
            if (id[i].first <= id[j].first){
                res[id[i].second] += (j - m - 1);
                temp[k++] = id[i++];
            } else{
                temp[k++] = id[j++];
            }
        }
        while(i <= m){
            res[id[i].second] += (j - m - 1);
            temp[k++] = id[i++];
        }
        while(j <= r){
            temp[k++] = id[j++];
        }
        for (i = l ; i <= r; i++){
            id[i] = temp[i - l];
        }
    }
};