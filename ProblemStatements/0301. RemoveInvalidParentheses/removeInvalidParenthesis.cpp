class Solution {
public:
    vector<string> removeInvalidParentheses(string s) {
        vector<string> res;
        unordered_set<string>vis;
        queue<string>q;
        bool f = false;
        q.push(s);
        vis.insert(s);
        while(!q.empty()){
            int s = q.size();
            for(int i {0}; i < s; i++){
                string c = q.front();
                q.pop();
                if(isValid(c)){
                    res.push_back(c);
                    f = true;
                }
                if (f){
                    continue;
                }
                for(int j {0}; j < c.length(); j++){
                    if(c[j] != '(' && c[j] != ')'){
                        continue;
                    }
                    string next = c.substr(0, j) + c.substr(j + 1);
                    if(!vis.count(next)){
                        vis.insert(next);
                        q.push(next);
                    }
                }
            }
            if(f){
                break;
            }
        }
        return res;
    }

    bool isValid(string& s){
        int c {0};
        for(char ch : s){
            if(ch == '('){
                c++;
            }
            if(ch == ')'){
                c--;
                if(c < 0){
                    return false;
                }
            }
        }
        return c == 0;
    }
};