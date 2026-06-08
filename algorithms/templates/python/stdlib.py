"""
Python 算法竞赛常用标准库速查。
"""

from collections import deque, defaultdict, Counter
import heapq
import bisect
import math
import sys
from itertools import permutations, combinations, product, accumulate
from functools import lru_cache, cmp_to_key
from operator import itemgetter


# ============================================================
# 1. collections：常用数据结构
# ============================================================

# deque：双端队列，常用于队列、栈、BFS
q = deque([1, 2, 3])
q.append(4)        # 右侧加入
q.appendleft(0)    # 左侧加入
q.pop()            # 右侧弹出
q.popleft()        # 左侧弹出

# defaultdict：默认字典，常用于建图、分组
graph = defaultdict(list)
graph[1].append(2)
graph[1].append(3)

# Counter：计数器，常用于统计频率
cnt = Counter([1, 2, 2, 3, 3, 3])
print(cnt[3])          # 输出 3
print(cnt.most_common(2))  # 出现次数最多的前 2 个元素


# ============================================================
# 2. heapq：堆 / 优先队列
# ============================================================

# Python 默认是小根堆
h = [9, 6, 3, 3]
heapq.heapify(h)
heapq.heappush(h, 3)
heapq.heappush(h, 1)
heapq.heappush(h, 2)
print(heapq.heappop(h))  # 输出 1

# 大根堆：存入相反数
h = []
heapq.heappush(h, -3)
heapq.heappush(h, -1)
heapq.heappush(h, -2)
print(-heapq.heappop(h))  # 输出 3


# ============================================================
# 3. bisect：二分查找
# ============================================================

arr = [1, 3, 3, 5, 7]

# bisect_left：第一个 >= x 的位置
print(bisect.bisect_left(arr, 3))   # 输出 1

# bisect_right：第一个 > x 的位置
print(bisect.bisect_right(arr, 3))  # 输出 3

# insort：插入元素并保持有序
bisect.insort(arr, 4)
print(arr)  # [1, 3, 3, 4, 5, 7]


# ============================================================
# 4. itertools：排列组合与迭代工具
# ============================================================

# permutations：排列，顺序不同算不同结果
print(list(permutations([1, 2, 3], 2)))

# combinations：组合，顺序不同不重复计算
print(list(combinations([1, 2, 3], 2)))

# product：笛卡尔积，常用于枚举状态
print(list(product([0, 1], repeat=3)))

# accumulate：前缀和
arr = [1, 2, 3, 4]
prefix = list(accumulate(arr, initial=0))
print(prefix)  # [1, 3, 6, 10]


# ============================================================
# 5. math：数学函数
# ============================================================

print(math.gcd(12, 18))       # 最大公约数，输出 6
print(math.lcm(12, 18))       # 最小公倍数，输出 36
print(math.sqrt(16))          # 浮点平方根，输出 4.0
print(math.isqrt(17))         # 整数平方根，输出 4
print(math.factorial(5))      # 阶乘，输出 120
print(math.comb(5, 2))        # 组合数 C(5, 2)，输出 10
print(math.perm(5, 2))        # 排列数 P(5, 2)，输出 20
print(math.ceil(7 / 3))       # 向上取整，输出 3


# ============================================================
# 6. functools：缓存、比较器等工具
# ============================================================

# lru_cache：记忆化搜索，常用于递归 DP
@lru_cache(None)
def fib(n):
    if n <= 1:
        return n
    return fib(n - 1) + fib(n - 2)

print(fib(10))  # 输出 55
fib.cache_clear()


# cmp_to_key：把比较函数转换为排序 key
def cmp(a, b):
    return (a > b) - (a < b)

arr = [3, 1, 2]
arr.sort(key=cmp_to_key(cmp))
print(arr)  # [1, 2, 3]


# ============================================================
# 7. operator：常用取值工具
# ============================================================

pairs = [(1, 3), (2, 2), (3, 1)]

# itemgetter(1)：按照元组的第 2 个元素排序
pairs.sort(key=itemgetter(1))
print(pairs)  # [(3, 1), (2, 2), (1, 3)]


# ============================================================
# 8. sys：输入输出与递归深度
# ============================================================

# 快速输入
# n = int(sys.stdin.readline())
# arr = list(map(int, sys.stdin.readline().split()))

# 设置递归深度，DFS 时常用
sys.setrecursionlimit(10**6)


# ============================================================
# 9. 内置函数：排序、枚举、聚合
# ============================================================

arr = [3, 1, 4, 2]

print(sorted(arr))              # 升序排序
print(sorted(arr, reverse=True)) # 降序排序
print(min(arr))                 # 最小值
print(max(arr))                 # 最大值
print(sum(arr))                 # 求和

# enumerate：同时获得下标和值
for i, x in enumerate(arr):
    print(i, x)

# zip：并行遍历多个列表
a = [1, 2, 3]
b = [4, 5, 6]

for x, y in zip(a, b):
    print(x, y)


# ============================================================
# 10. set：集合操作
# ============================================================

s = set()
s.add(1)       # 添加元素
s.add(2)
s.remove(1)    # 删除元素，元素不存在会报错
s.discard(3)   # 删除元素，元素不存在不会报错

a = {1, 2, 3}
b = {3, 4, 5}

print(a | b)  # 并集
print(a & b)  # 交集
print(a - b)  # 差集
print(a ^ b)  # 对称差集