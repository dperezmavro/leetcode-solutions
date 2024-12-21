from typing import List
from collections import defaultdict


class Solution:
    def closeStrings(self, word1: str, word2: str) -> bool:
        if len(word1) != len(word2):
            return False

        letters_1 = defaultdict(int)
        letters_2 = defaultdict(int)
        for idx, c in enumerate(word1):
            letters_1[c] += 1
            letters_2[word2[idx]] += 1
        
        if len(set(letters_1.keys()).difference(set(letters_2.keys()))):
            return False
        
        counts_1 = defaultdict(int)
        counts_2 = defaultdict(int)
        for i in letters_1.values():
            counts_1[i] += 1
        
        for i in letters_2.values():
            counts_2[i] += 1
        
        return counts_1 == counts_2
