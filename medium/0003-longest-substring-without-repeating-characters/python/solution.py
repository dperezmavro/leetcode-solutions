class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        max_len = 0
        curr_str = ""
        counts = dict()

        for c in s:
            curr_str += c
            if c in counts:
                counts[c] += 1
                while counts[c] > 1:
                    cc = curr_str[0]
                    curr_str = curr_str[1:]
                    counts[cc] -= 1
                    if counts[cc] == 0:
                        del counts[cc]
            else:
                counts[c] = 1
                if len(curr_str) > max_len:
                    max_len = len(curr_str)

        return max_len
