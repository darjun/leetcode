public class CanMakePalindroneFromSubstring {
  public static void main(String[] args) {
    System.out.println(
        new Solution().canMakePaliQueries("abcda",
            new int[][] { new int[] { 3, 3, 0 }, new int[] { 1, 2, 0 }, new int[] { 0, 3, 1 },
                new int[] { 0, 3, 2 } , new int[]{0,4,1}}));
  }
}