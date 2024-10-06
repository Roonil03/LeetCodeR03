class Solution {
    public boolean areSentencesSimilar(String sentence1, String sentence2) {
        String s1[] = sentence1.split(" ");
        String s2[] = sentence2.split(" ");
        if(s1.length<s2.length)
        {
            String[] temp = s1;
            s1 = s2;
            s2 = temp;
        }
        int left = 0, right = 0;
        int len1 = s1.length,len2 = s2.length;
        while(left < len2 && s1[left].equals(s2[left]))
        {
            left++;
        }
        while(right<len2 && s1[len1-right-1].equals(s2[len2-right-1]))
        {
            right++;
        }
        return right + left >= len2;
    }
}