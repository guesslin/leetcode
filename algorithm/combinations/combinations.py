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
        return list(combinations([x+1 for x in xrange(n)], k))

if __name__ == '__main__':
    s = Solution()
    print s.combine(5, 3)
