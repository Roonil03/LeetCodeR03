class Solution {
    public String compressedString(String word) {
        StringBuilder comp = new StringBuilder();
        int count = 0;
        for(int i = 0; i<word.length();i++)
        {
            char c = word.charAt(i);
            while(i < word.length() && word.charAt(i)==c && count <9)
            {
                i++;
                count++;
            }
            i--;
            comp.append(count).append(c);
            count = 0;
        }
        return comp.toString();
    }
}