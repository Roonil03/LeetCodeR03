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
    using ll = long long;
    ll mod = 1e9 + 7;
    ll tot = 0;
    ll res = 0;

    ll h1(TreeNode* n){
        if(!n){
            return 0;
        }
        return n->val + h1(n->left) + h1(n->right);
    }

    ll dfs(TreeNode* n){
        if(!n){
            return 0;
        }
        ll r = dfs(n->right);
        ll l = dfs(n->left);
        ll cur = n->val + l + r;
        res = max(res, cur * (tot - cur));
        return cur;
    }

    int maxProduct(TreeNode* root) {
        tot = h1(root);
        dfs(root);
        return (int)(res % mod);
    }
};