class Solution {
    public int[] earliestAndLatest(int n, int firstPlayer, int secondPlayer) {
    //     int[] res = new int[]{Integer.MAX_VALUE, Integer.MIN_VALUE};
    //     simulateRounds(n, firstPlayer, secondPlayer, 1, res);
    //     return res;
    // }
    // void simulateRounds(int n, int fp, int sp, int cur, int[] res)
    // {
    //     if(fp +sp == n+1)
    //     {
    //         res[0] = Math.min(res[0], cur);
    //         res[1] = Math.max(res[1], cur);
    //         return;
    //     }
    //     int n1 = (fp + 1) / 2;
    //     int n2 = (sp + 1) / 2;
    //     int n3 = (n+1) / 2;
    //     simulateRounds(n3, n2, n1, cur+1, res);
        if(firstPlayer + secondPlayer == n + 1)
        {
            return new int[]{1,1};
        }
        if(n == 3 || n == 4)
        {
            return new int[]{2,2};
        }
        if(firstPlayer - 1 > n - secondPlayer)
        {
            int t = n + 1 - firstPlayer;
            firstPlayer = n + 1 - secondPlayer;
            secondPlayer = t;
        }
        int m = (n + 1) /2 ;
        int min = n;
        int max = 1;
        if(secondPlayer * 2 <= n + 1)
        {
            int a = firstPlayer - 1;
            int b = secondPlayer - firstPlayer - 1;
            for(int i = 0; i <= a; i++)
            {
                for(int j = 0; j <= b; j++)
                {
                    int[] ret = earliestAndLatest(m, i+1, i+j+2);
                    min = Math.min(min, 1 + ret[0]);
                    max = Math.max(max, 1 + ret[1]);
                }
            }
        }
        else{
            int p4 = n + 1 - secondPlayer;
            int a = firstPlayer - 1;
            int b = p4 - firstPlayer - 1;
            int c = secondPlayer - p4 - 1;
            for (int i = 0; i <= a; i++) {
                for (int j = 0; j <= b; j++) {
                    int[] ret = earliestAndLatest(m, i + 1, i + j + 1 + (c + 1) / 2 + 1);
                    min = Math.min(min, 1 + ret[0]);
                    max = Math.max(max, 1 + ret[1]);
                }
            }
        }
        return new int[]{min, max};
    }
}