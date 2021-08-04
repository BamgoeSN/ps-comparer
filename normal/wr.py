from collections import deque
node, edge, start = [int(x) for x in input().split(' ')]

mylist = {}
for i in range(node):
    mylist[i+1] = []
for i in range(edge):
    a, b = [int(x) for x in input().split()]
    mylist[a].append(b)
    mylist[b].append(a)
for i in mylist:
    mylist[i].sort()

visited = [False for x in range(node + 1)]

dfs = []


def DFS(node):
    visited[node] = True
    dfs.append(node)
    for i in mylist[node]:
        if not visited[i]:
            DFS(i)


DFS(start)
print(dfs)


visited = [False for x in range(node + 1)]
visited[start] = True
q = deque([start])

bfs = []
while len(q):
    cur = q.popleft()
    bfs.append(cur)
    for i in mylist[cur]:
        if not visited[i]:
            visited[i] = True
            q.append(i)
print(bfs)
