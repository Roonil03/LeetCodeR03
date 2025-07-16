/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    int rob(TreeNode* root) {
        auto res = dfs(root);
        return max(res.first, res.second);
    }

    pair<int, int> dfs(TreeNode* n){
        if(!n){
            return {0,0};
        }
        auto l = dfs(n->left);
        auto r = dfs(n->right);
        int a = max(l.first, l.second) + max(r.first, r.second);
        int b = n->val + l.first + r.first;
        return {a, b};
    }
};