class SummaryRanges {
    TreeMap<Integer, Integer> in;

    public SummaryRanges() {
        in = new TreeMap<>();
    }
    
    public void addNum(int value) {
        Integer s = in.floorKey(value);
        if(s != null && in.get(s) >= value){
            return;
        }
        Integer l = in.lowerKey(value);
        Integer r = in.higherKey(value);
        int ns = value;
        int ne = value;
        if(l != null && in.get(l) + 1 >= value){
            ns = l;
            ne = Math.max(in.get(l), value);
            in.remove(l);
        }
        if(r != null && ne + 1 >= r){
            ne = in.get(r);
            in.remove(r);
        }
        in.put(ns, ne);
    }
    
    public int[][] getIntervals() {
        int[][] res = new int[in.size()][2];
        int i = 0;
        for(Map.Entry<Integer, Integer> e : in.entrySet()){
            res[i][0] = e.getKey();
            res[i][1] = e.getValue();
            i++;
        }
        return res;
    }
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * SummaryRanges obj = new SummaryRanges();
 * obj.addNum(value);
 * int[][] param_2 = obj.getIntervals();
 */