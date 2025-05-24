class TopVotedCandidate {
    int[] times;
    int[] leaders;
    
    public TopVotedCandidate(int[] persons, int[] times) {
        this.times = times;
        leaders = new int[times.length];
        Map<Integer, Integer> c = new HashMap<>();
        int lead = -1, r = 0;
        for (int i = 0; i < persons.length; i++) {
            int p = persons[i];
            c.put(p, c.getOrDefault(p, 0) + 1);
            if (c.get(p) >= r){
                lead = p;
                r = c.get(p);
            }
            leaders[i] = lead;
        }
    }
    
    public int q(int t) {
        int l = 0, r = times.length - 1, ans = 0;
        while (l <= r) {
            int m = l + (r - l) / 2;
            if (times[m] <= t) {
                ans = m;
                l = m + 1;
            } else {
                r = m - 1;
            }
        }
        return leaders[ans];
    }
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * TopVotedCandidate obj = new TopVotedCandidate(persons, times);
 * int param_1 = obj.q(t);
 */