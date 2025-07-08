class NumArray {
    // int[] tree;
    // int n;

    static class BinaryIndexedTree {
    int[] nums;
    int[] BIT;
    int n;

    public BinaryIndexedTree(int[] nums) {
      this.nums = nums;
      this.n = nums.length;
      BIT = new int[n + 1];
      for (int i = 0; i < n; i++) {
        init(i, nums[i]);
      }
    }

    void init(int i, int val) {
      i++;
      while (i <= n) {
        BIT[i] += val;
        i += (i & -i);
      }
    }

    void update(int i, int val) {
      int diff = val - nums[i];
      nums[i] = val;
      init(i, diff);
    }

    int getSum(int i) {
      i++;
      int ret = 0;
      while (i > 0) {
        ret += BIT[i];
        i -= (i & -i);
      }
      return ret;
    }
  }

  BinaryIndexedTree b;

    public NumArray(int[] nums) {
        // n = nums.length;
        // tree = new int[2*n];
        // for(int i = 0; i < n; i++){
        //     tree[n+1] = nums[i];
        // }
        // for(int i = n - 1; i > 0; i--){
        //     tree[i] = tree[2 * i] + tree[2 * i + 1];
        // }
        b = new BinaryIndexedTree(nums);
    }
    
    public void update(int index, int val) {
        // index += n;
        // tree[index] = val;
        // while(index > 1){
        //     index /= 2;
        //     tree[index] = tree[2 * index] + tree[2 * index + 1];
        // }
        b.update(index, val);
    }
    
    public int sumRange(int left, int right) {
    //     int sum = 0;
    //     left += n;
    //     right += n;
    //     while(left <= right){
    //         if ((left & 1) == 1){
    //             sum += tree[left];
    //             left++;
    //         }
    //         if((right & 1) == 1){
    //             sum += tree[right];
    //             right--;
    //         }
    //         left /= 2;
    //         right /= 2;
    //     }
    //     return sum;
    // }
        return b.getSum(right) - b.getSum(left - 1);
    }
}

/**
 * Your NumArray object will be instantiated and called as such:
 * NumArray obj = new NumArray(nums);
 * obj.update(index,val);
 * int param_2 = obj.sumRange(left,right);
 */