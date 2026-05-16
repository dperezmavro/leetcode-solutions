from typing import List

class Solution:
    def compress(self, chars: List[str]) -> int:
        count_formatter = lambda s: [x for x in "{}".format(s)]
        curr_c = ""
        curr_count = 0
        res = []
        for c in chars:
            if curr_c == "":
                curr_c = c
                res.append(curr_c)
            if c == curr_c:
                curr_count += 1
            # different char
            else:
                if curr_count > 1:
                    res.extend(count_formatter(curr_count))
                curr_c = c
                res.append(curr_c)
                curr_count = 1

        if curr_count > 1:
            res.extend(count_formatter(curr_count))

        for idx, v in enumerate(res):
            chars[idx] = v
        return len(res)
