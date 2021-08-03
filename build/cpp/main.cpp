#include <stdio.h>

void push(int *, int *, char *, int *, int *, int *, int *);        //arr,stack,sign,num,cur,top,scur
void pop(int *, int *, int *, char *, int *, int *, int *, int *);  //arr,stack,save,sign,cur,top,scur,acur

int main() {
    int times;          //몇 번 넣을건지
    int arr[100001];    //original
    int stack[100001];  //for stacking
    int save[100001];   //for camparing
    char sign[10000];   //for saving sign
    int num = 1;
    int cur = 0;   //for pointing number in arr
    int top = -1;  //for pointing number in stack
    int scur = 0;  //for pointing pos in sign
    int acur = 0;  //for pointing pos in save
    int cnt = 0;   //판별용
    int k = 0;

    scanf("%d", &times);

    for (int i = 0; i < times; i++)
        scanf("%d", &arr[i]);

    while (cur < times) {
        if (num > arr[cur])
            break;
        push(arr, stack, sign, &num, &cur, &top, &scur);
        pop(arr, stack, save, sign, &cur, &top, &scur, &acur);
    }

    for (int i = 0; i < times; i++) {
        if (arr[i] != save[i])
            cnt++;
    }

    if (cnt != 0)
        printf("NO\n");

    else
        while (sign[k] == '+' || sign[k] == '-') {
            printf("%c\n", sign[k]);
            k++;
        }
}

void push(int *carr, int *cstack, char *csign, int *cnum, int *ccur, int *ctop, int *cscur) {
    int pnum = *cnum;
    int pcur = *ccur;
    int ptop = *ctop;
    int pscur = *cscur;

    while (pnum <= carr[pcur]) {
        cstack[++ptop] = pnum++;
        csign[pscur++] = '+';
    }

    *cnum = pnum;
    *ccur = pcur;
    *ctop = ptop;
    *cscur = pscur;
}

void pop(int *carr, int *cstack, int *csave, char *csign, int *ccur, int *ctop, int *cscur, int *cacur) {
    int pcur = *ccur;
    int ptop = *ctop;
    int pscur = *cscur;
    int pacur = *cacur;

    while (cstack[ptop] >= carr[pcur]) {
        csave[pacur++] = cstack[ptop];
        csign[pscur++] = '-';
        if (carr[pcur] == cstack[ptop])
            pcur++;
        ptop--;
        if (ptop == -1)
            break;
    }

    *ccur = pcur;
    *ctop = ptop;
    *cscur = pscur;
    *cacur = pacur;
}
