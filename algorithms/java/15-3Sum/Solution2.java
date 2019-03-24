import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Solution2 {
    public List<List<Integer>> threeSum(int[] nums) {        
        List<List<Integer>> result = new ArrayList<>();
        if (nums.length < 3) {
            return result;
        }

        Arrays.sort(nums);
        for (int i = 0; i < nums.length - 2; i++) {
            // cut branch
            if (i > 0 && nums[i] == nums[i-1]) {
                continue;
            }

            int a = nums[i];
            int bi = i + 1;
            int ci = nums.length - 1;

            while (bi < ci) {
                int b = nums[bi];
                int c = nums[ci];
                int s = a + b + c;
                if (s > 0) {
                    --ci;
                } else if (s < 0) {
                    ++bi;
                } else {
                    ArrayList<Integer> triplet = new ArrayList<>();
                    triplet.add(a);
                    triplet.add(b);
                    triplet.add(c);
                    result.add(triplet);

                    // skip same b & c
                    ++bi;
                    while (bi < ci && nums[bi] == nums[bi - 1]) {
                        ++bi;
                    }

                    --ci;
                    System.out.println(ci);
                    while (bi < ci && nums[ci] == nums[ci + 1]) {
                        --ci;
                    }
                }
            }
        }
        return result;
    }
}