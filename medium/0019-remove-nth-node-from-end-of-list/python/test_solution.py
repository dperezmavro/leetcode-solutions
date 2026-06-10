import unittest
from typing import List
from solution import Solution, ListNode


class TestCase:
    def __init__(
        self,
        input: List[int],
        want: List[int],
        k: int,
    ):
        self.k = k
        self.want = want

        head = ListNode(val=input[0])
        node = head
        for i in input[1:]:
            nn = ListNode(val=i)
            node.next = nn
            node = nn
        self.head = head


class TestRunner(unittest.TestCase):
    def test_solution(self):
        tests = [
            TestCase(
                input=[1, 2, 3, 4, 5],
                want=[1, 2, 3, 5],
                k=2,
            ),
            TestCase(
                input=[1],
                want=[],
                k=1,
            ),
            TestCase(
                input=[1,2],
                want=[1],
                k=1,
            ),
            TestCase(
                input=[1,2],
                want=[2],
                k=2,
            ),
        ]

        for t in tests:
            s = Solution()
            arr = []
            h = s.removeNthFromEnd(t.head, t.k)
            while h:
                arr.append(h.val)
                h = h.next
            self.assertEqual(t.want, arr)


if __name__ == "__main__":
    unittest.main()
