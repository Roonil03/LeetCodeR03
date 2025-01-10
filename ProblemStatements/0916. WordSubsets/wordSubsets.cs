public class Solution {
    public IList<string> WordSubsets(string[] words1, string[] words2) {
        List<string> res = new List<string>();
        int [] freq = new int[26];
        foreach(var w in words2)
        {
            int [] wFreq = CountFrequency(w);
            for(int i = 0; i < 26; i++)
            {
                freq[i] = Math.Max(freq[i], wFreq[i]);
            }
        }
        foreach (var word in words1)
        {
            int[] wFreq = CountFrequency(word);
            bool asw = true;
            for (int i = 0; i < 26; i++)
            {
                if (wFreq[i] < freq[i])
                {
                    asw = false;
                    break;
                }
            }
            if (asw)
            {
                res.Add(word);
            }
        }
        return res;
    }
    int[] CountFrequency(string word)
    {
        int[] freq = new int[26];
        foreach (var ch in word)
        {
            freq[ch - 'a']++;
        }
        return freq;
    }
}