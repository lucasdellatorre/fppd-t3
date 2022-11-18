#include <iostream>
#include <vector>
#include <cstdlib>
#include <omp.h>
#include <time.h>
#include <chrono>
#include <getopt.h>

#define MAX 10000000000000
#define MIN 100000
#define NN 1000

using namespace std;

int main(int argc, char *argv[]) {
    int nthreads = atoi(argv[1]);
    srand(time(0));

    int prime_count = 0;
    int m=1, nthreadsn=1;

    for(int k=0;k<5;k++)
    {
        auto start = chrono::high_resolution_clock::now();
        #pragma omp parallel for num_threads(nthreadsn) reduction(+:prime_count)
        for(int i=0;i<NN*m;i++){
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
        auto finish = chrono::high_resolution_clock::now();
        m=m*10;
        nthreadsn++;
        auto time = chrono::duration<double>(finish - start).count();
        //write time in csv
    }

    cout<<"There are: " << prime_count << " prime numbers in the " << NN << " Generated Numbers" << endl;
}