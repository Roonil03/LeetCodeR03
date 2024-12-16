class Solution {
    public boolean equalFrequency(String word) {
        int[] freq = new int[26];
        for (char ch : word.toCharArray()) {
            freq[ch - 'a']++;
        }
        for (int i = 0; i < 26; i++) {
            if (freq[i] == 0)
            {
                continue;
            }
            freq[i]--;
            if (isValid(freq))
            {
                return true;
            }
            freq[i]++;
        }
        return false;
    }

    boolean isValid(int[] freq) {
        int uniqueFreq = 0;
        for (int count : freq) {
            if (count == 0)
            {
                continue;
            }
            if (uniqueFreq == 0)
            {
                uniqueFreq = count;
            }
            else if (uniqueFreq != count)
            {
                return false;
            }
        }
        return true;
    }
}
