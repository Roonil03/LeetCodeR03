// class MedianFinder {

//     private PriorityQueue<Integer> low;
//     private PriorityQueue<Integer> high;

//     public MedianFinder() {
//         low = new PriorityQueue<>(Collections.reverseOrder());
//         high = new PriorityQueue<>();
//     }
    
//     public void addNum(int num) {
//         low.offer(num);
//         high.offer(low.poll());
//         if (low.size() < high.size()) {
//             low.offer(high.poll());
//         }
//     }
    
//     public double findMedian() {
//         if (low.size() > high.size()) {
//             return low.peek();
//         } else {
//             return (low.peek() + high.peek()) / 2.0;
//         }
//     }
// }

// /**
//  * Your MedianFinder object will be instantiated and called as such:
//  * MedianFinder obj = new MedianFinder();
//  * obj.addNum(num);
//  * double param_2 = obj.findMedian();
//  */

class MedianFinder {
    int[] frequency = new int[200001];
    int[] jump = new int[2001];
    int mid_point = -1;
    int lower = 0;
    int higher = 0;

    public MedianFinder() {
        
    }
    
    public void addNum(int num) {
        if(mid_point == -1){
            mid_point = num + 100000;
        }        
        frequency[num + 100000]++;
        jump[num/100 + 1000]++;        
        if(num > mid_point - 100000){
            higher++;
        }
        else if(num < mid_point - 100000){
            lower++;
        }
        if(lower - higher > frequency[mid_point]){
            higher += frequency[mid_point];
            mid_point--;
            while(frequency[mid_point] == 0 && mid_point%100 != 0){
                mid_point--;
            }
            if(frequency[mid_point] == 0){
                int jump_point = mid_point/100 - 1;
                while(jump[jump_point] == 0){
                    jump_point--;
                    mid_point -= 100;
                }
            }
            while(frequency[mid_point] == 0){
                mid_point--;
            }
            lower -= frequency[mid_point];
        }
        else if(higher - lower > frequency[mid_point]){
            lower += frequency[mid_point];

            mid_point++;
            while(frequency[mid_point] == 0 && mid_point%100 != 0){
                mid_point++;
            }
            if(frequency[mid_point] == 0){
                int jump_point = mid_point/100;
                while(jump[jump_point] == 0){
                    jump_point++;
                    mid_point += 100;
                }
            }
            while(frequency[mid_point] == 0){
                mid_point++;
            }
            higher -= frequency[mid_point];
        }        
    }
    
    public double findMedian() {
        if(higher - lower == frequency[mid_point]){
            int mid = mid_point + 1;
            while(frequency[mid] == 0 && mid%100 != 0){
                mid++;
            }
            if(frequency[mid] == 0){
                int jump_point = mid/100;
                while(jump[jump_point] == 0){
                    jump_point++;
                    mid += 100;
                }
            }
            while(frequency[mid] == 0){
                mid++;
            }
            return (double)(mid + mid_point  - 200000)/2;
        }
        else if(lower - higher == frequency[mid_point]){
            int mid = mid_point - 1;
            while(frequency[mid] == 0 && mid%100 != 0){
                mid--;
            }
            if(frequency[mid] == 0){
                int jump_point = mid/100 - 1;
                while(jump[jump_point] == 0){
                    jump_point--;
                    mid -= 100;
                }
            }
            while(frequency[mid] == 0){
                mid--;
            }
            return (double)(mid + mid_point  - 200000)/2;
        }
        else{
            return (double)mid_point - 100000;
        }        
    }
}