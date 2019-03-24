public class FirstMissingPositive {
    public static void main(String args[]) {
        Solution s = new Solution();
        System.out.println(s.firstMissingPositive(new int[] { 1, 2, 0 }));
        System.out.println(s.firstMissingPositive(new int[] { 3, 4, -1, 1 }));
        System.out.println(s.firstMissingPositive(new int[] { 7, 8, 9, 11, 12 }));
    }
}

class Solution {
    public int firstMissingPositive(int[] nums) {
        int i = 0;
        while (i < nums.length) {
            // ignore negetive & those greater than array length.
            if (nums[i] <= 0 || nums[i] > nums.length) {
                i++;
                continue;
            }

            // i-th positive exists & at right position.
            if (nums[i] == i + 1) {
                i++;
                continue;
            }

            // move nums[i] to right position.
            // if nums[i] not in [1, n], don't bother to move.
            // if nums[i] is already at right position, don't move too.
            while (nums[i] > 0 && nums[i] <= nums.length && nums[i] != nums[nums[i]-1]) {
                int tmp = nums[nums[i]-1];
                nums[nums[i]-1] = nums[i];
                nums[i] = tmp;
            }

            i++;
        }
        
        for (i = 0; i < nums.length; i++) {
            if (nums[i] != i + 1) {
                return i + 1;
            }
        }
        return nums.length+1;
    }
}