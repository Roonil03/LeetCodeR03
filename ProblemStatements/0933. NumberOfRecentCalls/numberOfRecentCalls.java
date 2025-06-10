class RecentCounter {
    Queue<Integer> req;

    public RecentCounter() {
        req = new LinkedList<>();    
    }
    
    public int ping(int t) {
        req.offer(t);
        while(!req.isEmpty() && req.peek() < t - 3000){
            req.poll();
        }
        return req.size();
    }
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * RecentCounter obj = new RecentCounter();
 * int param_1 = obj.ping(t);
 */