#include <stdio.h>
#include <stdlib.h>

typedef struct list {
    int num;
    struct link* link;
} List;

List* insert(List* head, int no) {
    List* node = (List*)malloc(sizeof(List*));
    node->num = no;

    if (head == NULL) {
        head = node;
        node->link = head;
    } else {
        node->link = head->link;
        head->link = node;
        head = node;
    }

    return head;
}

List* print(List* head, int order) {
    List *point, *old;
    point = head->link;
    old = head;
    int cnt = 1;

    while (1) {
        if (cnt == order) {
            old->link = point->link;
            printf("%d ", point->num);
            free(point);
            break;
        }
        old = point;
        point = point->link;
        cnt++;
    }

    return head;
}

int main() {
    int times;  //몇 명 넣을것인지
    int order;  //몇 번째를 없앨것인지
    int lim = 0;

    scanf("%d %d", &times, &order);

    List* head = NULL;

    for (int i = 1; i <= times; i++)
        head = insert(head, i);

    while (lim != times) {
        head = print(head, order);
        lim++;
    }
}