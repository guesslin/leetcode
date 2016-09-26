#!/usr/bin/env python
# -*- coding: utf-8 -*-


class Solution(object):
    def wordPattern(self, pattern, str):
        """
        :type pattern: str
        :type str: str
        :rtype: bool
        """
        m = {}
        s = str.split()
        if len(s) != len(pattern):
            return False
        for i in xrange(len(pattern)):
            try:
                if s[i] != m[pattern[i]]:
                    return False
            except KeyError:
                m[pattern[i]] = s[i]
        r = set(m.values())
        if len(r) < len(m.values()):
            return False
        return True
