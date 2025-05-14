class Solution {
public:
    string mergeAlternately(string word1, string word2) {
        string res = "";
        int i {0};
        for(; i < word1.length() && i < word2.length(); i++){
            res += word1[i];
            res += word2[i];
        }
        while(i < word1.length()){
            res += word1[i++];
        }
        while(i < word2.length()){
            res += word2[i++];
        }
        return res;
    }
};