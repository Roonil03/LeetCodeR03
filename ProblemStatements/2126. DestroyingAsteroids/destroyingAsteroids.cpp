class Solution {
public:
    using ll = long long;

    bool asteroidsDestroyed(int mass, vector<int>& asteroids) {
        sort(asteroids.begin(), asteroids.end());
        ll temp = mass;
        for(int i {0}; i < asteroids.size(); i++){
            if(temp >= asteroids[i]){
                temp += asteroids[i];
            } else{
                return false;
            }
        }
        return true;
    }
};