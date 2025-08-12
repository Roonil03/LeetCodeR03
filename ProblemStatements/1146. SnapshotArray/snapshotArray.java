class SnapshotArray {
    List<List<int[]>> data;
    int sid;

    public SnapshotArray(int length) {
        data = new ArrayList<>();
        for(int i = 0; i < length; i++){
            data.add(new ArrayList<>());
            data.get(i).add(new int[]{0, 0});
        }
        sid = 0;
    }
    
    public void set(int index, int val) {
        List<int[]> hist = data.get(index);
        if(hist.get(hist.size() - 1)[0] == sid){
            hist.get(hist.size() - 1)[1] = val;
        } else{
            hist.add(new int[]{sid, val});
        }
    }
    
    public int snap() {
        return sid++;
    }
    
    public int get(int index, int snap_id) {
        List<int[]> hist = data.get(index);
        int l = 0, r = hist.size() - 1;
        while(l <= r){
            int m = l + (r - l) / 2;
            if(hist.get(m)[0] <= snap_id){
                l = m + 1;
            } else{
                r = m - 1;
            }
        }
        return hist.get(r)[1];
    }
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * SnapshotArray obj = new SnapshotArray(length);
 * obj.set(index,val);
 * int param_2 = obj.snap();
 * int param_3 = obj.get(index,snap_id);
 */