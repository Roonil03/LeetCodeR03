public class Solution {
    int[] pre;
    int tot;
    Random rand;

    public Solution(int[] w) {
        pre = new int[w.length];
        pre[0] = w[0];
        for(int i = 1;i < w.length; i++){
            pre[i] = pre[i - 1] + w[i];
        }
        tot = pre[w.length - 1];
        rand = new Random();
    }
    
    public int pickIndex() {
        int tar = rand.nextInt(tot) + 1;
        int l = 0, h = pre.length - 1;
        while(l < h){
            int m = l + (h - l) / 2;
            if(pre[m] < tar){
                l = m + 1;
            } else{
                h = m;
            }
        }
        return l;
    }
}

/**
 * Your Solution object will be instantiated and called as such:
 * Solution obj = new Solution(w);
 * int param_1 = obj.pickIndex();
 */ {
    
}
