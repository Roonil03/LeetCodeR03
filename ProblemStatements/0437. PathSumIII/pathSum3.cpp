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
    //using ll = long long;
    int pathSum(TreeNode* root, int targetSum) {
        unordered_map<long long, int> p;
        p[0] = 1;
        return dfs(root, 0, targetSum, p);
    }

    int dfs(TreeNode* root, long long sum, int target, unordered_map<long long, int> p){
        if(!root){
            return 0;
        }
        sum += root->val;
        int res = p[sum - target];
        p[sum]++;
        res += dfs(root->left, sum, target, p);
        res += dfs(root->right, sum, target, p);
        p[sum]--;
        return res;
    }
};