/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* reverseBetween(ListNode* head, int left, int right) {
        ListNode *leftNode = head;
        ListNode *ans = NULL;
        for(int i=1;i<left;i++){
            ans = leftNode;
            leftNode = leftNode->next;
        }
        
        ListNode* curr = leftNode;
        ListNode* prev = NULL;
        ListNode* nextt = NULL;
        
        ListNode* rightNode = NULL;
        for(int i=left;i<=right;i++){
            nextt = curr->next;
            if (i==right){
                rightNode = nextt;
            }
            curr->next = prev;
            prev = curr;
            curr = nextt;
        }
        
        if(left==1){
            head = prev;
        } else {
            ans->next = prev;
        }
        for(int i=left;i<=right;i++){
            if (i==right){
                prev->next = rightNode;
                break;
            }
            prev = prev->next;
        }
        return head;
    }
};