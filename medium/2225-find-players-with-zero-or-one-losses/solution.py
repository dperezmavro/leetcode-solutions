from typing import List
from collections import defaultdict


class Solution:
    def findWinners(self, matches: List[List[int]]) -> List[List[int]]:
        los_count = defaultdict(int)
        winners = set()
        losers = set()
        for i in matches:
            w = i[0]
            l = i[1]
            winners.add(w)
            losers.add(l)
            los_count[l] += 1
            
            if l in winners:
                winners.remove(l)
            if w in los_count:
                winners.remove(w)
            if los_count[l] > 1:
                losers.remove(l)

        w = list(winners)
        w.sort()
        l = list(losers)
        l.sort()
        return [w, l]
