class Solution {
    //     public boolean canConstruct(String s, int k) {
    //         if(k > s.length())
    //         {
    //             return false;
    //         }
    //         Map<Character, Integer>m = new HashMap<>();
    //         for(char c : s.toCharArray())
    //         {
    //             m.put(c, m.getOrDefault(c,0) + 1);
    //         }
    //         int o = 0;
    //         for(int c : m.values())
    //         {
    //             if(c % 2 == 1)
    //             {
    //                 o++;
    //             }
    //         }
    //         return o<=k;
    //     }
        public boolean canConstruct(String s, int k) {
            int odd = 0, n = s.length(), count[] = new int[26];
            for (int i = 0; i < n; ++i) {
                count[s.charAt(i) - 'a'] ^= 1;
                odd += count[s.charAt(i) - 'a'] > 0 ? 1 : -1;
            }
            return odd <= k && k <= n;
        }
    }
    