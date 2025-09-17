class FoodRatings {
    private static class FoodInfo{
        int rating;
        String cuisine;
        
        FoodInfo(int rating, String cuisine){
            this.rating = rating;
            this.cuisine = cuisine;
        }
    }

    private static class FoodItem{
        String name;
        int rating;

        FoodItem(String name, int rating){
            this.name = name;
            this.rating = rating;
        }

        public boolean equals(Object a){
            if (this == a){
                return true;
            }
            if(a == null || getClass() != a.getClass()){
                return false;
            }
            FoodItem temp = (FoodItem)a;
            return rating == temp.rating && Objects.equals(name, temp.name);
        }

        public int hashcode(){
            return Objects.hash(name, rating);
        }
    }

    final Map<String, FoodInfo> f1;
    final Map<String, TreeSet<FoodItem>> f2;

    public FoodRatings(String[] foods, String[] cuisines, int[] ratings) {
        f1 = new HashMap<>();
        f2 = new HashMap<>();
        for(int i = 0; i < foods.length; i++){
            String food = foods[i];
            String cuisine = cuisines[i];
            int rating = ratings[i];
            f1.put(food, new FoodInfo(rating, cuisine));
            f2.computeIfAbsent(cuisine, k -> new TreeSet<>((a, b)->{
                if (a.rating != b.rating){
                    return Integer.compare(b.rating, a.rating);
                }
                return a.name.compareTo(b.name);
            }));
            f2.get(cuisine).add(new FoodItem(food, rating));
        }
    }
    
    public void changeRating(String food, int newRating) {
        FoodInfo cur = f1.get(food);
        String cuisine = cur.cuisine;
        int old = cur.rating;
        TreeSet<FoodItem> fs = f2.get(cuisine);
        fs.remove(new FoodItem(food, old));
        fs.add(new FoodItem(food, newRating));
        cur.rating = newRating;
    }
    
    public String highestRated(String cuisine) {
        return f2.get(cuisine).first().name;
    }
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * FoodRatings obj = new FoodRatings(foods, cuisines, ratings);
 * obj.changeRating(food,newRating);
 * String param_2 = obj.highestRated(cuisine);
 */