class Solution {

    public boolean[] isArraySpecial(int[] nums, int[][] queries) {
        boolean[] res = new boolean[queries.length];
        ArrayList<Integer> viol = new ArrayList<>();
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] % 2 == nums[i - 1] % 2) {
                viol.add(i);
            }
        }
        for (int i = 0; i < queries.length; i++) {
            int start = queries[i][0];
            int end = queries[i][1];

            res[i] = !hasViol(start + 1, end, viol);
        }
        return res;
    }

    private boolean hasViol(int start, int end, ArrayList<Integer> viol) {
        int l = 0, r = viol.size() - 1;
        while (l <= r) {
            int mid = l + (r - l) / 2;
            int v = viol.get(mid);

            if (v < start) {
                l = mid + 1;
            } else if (v > end) {
                r = mid - 1;
            } else {
                return true;
            }
        }
        return false;
    }
}
