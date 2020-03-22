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

void trimSpace(char ** s, int originLen) {
    int slowOffset = 0;
    int fastOffset = 0;

    while (fastOffset < originLen) {
        while (fastOffset < originLen && *(*s+fastOffset) == ' ') fastOffset++;
        while (fastOffset < originLen && *(*s+fastOffset) != ' ') {
            *(*s+slowOffset) = *(*s+fastOffset);
            slowOffset++;
            fastOffset++;
        }
        while (fastOffset < originLen && *(*s+fastOffset) == ' ') fastOffset++;
        if (fastOffset < originLen) {
            *(*s+slowOffset) = ' '; 
            slowOffset++;
        }    
    }
    *(*s+slowOffset) ='\0';
}


char * reverseWords(char * s){
    if (s == NULL) {
        return NULL;
    }


    cout << "orig:" << s << "|" << endl;
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

    trimSpace(&s, length);
    cout << "trim:" << s << "|" << endl;
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

