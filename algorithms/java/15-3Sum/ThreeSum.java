import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;

public class ThreeSum {
    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.threeSum(new int[]{-1, 0, 1, 2, -1, 4}));
    }
}

class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        HashMap<Integer, Integer> m = new HashMap<>();
        Arrays.sort(nums);
        for (int i = 0; i < nums.length; i++) {
            m.put(nums[i], i);
        }

        List<List<Integer>> result = new ArrayList<>();
        HashMap<Integer, HashSet<Integer>> s = new HashMap<>();
        for (int i = 0; i < nums.length - 1; i++) {
            for (int j = i + 1; j < nums.length; j++) {
                int a = nums[i];
                int b = nums[j];
                int c = -(a + b);
                if (m.containsKey(c)) {
                    // index of c
                    int ci = m.get(c);
                    if (ci <= i || ci <= j) {
                        continue;
                    }

                    HashSet<Integer> inner = null;
                    if (s.containsKey(a)) {
                        inner = s.get(a);
                        if (inner.contains(b)) {
                            continue;
                        }
                    }

                    ArrayList<Integer> triplet = new ArrayList<>();
                    triplet.add(a);
                    triplet.add(b);
                    triplet.add(c);
                    result.add(triplet);

                    if (inner == null) {
                        inner = new HashSet<>();
                        s.put(a, inner);
                    }
                    inner.add(b);
                }
            }
        }
        return result;
    }
}