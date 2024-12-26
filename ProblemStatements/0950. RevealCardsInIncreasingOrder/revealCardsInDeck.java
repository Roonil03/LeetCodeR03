class Solution {
    public int[] deckRevealedIncreasing(int[] deck) {
        Arrays.sort(deck);
        Deque<Integer> q = new ArrayDeque<>();
        for(int i = 0; i < deck.length;i++){
            q.add(i);
        }
        int[] res = new int[deck.length];
        for(int i : deck){
            res[q.poll()] = i;
            if(!q.isEmpty()){
                q.add(q.poll());
            }
        }
        return res;
    }
}