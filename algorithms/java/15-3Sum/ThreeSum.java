public class ThreeSum {
    public static void main(String[] args) {
        Solution1 s1 = new Solution1();
        System.out.println(s1.threeSum(new int[] { -1, 0, 1, 2, -1, 4 }));
        
        Solution2 s2 = new Solution2();
        System.out.println(s2.threeSum(new int[] { -1, 0, 1, 2, -1, 4 }));
        System.out.println(s2.threeSum(new int[] { 0, 0, 0, 0 }));
    }
}