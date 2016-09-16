#!/usr/bin/env python
# -*- coding: utf-8 -*-


class Solution(object):
    def reverseWords(self, s):
        """
        :type s: str
        :rtype: str
        """
        return ' '.join(s.split()[::-1])


if __name__ == '__main__':
    s = Solution()
    string = 'the sky is blue'
    print s.reverseWords(string)
