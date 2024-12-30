class Trie {
    Trie[] ch = new Trie[26];
    int min_sz = Integer.MAX_VALUE;
    int best_i = 0;
    
    void insert(String w, int i, int j) {
        if (min_sz > w.length()) {
            best_i = i;
            min_sz = w.length();
        }
        if (j < 0) {
            return;
        }
        if (ch[w.charAt(j) - 'a'] == null) {
            ch[w.charAt(j) - 'a'] = new Trie();
        }
        ch[w.charAt(j) - 'a'].insert(w, i, j - 1);
    }
    
    int find(String q, int j) {
        if (j >= 0 && ch[q.charAt(j) - 'a'] != null) {
            return ch[q.charAt(j) - 'a'].find(q, j - 1);
        }
        return best_i;
    }
}

class Solution {
    public int[] stringIndices(String[] wordsContainer, String[] wordsQuery) {
        Trie t = new Trie();
        for (int i = 0; i < wordsContainer.length; ++i) {
            t.insert(wordsContainer[i], i, wordsContainer[i].length() - 1);
        }
        int[] res = new int[wordsQuery.length];
        for (int i = 0; i < wordsQuery.length; ++i) {
            res[i] = t.find(wordsQuery[i], wordsQuery[i].length() - 1);
        }
        return res;
    }
}
