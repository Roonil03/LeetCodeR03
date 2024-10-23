class Solution {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int[] merge = new int[nums1.length + nums2.length];
        int i = 0, j= 0 , k =0;
        while(i<nums1.length&&j<nums2.length){
            if(nums1[i]<=nums2[j]){
                merge[k++] = nums1[i++];
            }
            else{
                merge[k++] = nums2[j++];
            }
        }
        while (i < nums1.length) {
            merge[k++] = nums1[i++];
        }
        while(j<nums2.length){
            merge[k++] = nums2[j++];
        }
        // for(int a = 0; a<merge.length;a++){
        //     System.out.print(merge[a] + " ");
        // }
        if(merge.length%2==0){
            return (double)(merge[merge.length/2-1] + merge[merge.length/2])/2;
        }
        else{
            return merge[merge.length/2];
        }
    }
}
public class medianOfTwoSorted
{
    public static void main(String[] args) {
        Solution s1 = new Solution();
        System.out.println(s1.findMedianSortedArrays(new int[] {1,3}, new int[] {2}));
        System.out.println(s1.findMedianSortedArrays(new int[] {1,2}, new int[] {3,4}));
    }
}