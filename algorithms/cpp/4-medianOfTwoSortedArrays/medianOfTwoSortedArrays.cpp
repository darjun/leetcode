#include <iostream>
#include <vector>
#include <cassert>
#include <cmath>

using namespace std;

class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        // assert size of nums1 ge than size of nums2
        int m = nums1.size();
        int n = nums2.size();
        if (m > n) {
            return findMedianSortedArrays(nums2, nums1);
        }

        int imin = 0;
        int imax = m;
        int halfLen = (m + n + 1) / 2;
        while (imin <= imax) {
            int i = (imin + imax) / 2;
            int j = halfLen - i;

            if (i > 0 && j < n && nums1[i-1] > nums2[j]) {
                imax = i - 1;
            } else if (j > 0 && i < m && nums2[j-1] > nums1[i]) {
                imin = i + 1;
            } else {
                int maxLeft = 0;
                if (i == 0) {
                    maxLeft = nums2[j - 1];
                } else if (j == 0) {
                    maxLeft = nums1[i - 1];
                } else {
                    maxLeft = max(nums1[i - 1], nums2[j - 1]);
                }

                if ((m+n) % 2==1)
                    return maxLeft;

                int minRight = 0;
                if (i == m) {
                    minRight = nums2[j];
                } else if (j == n) {
                    minRight = nums1[i];
                } else {
                    minRight = min(nums1[i - 1], nums2[j - 1]);
                }

                return (maxLeft + minRight) / 2.0;
            }
        }

        return 0.0;
    }
};

int main()
{
    Solution s;
    vector<int> nums1{1, 3};
    vector<int> nums2{2};
    assert(s.findMedianSortedArrays(nums1, nums2) == 2.0);

    vector<int> nums3{1, 2};
    vector<int> nums4{3, 4};
    assert(s.findMedianSortedArrays(nums3, nums4) == 2.5);
}