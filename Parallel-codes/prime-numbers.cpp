#include <iostream>
#include <vector>
#include <cstdlib>
#include <omp.h>
#include <time.h>
#include <chrono>
#include <getopt.h>
#include <fstream>

#define MAX 10000000000
#define MIN 1000
#define NN 1000

using namespace std;

int main(int argc, char *argv[]) {
    int nthreads = atoi(argv[1]);
    srand(time(0));

    fstream fout;
    fout.open("pn-times.csv", ios::out | ios::app);
    for(int i=1;i<nthreads+1;i++)
        fout << i << ",";
    fout << endl;

    int prime_count = 0;
    int m=1, nthreadsn=1;

    for(int k=0;k<5;k++)
    {
        for(int l=1;l<nthreads+1;l++)
        {
            auto start = chrono::high_resolution_clock::now();
            #pragma omp parallel for num_threads(l) reduction(+:prime_count)
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
            auto time = chrono::duration<double>(finish - start).count();
            nthreadsn++;
            fout << time << ",";
            //write time in csv
        }
        fout << endl;
        m=m*10;
    }

    cout<<"There are: " << prime_count << " prime numbers in the " << NN << " Generated Numbers" << endl;
}