class WordDictionary {
    struct TrieNode{
        struct TrieNode *child[26];
        bool isEnd;
    };
    TrieNode* node;

public:
    WordDictionary() {
        node = new TrieNode();
    }
    
    void addWord(string word) {
        TrieNode* n = node;
        for (char c : word){
            if(!n -> child[c-'a']){
                n->child[c-'a'] = new TrieNode();
            }
            n = n->child[c-'a'];
        }
        n->isEnd = true;
    }
    
    bool search(string word) {
        return search(word.c_str(), node);
    }

private:
    bool search(const char* word, TrieNode* node){
        for(int i = 0; word[i] && node;i++){
            if(word[i] != '.'){
                node = node-> child[word[i] - 'a'];
            } else{
                TrieNode* temp = node;
                for(int j = 0; j < 26; j++){
                    node = temp->child[j];
                    if(search(word+i+1, node)){
                        return true;
                    }
                }
            }
        }
        return node && node->isEnd;
    }
};


/**
 * Your WordDictionary object will be instantiated and called as such:
 * WordDictionary* obj = new WordDictionary();
 * obj->addWord(word);
 * bool param_2 = obj->search(word);
 */