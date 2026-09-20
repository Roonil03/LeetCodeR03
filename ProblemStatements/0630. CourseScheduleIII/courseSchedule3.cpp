class Solution {
public:
    int scheduleCourse(vector<vector<int>>& courses) {
        ranges::sort(courses, [](auto& a, auto& b){
            return a[1] < b[1];
        });
        priority_queue<int>pq;
        int time {0};
        for(auto& v : courses){
            if(time + v[0] <= v[1]){
                pq.push(v[0]);
                time += v[0];
            } else if(!pq.empty() && pq.top () > v[0]){
                time += v[0] - pq.top();
                pq.pop();
                pq.push(v[0]);
            }
        }
        return pq.size();
    }
};