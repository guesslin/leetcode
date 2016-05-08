#!/usr/bin/env python
# -*- coding: utf-8 -*-


class Solution(object):
    def combine(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: List[List[int]]
        """
        from itertools import combinations
        result = []

        def add(x): return x+1
        for c in list(combinations(range(n), k)):
            result.append(map(add, c))

        return result

if __name__ == '__main__':
    s = Solution()
    print s.combine(5, 3)
