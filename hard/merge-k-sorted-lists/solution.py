from typing import List, Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class BinTreeNode:
    def __init__(self, val: int = 0):
        self.val = val
        self.left: BinTreeNode = None
        self.right: BinTreeNode = None

    def add(self, v: int):
        if v <= self.val:
            if self.left != None:
                self.left.add(v)
            else:
                self.left = BinTreeNode(v)
        else:
            if self.right != None:
                self.right.add(v)
            else:
                self.right = BinTreeNode(v)

    def getValues(self) -> List[int]:
        answer = []
        if self.left != None:
            answer = [x for x in self.left.getValues()]

        answer.append(self.val)

        if self.right != None:
            for x in self.right.getValues():
                answer.append(x)
        return answer


class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        answer = []
        if len(lists) == 0:
            return answer

        b: BinTreeNode = None
        for ls in lists:
            if ls == None:
                continue

            ln: ListNode = ls
            while ln != None:
                if b == None:
                    b = BinTreeNode(ln.val)
                else:
                    b.add(ln.val)
                ln = ln.next
        if b != None:
            values = b.getValues()
            if len(values) == 0:
                return None
            root: ListNode = ListNode(values[0])
            ln: ListNode = root
            for v in values[1:]:
                tmp = ListNode(v)
                ln.next = tmp
                ln = tmp
            return root
        return None

    def mergeKListsSane(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        answer = []
        if len(lists) == 0:
            return answer

        b: BinTreeNode = None
        for ls in lists:
            if ls == None:
                continue

            for n in ls:
                if b == None:
                    b = BinTreeNode(n)
                else:
                    b.add(n)
        if b != None:
            return b.getValues()
        return answer


s = Solution()
print(s.mergeKListsSane([]))
print(s.mergeKListsSane([[]]))
print(s.mergeKListsSane([[1, 4, 5], [1, 3, 4], [2, 6]]))
print(s.mergeKLists([[1, 4, 5], [1, 3, 4], [2, 6]]))
