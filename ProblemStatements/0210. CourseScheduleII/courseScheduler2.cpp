// class Solution {
// public:
//     vector<int> findOrder(int numCourses, vector<vector<int>>& prerequisites) {
//         vector<int> deg(numCourses, 0);
//         vector<vector<int>> g(numCourses);
//         for (auto& pre : prerequisites){
//             g[pre[1]].push_back(pre[0]);
//             deg[pre[0]]++;
//         }
//         queue<int> q;
//         for (int i = 0; i < numCourses; i++) {
//             if (deg[i] == 0) {
//                 q.push(i);
//             }
//         }
//         vector<int> order;
//         while (!q.empty()) {
//             int c = q.front();
//             q.pop();
//             order.push_back(c);
//             for (int n : g[c]) {
//                 deg[n]--;
//                 if (deg[n] == 0) {
//                     q.push(n);
//                 }
//             }
//         }
//         if (order.size() == numCourses) {
//             return order;
//         } else {
//             return {};
//         }
//     }
// };
class Solution {
    bool ans;
    private:
        void dfs(int node,vector<int> &vis,vector<int>  &visPath,stack<int> &st,vector<vector<int>> &adj)
        {
            
            if(ans)
            {
                vis[node]=1;
                visPath[node]=1;
    
                for(auto it: adj[node])
                {
                    if(!vis[it])
                    {
                        dfs(it,vis,visPath,st,adj);
                    }
                    else
                    {
                        if(visPath[it])
                        {
                            ans=false;
                            break;
                        }
    
                    }
    
    
                }
    
                st.push(node);
                visPath[node]=0;
    
            }
    
    
        }
    
    public:
        vector<int> findOrder(int numCourses, vector<vector<int>>& prerequisites) {
    
            vector<vector<int>> adj(numCourses);
            vector<int> vis(numCourses,0);
            vector<int> visPath(numCourses,0);
            ans=true;
            stack<int> st;
            for(int i=0;i<prerequisites.size();i++)
            {
                int ai=prerequisites[i][0];
                int bi=prerequisites[i][1];
    
                adj[bi].push_back(ai);
            }
    
            for(int i=0;i<numCourses;i++)
            {
                if(!vis[i]) {
                    dfs(i,vis,visPath,st,adj);
                    if(!ans) 
                    {
                        
                        break;
                    }
                }
    
            }
    
        vector<int> a;
    
        while(!st.empty() && ans)
        {
    
            a.push_back(st.top());
            st.pop();
        }
    
        return a;
        
        }
    };