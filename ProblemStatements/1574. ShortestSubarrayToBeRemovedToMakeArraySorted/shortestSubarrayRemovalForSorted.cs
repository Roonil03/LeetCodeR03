public class Solution {
    public int FindLengthOfShortestSubarray(int[] arr) {
        int n = arr.Length;
        int left = 0;
        while (left + 1 < n && arr[left] <= arr[left + 1]) {
            left++;
        }
        if (left == n - 1) {
            return 0;
        }
        int right = n - 1;
        while (right > left && arr[right - 1] <= arr[right]) {
            right--;
        }
        int remove = Math.Min(n - left - 1, right);
        for (int i = 0, j = right; i <= left && j < n; ) {
            if (arr[i] <= arr[j]) {
                remove = Math.Min(remove, j - i - 1);
                i++;
            } else {
                j++;
            }
        }        
        return remove;
    }
}
