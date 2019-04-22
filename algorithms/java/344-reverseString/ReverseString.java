public class ReverseString {
    public static void main(String[] args) {
        Solution s = new Solution();

        char[] s1 = new char[] { 'h', 'e', 'l', 'l', 'o' };
        s.reverseString(s1);
        System.out.println(s1);
    }
}
class Solution {
    public void reverseString(char[] s) {
        for (int i = 0, j = s.length - 1; i < j; i++, j--) {
            char temp = s[i];
            s[i] = s[j];
            s[j] = temp;
        }
    }
}