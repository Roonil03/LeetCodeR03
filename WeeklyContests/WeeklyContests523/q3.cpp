class BusBooking {
public:
    int tot {0}, count {0};
    vector<uint32_t> tt;
    
    BusBooking(int n)  : tt(n, 0){
        // tot, count, tt
        // int tot {0}, count {0};
        // vector<uint32_t> tt;
    }
    
    void toggle(string seat) {
        int r {0}, c = seat[0] - 'A', sh = c * 8;
        for(int i {1}; i < seat.size(); i++){
            r = r * 10 + (seat[i] - '0');
        }
        uint32_t &s = tt[r - 1];
        if(s & (1 << sh)){
            tot -= (s >> (sh + 4)) & 15;
            count--;
            s &= ~(255 << sh);
        } else{
            int cost = 1 + 2 * (((c == 0) & ((s >> 8) & 1)) | ((c == 3) & ((s >> 16) & 1)));
            tot += cost;
            count++;
            s |= (1 << sh) | (cost << (sh + 4));
        }
    }
    
    int getTotalTime() {
        return tot;
    }
    
    int getMinTime() {
        return count;
    }
};

/**
 * Your BusBooking object will be instantiated and called as such:
 * BusBooking* obj = new BusBooking(n);
 * obj->toggle(seat);
 * int param_2 = obj->getTotalTime();
 * int param_3 = obj->getMinTime();
 */