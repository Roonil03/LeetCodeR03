class DetectSquares {

    Map<Integer, Map<Integer, Integer>> p;

    public DetectSquares() {
        p = new HashMap<>();    
    }
    
    public void add(int[] point) {
        int x = point[0];
        int y = point[1];
        p.computeIfAbsent(x, k -> new HashMap<>()).merge(y, 1, Integer::sum);
    }
    
    public int count(int[] point) {
        int x1 = point[0];
        int y1 = point[1];
        if(!p.containsKey(x1)){
            return 0;
        }
        int ts = 0;
        for(Map.Entry<Integer, Map<Integer, Integer>> en : p.entrySet()){
            int x2 = en.getKey();
            if(x2 == x1){
                continue;
            }
            int s = Math.abs(x2 - x1);
            Map<Integer, Integer> a = p.get(x1);
            Map<Integer, Integer> b = en.getValue();
            int ty = y1 + s;
            ts += b.getOrDefault(y1, 0) * a.getOrDefault(ty, 0) * b.getOrDefault(ty, 0);
            int by = y1 - s;
            ts += b.getOrDefault(y1, 0) * a.getOrDefault(by, 0) * b.getOrDefault(by, 0);
        }
        return ts;
    }
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * DetectSquares obj = new DetectSquares();
 * obj.add(point);
 * int param_2 = obj.count(point);
 */