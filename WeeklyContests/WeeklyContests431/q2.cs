//  Could not solve it entirely during the time of the contest, 
//  it has been solved post the time limit of the contest.

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

// The provided code defines a Solution class with a method CalculateScore that computes a score 
// based on the positions of characters in a given string s. The method initializes a list of 26 
// stacks, one for each letter of the alphabet. As it iterates through the string, it calculates 
// the mirror character for each character (the character at the same distance from the end of 
// the alphabet). If the stack corresponding to the mirror character is not empty, it pops from 
// the stack and adds the difference between the current index and the popped index to the score. 
// If the stack is empty, it pushes the current index onto the stack for the current character. 
// The time complexity is (O(n)), where (n) is the length of the string, and the space complexity 
// is (O(n)) due to the stacks used to store indices.