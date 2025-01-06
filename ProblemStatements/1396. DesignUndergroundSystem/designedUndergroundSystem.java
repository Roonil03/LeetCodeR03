class UndergroundSystem {
    Map<Integer, Pair<String, Integer>> ci;
    Map<String, Pair<Integer, Integer>> tt;

    UndergroundSystem() {
        ci = new HashMap<>();
        tt = new HashMap<>();
    }

    void checkIn(int id, String sn, int t) {
        ci.put(id, new Pair<>(sn, t));
    }

    void checkOut(int id, String sn, int t) {
        Pair<String, Integer> chk = ci.get(id);
        String ss = chk.getKey();
        int st = chk.getValue();
        int time = t - st;
        String route = ss + ":" + sn;
        Pair<Integer, Integer> tr = tt.getOrDefault(route, new Pair<>(0, 0));
        int total = tr.getKey() + time;
        int count = tr.getValue() + 1;
        tt.put(route, new Pair<>(total, count));
        ci.remove(id);
    }

    double getAverageTime(String ss, String es) {
        String route = ss + ":" + es;
        Pair<Integer, Integer> tr = tt.get(route);
        int total = tr.getKey();
        int count = tr.getValue();
        return (double) total / count;
    }

    class Pair<K, V> {
        K k;
        V v;

        Pair(K k, V v) {
            this.k = k;
            this.v = v;
        }

        K getKey() {
            return k;
        }

        V getValue() {
            return v;
        }
    }
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * UndergroundSystem obj = new UndergroundSystem();
 * obj.checkIn(id,stationName,t);
 * obj.checkOut(id,stationName,t);
 * double param_3 = obj.getAverageTime(startStation,endStation);
 */