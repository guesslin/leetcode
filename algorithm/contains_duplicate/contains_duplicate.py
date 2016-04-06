#!/usr/bin/env python
# -*- coding: utf-8 -*-


class Solution(object):
    def containsDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        if len(nums) != len(set(nums)):
            return True
        return False
