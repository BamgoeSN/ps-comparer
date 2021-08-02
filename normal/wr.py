import sys
input = sys.stdin.readline

N = int(input())
schedule = [list(map(int,input().rstrip().split(" "))) for _ in range(N)]
dp = [0]*30
for i in range(N):
    if i+1<schedule[i][0]:
        a = max(schedule[i][1],dp[i+schedule[i][0]+1])
        for t in range(i+schedule[i][0]+1,N+2):
            dp[t] = max(a,dp[t])
    else:
        a = max(max(dp[:i+2])+schedule[i][1],dp[i+schedule[i][0]+1])
        for t in range(i+schedule[i][0]+1,N+2):
            dp[t] = max(a,dp[t])

print(max(dp[:N+2]))
