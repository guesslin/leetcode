#!/usr/bin/env python
# -*- coding: utf-8 -*-


class Solution:
    # @param n, an integer
    # @return an integer
    def hammingWeight(self, n):
        s = 0
        while n:
            s += n % 2
            n = int(n/2)
        return s


if __name__ == '__main__':
    import sys
    f = Solution()
    if len(sys.argv) == 2:
        print f.hammingWeight(int(sys.argv[1]))
