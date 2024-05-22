### Title

| 算法 | 描述 | 时间复杂度(n:点数 m:边数) |
| --- | --- | --- |
| Dijstra (朴素) | 适用于稠密图 | O(n^2) |
| Dijstra (堆优化) | 适用于稀疏图 | O(m*logn) |
| Floyd | 多源最短路 | O(n^3) |
| Bellman Frod | 边数限制（负权） | O(n*m) |
| SPFA | 理论效率很高（负权, 竞赛中正权图有时卡该算法） | O(k*m) (k一般为2) |
| Johnson | 使用带有负权的全员最短路 | O(n(n+m)logm) |

### Dijstra
dijkstra的算法思想是从以上最短距离数组中每次选择一个最近的点，将其作为下一个点，然后重新计算从起始点经过该点到其他所有点的距离，更新最短距离数据。已经选取过的点就是确定了最短路径的点，不再参与下一次计算。

position: resources/html

doc: [一篇文章讲透Dijkstra最短路径算法](https://www.cnblogs.com/goldsunshine/p/12978305.html)