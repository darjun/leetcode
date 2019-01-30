#include <iostream>
#include <vector>
#include <map>

using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        map<int, int> indices;
        for (int i = 0; i < nums.size(); i++) {
            auto iter = indices.find(nums[i]);
            if (iter == indices.end()) {
                indices[target-nums[i]] = i;
            }
            else {
                return vector<int>({iter->second, i});
            }
        }
        return vector<int>();
    }
};

int main()
{
    Solution s;
    vector<int> nums{2, 7, 11, 15};
    int target = 9;
    for (auto index: s.twoSum(nums, target)) {
        std::cout << index << " ";
    }
    std::cout << std::endl;
    return 0;
}