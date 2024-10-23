using System;
using System.Collections.Generic;
using System.Linq;
public class Solution
{
    public int LengthOfLongestSubstring(string s)
    {
        char[] chars = s.ToCharArray();
        List<string> words = new List<string>();
        bool newWord = true;

        for (int i = 0; i < chars.Length; i++)
        {
            if (newWord)
            {
                words.Add(chars[i].ToString());
                newWord = false;
            }
            if (i < chars.Length - 1 && !words[words.Count - 1].Contains(chars[i + 1]))
            {
                words[words.Count - 1] = words[words.Count - 1] + chars[i + 1].ToString();
            }
            else
            {
                newWord = true;
                i = words.Count - 1;
            }
        }
        return words.Any() ? words.Select(x => x.Length).Max() : 0;
    }

    public static void Main(string[] args)
    {
        Solution solution = new Solution();
        Console.WriteLine("Enter a string:");
        string input = Console.ReadLine();
        if (string.IsNullOrEmpty(input))
        {
            Console.WriteLine("The input string is empty.");
            return;
        }
        int result = solution.LengthOfLongestSubstring(input);
        Console.WriteLine($"The length of the longest substring without repeating characters is: {result}");
    }
}