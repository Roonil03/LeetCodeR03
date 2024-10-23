#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

class largestNumber {
public:
    static bool compare(int s1, int s2) {
        return std::to_string(s1) + std::to_string(s2) > std::to_string(s2) + std::to_string(s1);
    }
    
    std::string largestNumber(std::vector<int>& nums) {
        std::sort(nums.begin(), nums.end(), compare);
        if (nums[0] == 0) {
            return "0";
        } else {
            std::string ans;
            for (int n : nums) {
                ans += std::to_string(n);
            }
            return ans;
        }
    }
};

int main() {
    largestNumber ln; 
    std::vector<int> nums = {10, 2, 9, 39, 0};

    std::string result = ln.largestNumber(nums);
    std::cout << "Largest Number: " << result << std::endl;

    return 0;
}
