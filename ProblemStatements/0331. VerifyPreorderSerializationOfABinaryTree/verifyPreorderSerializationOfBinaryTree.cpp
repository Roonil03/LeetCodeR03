class Solution {
public:
    bool isValidSerialization(string preorder) {
        vector<string> st;
        stringstream ss(preorder);
        string node;
        while(getline(ss, node, ',')){
            st.push_back(node);
            while(st.size() >= 3 && st[st.size() - 1] == "#" && st[st.size() - 2] == "#" && st[st.size() - 3] != "#"){
                st.pop_back();
                st.pop_back();
                st.pop_back();
                st.push_back("#");
            }
        }
        return st.size() == 1 && st[0] == "#";
    }
};