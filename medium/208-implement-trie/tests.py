import unittest
from solution import Trie
from typing import List


class TestCase:
    def __init__(self, input: List[str], want: List[List[int]]):
        self.want = want
        self.input = input


class TestSolution(unittest.TestCase):
    def test_equalSubstring(self):
        def t1():
            t = Trie()
            t.insert("apple")
            self.assertEqual(True, t.search("apple"))
            self.assertEqual(False, t.search("app"))
            self.assertEqual(True, t.startsWith("app"))
            t.insert("app")
            self.assertEqual(True, t.search("app"))


        def t2():
            t = Trie()
            t.insert("a")
            self.assertEqual(True, t.search("a"))
            self.assertEqual(True, t.startsWith("a"))
        
        def t3():
            t = Trie()
            t.insert("ab")
            self.assertEqual(False, t.search("a"))
            self.assertEqual(True, t.startsWith("a"))
        
        t1()
        t2()
        t3()
