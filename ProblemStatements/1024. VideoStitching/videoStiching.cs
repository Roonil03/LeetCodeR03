public class Solution {
    public int VideoStitching(int[][] clips, int time) {
        Array.Sort(clips, (a, b) => a[0].CompareTo(b[0]));
        int end = 0, farthest = 0, count = 0, i = 0;        
        while (i < clips.Length && end < time) {
            while (i < clips.Length && clips[i][0] <= end) {
                farthest = Math.Max(farthest, clips[i][1]);
                i++;
            }
            if (end == farthest) {
                return -1;
            }
            end = farthest;
            count++;
        }        
        return end >= time ? count : -1;
    }
}