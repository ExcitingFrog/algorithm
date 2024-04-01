### Title

| 算法 | 描述 | 时间复杂度(n:点数 m:边数) |
| --- | --- | --- |
| Dijstra (朴素) | 适用于稠密图 | O(n^2) |
| Dijstra (堆优化) | 适用于稀疏图 | O(m*logn) |
| Floyd | 多源最短路 | O(n^3) |
| Bellman Frod | 边数限制（负权） | O(n*m) |
| SPFA | 理论效率很高（负权, 竞赛中正权图有时卡该算法） | O(k*m) (k一般为2) |
| Johnson | 使用带有负权的全员最短路 | O(n(n+m)logm) |