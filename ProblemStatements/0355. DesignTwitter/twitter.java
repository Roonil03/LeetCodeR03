class Twitter {
    static int stamp = 0;
    Map<Integer, Set<Integer>> follows;
    Map<Integer, List<Tweet>> tweets;

    class Tweet{
        int id;
        int time;

        public Tweet(int id, int time){
            this.id = id;
            this.time = time;
        }
    }

    public Twitter() {
        follows = new HashMap<>();
        tweets = new HashMap<>();  
    }
    
    public void postTweet(int userId, int tweetId) {
        tweets.putIfAbsent(userId, new ArrayList<>());
        tweets.get(userId).add(new Tweet(tweetId, stamp++));
    }
    
    public List<Integer> getNewsFeed(int userId) {
        PriorityQueue<Tweet> pq = new PriorityQueue<>((a, b) -> b.time - a.time);
        List<Tweet> ut = tweets.get(userId);
        if(ut != null){
            for(Tweet t : ut){
                pq.offer(t);
            }
        }
        Set<Integer> f = follows.get(userId);
        if(f != null){
            for(int i : f){
                List<Tweet> ft = tweets.get(i);
                if(ft != null){
                    for(Tweet t : ft){
                        pq.offer(t);
                    }
                }
            }
        }
        List<Integer> res = new ArrayList<>();
        int count = 0;
        while(!pq.isEmpty() && count < 10){
            res.add(pq.poll().id);
            count++;
        }
        return res;
    }
    
    public void follow(int followerId, int followeeId) {
        if (followerId != followeeId){
            follows.putIfAbsent(followerId, new HashSet<>());
            follows.get(followerId).add(followeeId);
        }
    }
    
    public void unfollow(int followerId, int followeeId) {
        if(follows.containsKey(followerId)){
            follows.get(followerId).remove(followeeId);
        }
    }
}

/**
 * Your Twitter object will be instantiated and called as such:
 * Twitter obj = new Twitter();
 * obj.postTweet(userId,tweetId);
 * List<Integer> param_2 = obj.getNewsFeed(userId);
 * obj.follow(followerId,followeeId);
 * obj.unfollow(followerId,followeeId);
 */