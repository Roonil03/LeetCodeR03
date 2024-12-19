public class Solution {
    public int MaxChunksToSorted(int[] arr) {
        int count = 0;
        int k = 0;
        for(int i = 0; i< arr.Length; i++)
        {
            k = Math.Max(k, arr[i]);
            if(k == i)
            {
                count++;
            }
        }
        return count;
    }
}