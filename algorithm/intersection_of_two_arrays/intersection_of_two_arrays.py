#!/usr/bin/env python
# -*- coding: utf-8 -*-


class Solution(object):
    def intersection(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: List[int]
        """
        return list(set(nums1) & set(nums2))


if __name__ == '__main__':
    s = Solution()
    print s.intersection(range(10), [x for x in range(10) if x % 2 == 0])
