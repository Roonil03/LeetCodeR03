class RandomizedCollection {

    List<Integer> list;
    Map<Integer, Set<Integer>> map;
    Random rand;

    public RandomizedCollection() {
        list = new ArrayList<>();
        map = new HashMap<>();
        rand = new Random();
    }
    
    public boolean insert(int val) {
        boolean np = !map.containsKey(val);
        map.computeIfAbsent(val, k -> new HashSet<>()).add(list.size());
        list.add(val);
        return np;
    }
    
    public boolean remove(int val) {
        if(!map.containsKey(val) || map.get(val).isEmpty()){
            return false;
        }
        int id = map.get(val).iterator().next();
        map.get(val).remove(id);
        int a = list.size() - 1;
        int b = list.get(a);
        list.set(id, b);
        if(id != a){
            map.get(b).add(id);
            map.get(b).remove(a);
        }
        list.remove(a);
        if(map.get(val).isEmpty()){
            map.remove(val);
        }
        return true;
    }
    
    public int getRandom() {
        return list.get(rand.nextInt(list.size()));
    }
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * RandomizedCollection obj = new RandomizedCollection();
 * boolean param_1 = obj.insert(val);
 * boolean param_2 = obj.remove(val);
 * int param_3 = obj.getRandom();
 */