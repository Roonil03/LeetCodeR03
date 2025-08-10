class TimeMap {

    Map<String, List<Pair>> m;

    public TimeMap() {
        m = new HashMap<>();
    }
    
    public void set(String key, String value, int timestamp) {
        m.computeIfAbsent(key, k -> new ArrayList<>()).add(new Pair(timestamp, value));
    }
    
    public String get(String key, int timestamp) {
        if(!m.containsKey(key)){
            return "";
        }
        List<Pair> lt = m.get(key);
        int l = 0, r = lt.size() - 1;
        String res = "";
        while(l <= r){
            int m = l + (r - l) / 2;
            if(lt.get(m).timestamp <= timestamp){
                res = lt.get(m).value;
                l = m + 1;
            } else{
                r = m - 1;
            }
        }
        return res;
    }

    class Pair{
        int timestamp;
        String value;
        Pair(int timestamp, String value){
            this.timestamp = timestamp;
            this.value = value;
        }
    }
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * TimeMap obj = new TimeMap();
 * obj.set(key,value,timestamp);
 * String param_2 = obj.get(key,timestamp);
 */