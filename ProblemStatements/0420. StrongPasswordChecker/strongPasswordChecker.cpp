class Solution {
public:
    int strongPasswordChecker(string password) {
        int n = password.size();
        bool low = false, high = false, dig = false;
        for(char c : password){
            if(islower(c)){
                low = true;
            } else if(isupper(c)){
                high = true;
            } else if(isdigit(c)){
                dig = true;
            }
        }
        int a = (!low) + (!high) + (!dig);
        int rep {0};
        vector<int> r;
        for(int i {0}; i < n; ){
            int j = i;
            while(j < n && password[j] == password[i]){
                j++;
            }
            int l = j - i;
            if(l >= 3){
                rep += l / 3;
                r.push_back(l);
            }
            i = j;
        }
        if(n < 6){
            return max(a, 6 - n);
        }
        if(n <= 20){
            return max(a, rep);
        }
        int del = n - 20;
        int rem = del;
        for(int i {0}; i < r.size() && rem > 0; i++){
            if(r[i] >= 3 && r[i] % 3 == 0){
                r[i]--;
                rem--;
                rep--;
            }
        }
        for(int i {0}; i < r.size() && rem > 0; i++){
            if(r[i] >= 3 && r[i]% 3 == 1){
                int temp = min(2, rem);
                r[i] -= temp;
                rem -= temp;
                if(temp == 2){
                    rep--;
                }
            }
        }
        for(int i {0}; i < r.size() && rem > 0; i++){
            if(r[i] >= 3){
                int temp = min(rem, r[i] - 2);
                r[i] -= temp;
                rem -= temp;
                rep -= temp / 3;
            }
        }
        return del + max(a, rep);
    }
};