public class Solution {
    public double MinAreaFreeRect(int[][] pts) {
        int n = pts.Length;
        double res = double.MaxValue;
        if (n < 4) return 0.0;
        var map = new Dictionary<string, List<int[]>>();
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                long d = (long)(pts[i][0] - pts[j][0]) * (pts[i][0] - pts[j][0]) + 
                         (long)(pts[i][1] - pts[j][1]) * (pts[i][1] - pts[j][1]);
                double cx = (double)(pts[j][0] + pts[i][0]) / 2;
                double cy = (double)(pts[j][1] + pts[i][1]) / 2;
                string key = d + "+" + cx + "+" + cy;
                if (!map.ContainsKey(key)) {
                    map[key] = new List<int[]>();
                }
                map[key].Add(new int[]{i, j});
            }
        }

        foreach (var key in map.Keys) {
            if (map[key].Count > 1) {
                var list = map[key];
                for (int i = 0; i < list.Count; i++) {
                    for (int j = i + 1; j < list.Count; j++) {
                        int p1 = list[i][0];
                        int p2 = list[j][0];
                        int p3 = list[j][1];
                        double l1 = Math.Sqrt(Math.Pow(pts[p1][0] - pts[p2][0], 2) + Math.Pow(pts[p1][1] - pts[p2][1], 2)); 
                        double l2 = Math.Sqrt(Math.Pow(pts[p1][0] - pts[p3][0], 2) + Math.Pow(pts[p1][1] - pts[p3][1], 2));
                        double area = l1 * l2;
                        res = Math.Min(res, area);
                    }
                }
            }
        }
        return res == double.MaxValue ? 0.0 : res;
    }
}
