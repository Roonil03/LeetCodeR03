class Solution {
    public:
        int minMutation(string startGene, string endGene, vector<string>& bank) {
            // unordered_set<string> s{bank.begin(), bank.end()};
            // if(!s.count(endGene)){
            //     return -1;
            // }
            // queue<string> q;
            // q.push(startGene);
            // int res {0}, s1;
            // string cur,temp;
            // while(!q.empty()){
            //     cur = q.front();
            //     q.pop();
            //     if(cur == endGene){
            //         return res;
            //     }
            //     s.erase(cur);
            //     for(int i{0}; i < 8; i++){
            //         temp = cur;
            //         temp[i] = 'A';
            //         if(s.count(temp)){
            //             q.push(temp);
            //         }
            //         temp[i] = 'C';
            //         if(s.count(temp)){
            //             q.push(temp);
            //         }
            //         temp[i] = 'G';
            //         if(s.count(temp)){
            //             q.push(temp);
            //         }
            //         temp[i] = 'T';
            //         if(s.count(temp)){
            //             q.push(temp);
            //         }
            //     }
            //     res++;
            // }
            // return -1;
            map<string,int> dist;
            dist[startGene] = 0;
            queue<string> q;
            q.push(startGene);
            while (q.size()) {
                auto u = q.front(); q.pop();
                for (auto v : bank) {
                    if (dist.count(v)){
                        continue;
                    }
                    int cnt = 0;
                    for (int i = 0; i < 8; i++) {
                        if (u[i] != v[i]){
                            cnt++;
                        }
                    }
                    if (cnt == 1){
                        dist[v] = dist[u] + 1;
                        q.push(v);
                    }
                }
            }
            if (dist.count(endGene)){
                return dist[endGene];
            } else{
                return -1;
            }
        }
    };