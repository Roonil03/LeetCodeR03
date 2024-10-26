#include <stdio.h>
#include <stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode* left;
    struct TreeNode* right;
};

int id = 0;
int *mp1, *mp2, *sz;

int dfs(struct TreeNode* root, int h) {
    if (!root) return 0;
    mp1[root->val] = id;
    mp2[id] = h;
    id++;
    int lz = dfs(root->left, h + 1);
    int rz = dfs(root->right, h + 1);
    sz[mp1[root->val]] = 1 + lz + rz;
    return 1 + lz + rz;
}

int* treeQueries(struct TreeNode* root, int* queries, int queriesSize, int* returnSize) {
    *returnSize = queriesSize;
    int* returnArr = (int*)malloc(sizeof(int) * queriesSize);
    int maxNodes = 100001;
    mp1 = (int*)calloc(maxNodes, sizeof(int));
    mp2 = (int*)calloc(maxNodes, sizeof(int));
    sz = (int*)calloc(maxNodes, sizeof(int));
    
    dfs(root, 0);
    
    int* left = (int*)malloc(id * sizeof(int));
    int* right = (int*)malloc(id * sizeof(int));
    
    for (int i = 0; i < id; i++) {
        left[i] = mp2[i];
        if (i) left[i] = fmax(left[i - 1], left[i]);
    }
    
    for (int i = id - 1; i >= 0; i--) {
        right[i] = mp2[i];
        if (i + 1 < id) right[i] = fmax(right[i], right[i + 1]);
    }
    
    for (int i = 0; i < queriesSize; i++) {
        int nodeid = mp1[queries[i]];
        int l = nodeid, r = l + sz[nodeid] - 1;
        int h = 0;
        if (l > 0) h = fmax(h, left[l - 1]);
        if (r + 1 < id) h = fmax(h, right[r + 1]);
        returnArr[i] = h;
    }
    
    free(mp1);
    free(mp2);
    free(sz);
    free(left);
    free(right);
    return returnArr;
}

void freeTree(struct TreeNode* root) {
    if (root) {
        freeTree(root->left);
        freeTree(root->right);
        free(root);
    }
}