public class ClimingStairs {
    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.climbStairs(2));
        System.out.println(s.climbStairs(3));
    }
}

class Solution {
    public int climbStairs(int n) {
        int p1 = 1; // 1 way to clime 1 steps, 1
        int p2 = 2; // 2 ways to clime 2 steps, 1 + 1 and 2
        if (n == 1) {
            return p1;
        }
        if (n == 2) {
            return p2;
        }

        for (int i = 3; i <= n; i++) {
            int p3 = p1 + p2;
            p1 = p2;
            p2 = p3;
        }

        return p2;
    }
}