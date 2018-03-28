#include <iostream>

using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode* head = NULL;
        ListNode* tail = NULL;
        int carry = 0;
        while (l1 != NULL || l2 != NULL) {
            int digit = (l1 ? l1->val : 0) + (l2 ? l2->val : 0) + carry;
            carry = digit / 10;
            digit %= 10;

            ListNode* newNode = new ListNode(digit);
            if (head == NULL) {
                head = newNode;
            }
            if (tail != NULL) {
                tail->next = newNode;
            }
            tail = newNode;

            if (l1) l1 = l1->next;
            if (l2) l2 = l2->next;
        }

        // 最后如果有进位，生成最高位
        if (carry) {
            tail->next = new ListNode(carry);
        }
        return head;
    }
};

void destroyList(ListNode* l)
{
    while (l != NULL) {
        ListNode* tmp = l->next;
        delete(l);
        l = tmp;
    }
}

void printList(ListNode* l)
{
    while (l != NULL) {
        std::cout << l->val << " ";
        l = l->next;
    }
    std::cout << std::endl;
}

int main()
{
    ListNode* l1 = new ListNode(2);
    l1->next = new ListNode(4);
    l1->next->next = new ListNode(3);

    ListNode* l2 = new ListNode(5);
    l2->next = new ListNode(6);
    l2->next->next = new ListNode(4);

    Solution s;
    ListNode* result = s.addTwoNumbers(l1, l2);

    printList(result);

    destroyList(l1);
    destroyList(l2);
    destroyList(result);

    return 0;
}