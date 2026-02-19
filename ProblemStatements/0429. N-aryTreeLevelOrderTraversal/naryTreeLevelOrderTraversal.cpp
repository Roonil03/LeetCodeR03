/*
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val) {
        val = _val;
    }

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};
*/

class Solution {
public:
    vector<vector<int>> levelOrder(Node* root) {
        vector<vector<int>> res;
        if(!root){
            return res;
        }
        queue<Node*> q;
        q.push(root);
        while(!q.empty()){
            int s = q.size();
            vector<int> temp;
            for(int i {0}; i < s; i++){
                Node* n = q.front();
                q.pop();
                temp.push_back(n->val);
                for(Node* c : n->children){
                    q.push(c);
                }
            }
            res.push_back(temp);
        }
        return res;
    }
};