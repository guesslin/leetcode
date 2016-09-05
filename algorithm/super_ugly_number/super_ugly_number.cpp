#include <iostream>
#include <vector>

using std::vector;


int min(vector<int>& nums) {
	int m = nums[0];
	for (unsigned int i = 1; i < nums.size(); i++) {
		if (m > nums[i]) {
			m = nums[i];
		}
	}
	return m;
}

int nthSuperUglyNumber(int n, vector<int>& primes) {
	vector<int> ugly(n,n);
	vector<int> m = primes;
	vector<int> index(primes.size(), 0);
	ugly[0] = 1;
	for (int i = 1; i < n; i++) {
		int minV = min(m);
		ugly[i] = minV;
		for (unsigned int j = 0; j < index.size(); j++) {
			if (minV == m[j]) {
				index[j]++;
				m[j] = primes[j] * ugly[index[j]];
			}
		}
	}
	return ugly[n-1];
}

int main(void) {
	vector<int> primes;
	primes.push_back(2);
	primes.push_back(7);
	primes.push_back(13);
	primes.push_back(19);
	std::cout << nthSuperUglyNumber(12, primes) << std::endl;
	return 0;
}
