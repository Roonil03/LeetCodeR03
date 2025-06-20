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
    TreeNode* deleteNode(TreeNode* root, int key) {
        if(!root){
            return nullptr;
        }
        if(key < root->val){
            root->left = deleteNode(root->left, key);
        } else if(key > root->val){
            root->right = deleteNode(root->right, key);
        } else{
            if(!root->left){
                TreeNode* right = root->right;
                delete root;
                return right;
            } else if(!root->right){
                TreeNode* left = root->left;
                delete root;
                return left;
            } else{
                TreeNode* m = findMin(root->right);
                root->val = m->val;
                root->right = deleteNode(root->right, m->val);
            }
        }
        return root;
    }

    TreeNode* findMin(TreeNode* m){
        while(m && m->left){
            m = m->left;
        }
        return m;
    }
};