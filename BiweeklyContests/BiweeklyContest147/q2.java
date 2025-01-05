/*Had to leave question due to multiple failed attempts, and the program not taking a code that is optimized to O(n) */

class TaskManager {
    private Map<Integer, int[]> map;
    private PriorityQueue<int[]> pq;
    private Set<Integer> deletedTasks;

    public TaskManager(List<List<Integer>> tasks) {
        pq = new PriorityQueue<>((a,b) -> a[0] != b[0] ? b[0] - a[0] : b[1] - a[1]);
        map = new HashMap<>();
        for(List<Integer> t : tasks){
            int u = t.get(0), id = t.get(1), p = t.get(2);
            map.put(id, new int[]{u,p});
            pq.offer(new int[]{p,id,u});
        }
    }

    public void add(int u, int id, int p) {
        map.put(id, new int[]{u,p});
        pq.offer(new int[]{p,id,u});
    }

    public void edit(int id, int np) {
        if(!map.containsKey(id)) return;
        int[] arr = map.get(id);
        arr[1] = np;
        map.put(id, arr);
        pq.offer(new int[]{np,id,arr[0]});
    }

    public void rmv(int id) {
        map.remove(id);
    }

    public int execTop() {
        while(!pq.isEmpty()){
            int[] top = pq.peek();
            int p = top[0], id = top[1], u = top[2];
            if(!map.containsKey(id)) {
                pq.poll();
            } else {
                int[] cur = map.get(id);
                if(cur[0]==u && cur[1]==p) {
                    pq.poll();
                    map.remove(id);
                    return u;
                }
                pq.poll();
            }
        }
        return -1;
    }
}
/**
 * Your TaskManager object will be instantiated and called as such:
 * TaskManager obj = new TaskManager(tasks);
 * obj.add(userId,taskId,priority);
 * obj.edit(taskId,newPriority);
 * obj.rmv(taskId);
 * int param_4 = obj.execTop();
 */