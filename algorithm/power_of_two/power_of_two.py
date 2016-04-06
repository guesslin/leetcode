class Solution(object):
    def isPowerOfTwo(self, n):
        """
        :type n: int
        :rtype: bool
        """
        pot = [2**x for x in xrange(64)]
        return n in pot
