import java.util.Arrays;
import java.util.Collections;
import java.util.Comparator;
import java.util.Map;
import java.util.TreeMap;

public class SlidingWindowMaximum {
    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(Arrays.toString(s.maxSlidingWindow(new int[]{1, 3, -1, -3, 5, 3, 6, 7}, 3)));
    }
}

class Solution {
    public int[] maxSlidingWindow(int[] nums, int k) {
        if (k == 0) {
            return new int[]{};
        }

        int[] sliding = new int[nums.length - k + 1];
        int index = 0;
        
        TreeMap<Integer, Integer> t = new TreeMap<>(Collections.reverseOrder());

        // first sliding window
        for (int i = 0; i < k; i++) {
            if (t.containsKey(nums[i])) {
                t.put(nums[i], t.get(nums[i]) + 1);
            } else {
                t.put(nums[i], 1);
            }
        }
        sliding[index++] = t.firstKey();

        for (int i = k; i < nums.length; i++) {
            int keyToRemove = nums[i - k];
            int count = t.get(keyToRemove);
            if (count == 1) {
                t.remove(keyToRemove);
            } else {
                t.put(keyToRemove, count - 1);
            }
            if (t.containsKey(nums[i])) {
                t.put(nums[i], t.get(nums[i]) + 1);
            } else {
                t.put(nums[i], 1);
            }

            sliding[index++] = t.firstKey();
        }

        return sliding;
    }
}