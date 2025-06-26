class SmallestInfiniteSet {
    TreeSet<Integer> s;
    int c;
    
    public SmallestInfiniteSet() {
        s = new TreeSet<>();
        c = 1;    
    }
    
    public int popSmallest() {
        if(!s.isEmpty()){
            return s.pollFirst();
        }
        return c++;
    }
    
    public void addBack(int num) {
        if (num < c){
            s.add(num);
        }
    }
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * SmallestInfiniteSet obj = new SmallestInfiniteSet();
 * int param_1 = obj.popSmallest();
 * obj.addBack(num);
 */