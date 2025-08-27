class Solution {
public:
    int leastInterval(vector<char>& tasks, int n) {
        vector<int> freq(26, 0);
        for(char t : tasks){
            freq[t-'A']++;
        }
        sort(freq.begin(), freq.end(), greater<int>());
        int m1 = freq[0];
        int m2 {1};
        for(int i {1}; i < 26; i++){
            if(freq[i] == m1){
                m2++;
            } else{
                break;
            }
        }
        int a = tasks.size() - m1 * m2;
        int b = (m1 - 1) * (n - (m2 - 1));
        int temp = max(0, b - a);
        return tasks.size() + temp;
    }
};