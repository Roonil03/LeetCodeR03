class FindSumPairs {

    int[] nums1, nums2;
    Map<Integer, Integer>freq;

    public FindSumPairs(int[] nums1, int[] nums2) {
        this.nums1 = nums1;
        this.nums2 = nums2;
        freq = new HashMap<>();
        for(int n : nums2){
            freq.put(n, freq.getOrDefault(n, 0) + 1);
        }
    }
    
    public void add(int index, int val) {
        int a = nums2[index];
        freq.put(a, freq.get(a) - 1);
        if(freq.get(a) == 0){
            freq.remove(a);
        }
        nums2[index] += val;
        freq.put(nums2[index], freq.getOrDefault(nums2[index], 0) + 1);
    }
    
    public int count(int tot) {
        int res = 0;
        for(int n : nums1){
            int tar = tot - n;
            res += freq.getOrDefault(tar, 0);
        }
        return res;
    }
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * FindSumPairs obj = new FindSumPairs(nums1, nums2);
 * obj.add(index,val);
 * int param_2 = obj.count(tot);
 */