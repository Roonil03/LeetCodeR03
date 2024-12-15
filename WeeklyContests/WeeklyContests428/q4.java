/* Could not solve it entirely during the time of the contest, 
 * it has been solved post the time limit of the contest.
 */

import java.util.*;

class Solution {
    public int makeStringGood(String s) {
        Map<Character, Integer> cntr = new HashMap<>();
        for (char ch : s.toCharArray()) 
        {
            cntr.put(ch, cntr.getOrDefault(ch, 0) + 1);
        }        
        if (cntr.size() == 1) 
        {
            return 0;
        }        
        int maxFreq = Collections.max(cntr.values());
        int result = Integer.MAX_VALUE;
        for (int freq = 1; freq <= maxFreq; freq++) 
        {
            result = Math.min(result, checkFreq(freq, cntr));
        }        
        return result;
    }

    private int checkFreq(int freq, Map<Character, Integer> cntr) {
        Map<String, Integer> memo = new HashMap<>();
        return dp(97, 0, freq, cntr, memo);
    }

    private int dp(int ch, int carry, int freq, Map<Character, Integer> cntr, Map<String, Integer> memo) 
    {
        if (ch > 97 + 26) 
        {
            return carry;
        }        
        String key = ch + "," + carry;
        if (memo.containsKey(key)) 
        {
            return memo.get(key);
        }        
        char c = (char) ch;
        int res = 0;
        if (!cntr.containsKey(c)) 
        {
            res = dp(ch + 1, 0, freq, cntr, memo) + carry;
        } 
        else{
            int count = cntr.get(c);
            if (count >= freq) 
            {
                res = dp(ch + 1, count - freq, freq, cntr, memo) + carry;
            } 
            else{
                res = dp(ch + 1, count, freq, cntr, memo) + carry;
                res = Math.min(res, dp(ch + 1, 0, freq, cntr, memo) + Math.max(carry, freq - count));
            }
        }        
        memo.put(key, res);
        return res;
    }
}

/*
The code aims to find the minimum number of deletions required to make a given string "good," where a "good" string 
is defined as having each character occurring at least a certain frequency. It first counts the frequency of each character 
in the string. Then, it uses a dynamic programming approach to explore different combinations of character frequencies and 
calculate the minimum deletions needed to achieve each target frequency. By iterating through possible target frequencies, 
the code determines the overall minimum number of deletions necessary to make the string "good."
 */