#include <iostream>
using namespace std;
void  _r(char ** s, int length) {
    char* head = *s;
    char* tail = *s + length - 1;
    char tmp;
    while (head < tail) {
        tmp = *head;
        *head = *tail;
        *tail = tmp;
        head++;
        tail--;
    }
    return;
}

void trimSpace(char ** s) {
    //去除多余空格
    char *str = *s;
    //去除首部空格
    while(*str != '\0') {
        if (*str != ' ') {
            *s = str;
            break;
        }
        str++;
    }
    str = *s;
    int i,j;
    i = 0;
    j = 0;
    //去除中间或尾部空格
    while(*(str+i) != '\0') {
        if (*(str+j) == ' ') {
            if (*(str+j+1) == ' ') {
                j++;
                continue;
            } else if (*(str+j+1) == '\0' ) {
                //去掉尾部空格
                *(str+i) = '\0';
                break;
            }
        } else if (*(str+j) == '\0' ) {
                //去掉尾部空格
                *(str+i) = '\0';
                break;
        } 
        if (*(str+i) != *(str+j)) {
            *(str+i) = *(str+j);
        }
        i++;
        j++; 
    }
}


char * reverseWords(char * s){
    if (s == NULL) {
        return NULL;
    }

    trimSpace(&s);
    cout << "trim:" << s << endl;
    int length = 0;
    char* tail = s;
    while (*tail != '\0') {
        tail++;
    }

    length = tail - s;

    _r(&s, length);
    
    char* pWordHead = s;
    char* pNext = s;

    while(pNext - s <= length) {
        if ((*(pNext) == ' ') || (pNext - s == length)) {
            _r(&pWordHead, (pNext-pWordHead));
            pNext++;
            pWordHead = pNext;
        } else {
            pNext++;
        }
    }

    return s;
}

int main() {
    cout << "cpp strive for greatness" << endl;

    char cstring[] = "  This is not! ok?  ";

    char* r;
    r = reverseWords(cstring);
    
    cout << r << endl;

    return 0;
}

