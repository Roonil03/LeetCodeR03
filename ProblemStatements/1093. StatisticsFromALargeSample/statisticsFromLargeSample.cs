public class Solution {
    public double[] SampleStats(int[] count) {
        double min = -1;
        double max = 0;
        double sum = 0;
        long total = 0;
        int mode = 0;
        int maxF = 0;        
        for (int i = 0; i < count.Length; i++) {
            if (count[i] > 0) {
                if (min == -1) {
                    min = i;
                }
                max = i;                
                sum += (double)i * count[i];
                total += count[i];                
                if (count[i] > maxF) {
                    maxF = count[i];
                    mode = i;
                }
            }
        }        
        double mean = sum / total;        
        double med = 0;
        if (total % 2 == 0) {
            long t1 = total / 2;
            long t2 = t1 + 1;
            long cc = 0;
            int m1 = 0, m2 = 0;
            
            for (int i = 0; i < count.Length; i++) {
                cc += count[i];
                if (m1 == 0 && cc >= t1) {
                    m1 = i;
                }
                if (cc >= t2) {
                    m2 = i;
                    break;
                }
            }
            med = (m1 + m2) / 2.0;
        }
        else {
            long aim = (total + 1) / 2;
            long cc = 0;            
            for (int i = 0; i < count.Length; i++) {
                cc += count[i];
                if (cc >= aim) {
                    med = i;
                    break;
                }
            }
        }        
        return new double[] { min, max, mean, med, mode };
    }
}