public class MajorityElement {
    public static void main(String args[]) {
        Solution s = new Solution();
        System.out.println(s.majorityElement(new int[]{3, 2, 3}));
        System.out.println(s.majorityElement(new int[]{2, 2, 1, 1, 1, 2, 2}));
    }
}

class Solution {
    public int majorityElement(int[] nums) {
        int me = nums[0];
        int count = 1;
        
        for (int i = 1; i < nums.length; i++) {
            if (me == nums[i]) {
                count++;
            } else {
                count--;
                if (count == 0) {
                    me = nums[i];
                    count = 1;
                }
            }
        }

        return me;
    }
}