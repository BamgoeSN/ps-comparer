from collections import deque

N = int(input())
arr = [list(input()) for _ in range(N)]  # 그림 담는 arr
vector = ((-1, 0), (1, 0), (0, 1), (0, -1))  # 왼쪽, 오른쪽, 아래, 위
visited = None

# 유효성 판단: 범위내에 있고 간 적이 없으며 같은색일경우


def validate(x, y, nowColor):
    return 0 <= x < N and 0 <= y < N and not visited[y][x] and arr[y][x] == nowColor


def bfs():
    global visited

    visited = [[False]*N for _ in range(N)]  # 방문한 곳 초기화
    cnt = 0  # 개수

    # N*N을 모두 순회
    for y in range(N):
        for x in range(N):
            # 방문한 적이 없다면 방문
            if not visited[y][x]:
                deq = deque()
                nowColor = arr[y][x]  # 지금의 색
                deq.append((x, y, nowColor))
                visited[y][x] = True  # 방문한걸로 표시
                cnt += 1  # 해당영역에 대해 카운트해줌

                # 다음위치로 이동
                while deq:
                    x, y, nowColor = deq.popleft()
                    for vx, vy in vector:
                        nx, ny = (x+vx, y+vy)  # nextX, nextY
                        if validate(nx, ny, nowColor):
                            visited[ny][nx] = True
                            deq.append((nx, ny, nowColor))
    return cnt


# 일반인이 봤을 경우
a1 = bfs()

# 적록색약을 위한 R -> G
for y in range(N):
    for x in range(N):
        if arr[y][x] == 'R':
            arr[y][x] = 'G'

# 적록 색약이 봤을 경우
a2 = bfs()

print(a1, a2)
