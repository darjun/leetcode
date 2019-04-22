import java.util.ArrayList;

public class ReverseWordsInAString {
    public static void main(String[] args) {
        Solution s = new Solution();

        // System.out.println(s.reverseWords("the sky is blue"));
        // System.out.println(s.reverseWords("  hello world!  "));
        System.out.println(s.reverseWords("a good   example"));
        System.out.println(s.reverseWords("   example "));
        System.out.println(s.reverseWords("   a "));
    }
}

class Solution {
    public String reverseWords(String s) {
        StringBuilder sb = new StringBuilder();

        int i = 0;
        while (i < s.length() && s.charAt(i) == ' ') {
            i++;
        }

        int j = s.length() - 1;
        while (j >= 0 && s.charAt(j) == ' ') {
            j--;
        }

        while (i <= j) {
            int index = s.lastIndexOf(' ', j);
            sb.append(s, index + 1, j + 1);
            sb.append(' ');

            j = index - 1;
            while (j >= 0 && s.charAt(j) == ' ') {
                j--;
            }
        }

        // remove last space
        if (sb.length() > 0) {
            sb.deleteCharAt(sb.length() - 1);
        }
        return sb.toString();
    }
}