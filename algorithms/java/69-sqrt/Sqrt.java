public class Sqrt {
    public static void main(String[] args) {
        Solution s = new Solution();

        // System.out.println(s.mySqrt(4));
        // System.out.println(s.mySqrt(8));
        System.out.println(s.mySqrt(2147395599));
        System.out.println(s.mySqrt(2147483647));
    }
}

class Solution {
    public int mySqrt(int x) {
        if (x == 0) {
            return 0;
        }

        if (x <= 3) {
            return 1;
        }

        int low = 1;
        int high = x / 2;
        
        while (low <= high) {
            int mid = (low + high) >>> 1;
            if (mid == x / mid) {
                return mid;
            } else if (mid > x / mid) {
                high = mid - 1;
            } else {
                low = mid + 1;
            }
        }

        return high;
    }
}