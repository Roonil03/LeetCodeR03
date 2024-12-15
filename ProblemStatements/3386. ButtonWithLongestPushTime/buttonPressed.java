class Solution {
    public int buttonWithLongestTime(int[][] events) {
        int maxi = 0, ind = Integer.MAX_VALUE, prevTime = 0;
        for (int[] event : events) {
            int duration = event[1] - prevTime;
            if (duration > maxi || (duration == maxi && event[0] < ind)) {
                ind = event[0];
                maxi = duration;
            }
            prevTime = event[1];
        }
        return ind;
    }
}