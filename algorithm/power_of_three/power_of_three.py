#!/usr/bin/env python
# -*- coding: utf-8 -*-

from math import log10


class Solution(object):
    def isPowerOfThree(self, n):
        """
        :type n: int
        :rtype: bool
        """
        return ((log10(n) / log10(3)) % 1) == 0

if __name__ == '__main__':
    s = Solution()
    print s.isPowerOfThree(81)
