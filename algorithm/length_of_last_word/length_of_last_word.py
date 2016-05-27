#!/usr/bin/env python
# -*- coding: utf-8 -*-


class Solution(object):
    def lengthOfLastWord(self, s):
        """
        :type s: str
        :rtype: int
        """
        try:
            return len(s.split()[-1])
        except IndexError:
            return 0


if __name__ == '__main__':
    s = Solution()
    print s.lengthOfLastWord("Hello world")
    print s.lengthOfLastWord("a ")
    print s.lengthOfLastWord("b   a    ")
