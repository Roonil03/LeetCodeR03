public class Solution {
    public bool MakeStringsEqual(string s, string target) {
        // int ones = 0, onesT = 0;
        // for(int i = 0; i < s.Length; i++)
        // {
        //     if(s[i] == '1')
        //     {
        //         ones++;
        //     }
        //     if(target[i] == '1')
        //     {
        //         onesT++;
        //     }
        // }
        // return ones==onesT;
        return s.Contains("1") == target.Contains("1");
    }
}