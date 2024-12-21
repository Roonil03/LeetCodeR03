public class Solution {
    public int MaximumValue(string[] strs) {
        int res = 0;
        foreach(var s in strs)
        {
            if(int.TryParse(s, out int num))
            {
                res = Math.Max(res, num);
            }
            else{
                res = Math.Max(res, s.Length);
            }
        }
        return res;
    }
}
