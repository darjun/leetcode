#include <iostream>
#include <string>
#include <cassert>

using namespace std;

class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        // empty or one-length string
        if (s.size() <= 1) {
            return s.size();
        }

        int maxLength = 0;
        // int leftIndex = 0;
        // int rightIndex = 0;
        int newStartIndex = 0;
        int indice[256];
        for (int i = 0; i < 256; i++) {
            indice[i] = -1;
        }

        for (int i = 0; i < s.size(); i++) {
            if (indice[s[i]] == -1) {
                if (i - newStartIndex + 1 > maxLength) {
                    // leftIndex = newStartIndex;
                    // rightIndex = i;
                    maxLength = i - newStartIndex + 1;
                }
            } else {
                // reset indice before repeat character
                int repeatIndex = indice[s[i]];
                for (int j = newStartIndex; j <= repeatIndex; j++)
                {
                    indice[s[j]] = -1;
                }

                // set new start
                newStartIndex = repeatIndex + 1;
            }
            indice[s[i]] = i;
        }

        // cout << "max length: " << maxLength << " left index: " << leftIndex << " right index: " << rightIndex << endl;
        return maxLength;
    }
};

int main()
{
    Solution s;
    assert(s.lengthOfLongestSubstring("abcabcbb") == 3);
    assert(s.lengthOfLongestSubstring("bbb") == 1);
    assert(s.lengthOfLongestSubstring("pwwkew") == 3);
}