#include<iostream>
#include<queue>
using namespace std;
int N, M;
int answer = 0x3f3f3f;
bool visit[11][11][11][11] = { false, };
queue <pair<pair<int, int>, int>> R;
queue <pair<pair<int, int>, int>> B;
char MAP[11][11];
void input() {
	cin >> N >> M;
	int ry, rx, by, bx;
	for (int i = 1; i <= N; i++) {
		for (int j = 1; j <= M; j++) {
			cin >> MAP[i][j];
			if (MAP[i][j] == 'R') {
				R.push(make_pair(make_pair(i, j), 0));
				ry = i;
				rx = j;
			}
			if (MAP[i][j] == 'B') {
				by = i;
				bx = j;
				B.push(make_pair(make_pair(i, j), 0));
			}
		}
	}
	visit[ry][rx][by][bx] = true;
}

void m_left(int Rcy, int Rcx, int Bcy, int Bcx, int cnt) {
	int nry = Rcy;
	int nrx = Rcx;
	int nby = Bcy;
	int nbx = Bcx;
	while (MAP[nry][nrx -1] != '#' && MAP[nry][nrx-1] != 'O') {
		nrx--;
	}
	if (MAP[nry][nrx - 1] == 'O') {
		nrx--;
	}

	while (MAP[nby][nbx - 1] != '#' && MAP[nby][nbx - 1] != 'O') {
		nbx--;
	}
	if (MAP[nby][nbx - 1] == 'O') {
		nbx--;
	}

	if (MAP[nry][nrx] != 'O') {//겹쳤을때 
		if (nry == nby && nrx == nbx) {
			if (Bcx > Rcx) {
				nbx = nbx + 1;
				while (MAP[nby][nbx] == '#') {
					nbx++;
				}
			}
			else {
				nrx = nrx + 1;
				while (MAP[nry][nrx] == '#') {
					nrx++;
				}
			}
		}
	}
	if (Rcy == nry && Rcx == nrx && Bcx == nbx && Bcy == nby) return;
	if (visit[nry][nrx][nby][nbx]) return;
	visit[nry][nrx][nby][nbx] = true;
	R.push(make_pair(make_pair(nry, nrx), cnt+1));
	B.push(make_pair(make_pair(nby, nbx), cnt + 1));
}


void m_right(int Rcy, int Rcx, int Bcy, int Bcx, int cnt) {
	int nry = Rcy;
	int nrx = Rcx;
	int nby = Bcy;
	int nbx = Bcx;
	while (MAP[nry][nrx + 1] != '#'&& MAP[nry][nrx + 1] != 'O') {
		nrx++;
	}
	if (MAP[nry][nrx + 1] == 'O') {
		nrx++;
	}

	while (MAP[nby][nbx + 1] != '#' && MAP[nby][nbx + 1] != 'O') {
		nbx++;
	}
	if (MAP[nby][nbx + 1] == 'O') {
		nbx++;
	}

	if (MAP[nry][nrx] != 'O') {//겹쳤을때 
		if (nry == nby && nrx == nbx) {
			if (Bcx > Rcx) {
				nrx = nrx - 1;
				while (MAP[nry][nrx] == '#') {
					nrx--;
				}
			}
			else {
				nbx = nbx - 1;
				while (MAP[nby][nbx] == '#') {
					nbx--;
				}
			}
		}
	}
	if (Rcy == nry && Rcx == nrx && Bcx == nbx && Bcy == nby) return;
	if (visit[nry][nrx][nby][nbx]) return;
	visit[nry][nrx][nby][nbx] = true;
	R.push(make_pair(make_pair(nry, nrx), cnt + 1));
	B.push(make_pair(make_pair(nby, nbx), cnt + 1));
}


void m_down(int Rcy, int Rcx, int Bcy, int Bcx, int cnt) {
	int nry = Rcy;
	int nrx = Rcx;
	int nby = Bcy;
	int nbx = Bcx;
	while (MAP[nry+1][nrx] != '#' && MAP[nry+1][nrx]!='O') {
		nry++;
	}
	if (MAP[nry+1][nrx] == 'O') {
		nry++;
	}

	while (MAP[nby+1][nbx] != '#'&& MAP[nby+1][nbx]!='O') {
		nby++;
	}
	if (MAP[nby+1][nbx] == 'O') {
		nby++;
	}

	if (MAP[nry][nrx] != 'O') {//겹쳤을때 
		if (nry == nby && nrx == nbx) {
			if (Bcy > Rcy) {
				nry = nry - 1;
				while (MAP[nry][nrx] == '#') {
					nry--;
				}
			}
			else {
				nby = nby - 1;
				while (MAP[nby][nbx] == '#') {
					nby--;
				}

			}
		}
	}
	if (Rcy == nry && Rcx == nrx && Bcx == nbx && Bcy == nby) return;
	if (visit[nry][nrx][nby][nbx]) return;
	visit[nry][nrx][nby][nbx] = true;
	R.push(make_pair(make_pair(nry, nrx), cnt + 1));
	B.push(make_pair(make_pair(nby, nbx), cnt + 1));
}


void m_up(int Rcy, int Rcx, int Bcy, int Bcx, int cnt) {

	int nry = Rcy;
	int nrx = Rcx;
	int nby = Bcy;
	int nbx = Bcx;
	while (MAP[nry - 1][nrx] != '#'&& MAP[nry-1][nrx]!='O') {
		nry--;
	}
	if (MAP[nry - 1][nrx] == 'O') {
		nry--;
	}

	while (MAP[nby - 1][nbx] != '#'&& MAP[nby-1][nbx]!='0') {
		nby--;
	}
	if (MAP[nby - 1][nbx] == 'O') {
		nby--;
	}


	if (MAP[nry][nrx] != 'O') {//겹쳤을때 
		if (nry == nby && nrx == nbx) {
			if (Bcy > Rcy) {
				nby = nby + 1;
				while (MAP[nby][nbx] == '#') {
					nby++;
				}
			}
			else {
				nry = nry + 1;
				while (MAP[nry][nrx]== '#') {
					nry++;
				}
			}
		}
	}
	if (Rcy == nry && Rcx == nrx && Bcx == nbx && Bcy == nby) return;
	if (visit[nry][nrx][nby][nbx]) return;
	visit[nry][nrx][nby][nbx] = true;
	R.push(make_pair(make_pair(nry, nrx), cnt + 1));
	B.push(make_pair(make_pair(nby, nbx), cnt + 1));
}








int solve() {
	while (!R.empty()||!B.empty()) {
		int Rcy = R.front().first.first;
		int Rcx = R.front().first.second;
		int Rcnt = R.front().second;
		R.pop();

		int Bcy = B.front().first.first;
		int Bcx = B.front().first.second;
		int Bcnt = B.front().second;
		B.pop();
		if (Rcnt > 10) {
			return -1;
		}

		if (MAP[Bcy][Bcx] == 'O') {
			continue;
		}
		if (MAP[Rcy][Rcx] == 'O' && MAP[Bcy][Bcx] != 'O') {
			return Rcnt;
		}
		
		m_left(Rcy, Rcx, Bcy, Bcx, Rcnt);
		m_right(Rcy, Rcx, Bcy, Bcx, Rcnt);
		m_down(Rcy, Rcx, Bcy, Bcx, Rcnt);
		m_up(Rcy, Rcx, Bcy, Bcx, Rcnt);
	}
	return -1;
	
}

int main() {
	input();
	cout<<solve();
}