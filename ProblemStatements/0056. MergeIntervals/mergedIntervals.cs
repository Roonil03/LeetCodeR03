public class Solution {
    public int[][] Merge(int[][] intervals) {
        Array.Sort(intervals, (a,b) => a[0].CompareTo(b[0]));
        List<int[]> mer = new List<int[]>();
        foreach(var i in intervals)
        {
            if (mer.Count == 0 || mer[^1][1] < i[0])
            {
                mer.Add(i);
            }
            else {
                mer[^1][1] = Math.Max(mer[^1][1], i[1]);
            }
        }
        return mer.ToArray();
    }
}