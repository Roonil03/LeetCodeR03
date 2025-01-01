public class Solution {
    public int MaxScore(string s) {
        int t1 = 0;
        foreach(char c in s){
            if (c == '1'){
                t1++;
            }
        }
        int res = int.MinValue;
        int c1 = 0, c0 = 0;
        for(int i = 0; i< s.Length - 1; i++){
            if(s[i] == '0'){
                c0++;
            }
            else{
                c1++;
            }
            int ls = c0;
            int rs = t1-c1;
            int ts = ls + rs;
            if(ts > res){
                res = ts;
            }
        }
        return res;
    }
}