public class Solution {
    public IList<IList<int>> Subsets(int[] nums) {
        IList<IList<int>> res = new List<IList<int>>();
        List<int>sub = new List<int>();
        bt(nums, 0, sub, res);
        return res;
    }
    void bt(int[] nums, int i, List<int> sub, IList<IList<int>> res)
    {
        res.Add(new List<int>(sub));
        for(int a = i; a < nums.Length; a++)
        {
            sub.Add(nums[a]);
            bt(nums, a + 1, sub, res);
            sub.RemoveAt(sub.Count - 1);
        }
    }
}