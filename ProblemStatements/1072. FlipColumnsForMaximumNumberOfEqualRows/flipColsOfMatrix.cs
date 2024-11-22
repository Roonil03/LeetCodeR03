public class Solution {
    public int MaxEqualRowsAfterFlips(int[][] matrix) {
        var patternCounts = new Dictionary<string, int>();
        foreach (var row in matrix) {
            var pattern = new StringBuilder();
            var invPattern = new StringBuilder();
            foreach (var cell in row) {
                pattern.Append(cell == row[0] ? '0' : '1');
                invPattern.Append(cell == row[0] ? '1' : '0');
            }
            var str = pattern.ToString();
            if (!patternCounts.ContainsKey(str)) {
                patternCounts[str] = 0;
            }
            patternCounts[str]++;
        }
        int rows = 0;
        foreach (var count in patternCounts.Values) {
            rows = Math.Max(rows, count);
        }
        return rows;
    }
}
