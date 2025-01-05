public class Solution {
    public long CalculateScore(string s) {
        List<Stack<int>> st = new List<Stack<int>>(26);
        for (int i = 0; i < 26; i++)
        {
            st.Add(new Stack<int>());
        }

        long score = 0;
        for (int i = 0; i < s.Length; i++)
        {
            char c = s[i];
            char mirrorChar = (char)('z' - (c - 'a'));
            int mirrorIndex = mirrorChar - 'a';

            if (st[mirrorIndex].Count > 0)
            {
                int j = st[mirrorIndex].Pop();
                score += (i - j);
            }
            else
            {
                st[c - 'a'].Push(i);
            }
        }
        return score;
    }
}