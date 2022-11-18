#include <iostream>
#include <vector>
#include <cstdlib>
#include <omp.h>
#include <time.h>

#define MAX 10000000000000
#define MIN 100000
#define NN 100

using namespace std;

int main(int argc, char *argv[]) {
    int nthreads = atoi(argv[1]);
    srand(time(0));

    int prime_count = 0;

    #pragma omp parallel for num_threads(nthreads) reduction(+:prime_count)
    for(int i=0;i<NN;i++){
        int n = rand() % (MAX - MIN + 1) + MIN;
        int sup = 1;
        for(int j=2;j<n;j++){
            if(n%j==0){
                sup=0;
                break;
            }
        }
        if(sup)
            prime_count++;
    }

    cout<<"There are: " << prime_count << " prime numbers in the " << NN << " Generated Numbers" << endl;
}