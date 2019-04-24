public class StringToInteger {
    public static void main(String[] args) {
        Solution s = new Solution();
        // System.out.println(s.myAtoi("42"));
        // System.out.println(s.myAtoi("4193 with words"));
        // System.out.println(s.myAtoi("words and 987"));
        // System.out.println(s.myAtoi("-91283472332"));
        // System.out.println(s.myAtoi("91283472332"));
        System.out.println(s.myAtoi("-2147483649"));
    }
}

class Solution {
    public int myAtoi(String str) {
        int startPos = 0;
        while (startPos < str.length() && str.charAt(startPos) == ' ') {
            startPos++;
        }

        if (startPos >= str.length()) {
            return 0;
        }

        int sign = 1;
        if (str.charAt(startPos) == '+') {
            startPos++;
        } else if (str.charAt(startPos) == '-') {
            startPos++;
            sign = -1;
        }

        int result = 0;
        int maxLeftPart = Integer.MAX_VALUE / 10; // 214748364
        int minLeftPart = Integer.MIN_VALUE / 10; // -214748364
        int maxRightDigit = 7; // 2147483647
        int minRightDigit = 8; // -2147483648
        while (startPos < str.length()) {
            char ch = str.charAt(startPos++);
            // digit out of range
            if (ch < '0' || ch > '9') {
                break;
            }

            int digit = ch - '0';
            if (sign == 1) {
                // number out of range
                if (result > maxLeftPart || (result == maxLeftPart && digit > maxRightDigit)) {
                    result = Integer.MAX_VALUE;
                    break;
                }

                result = result * 10 + digit;
            } else {
                // number out of range
                if ((result != 0 && result < minLeftPart) || (result == minLeftPart && digit > minRightDigit)) {
                    result = Integer.MIN_VALUE;
                    break;
                }

                result = result * 10 - digit;
            }
        }

        return result;
    }
}