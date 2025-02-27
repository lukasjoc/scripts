#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv) {
    unsigned int target = atoi(argv[1]);
    printf("Equal Divisors of: %d", target);
    printf("\n---------------------\n");
    for (int i = 0; i < target; i++) {
        if(target%i == 0) printf("\n %d", i);
    }
    printf("\n");
    return 0;
}
