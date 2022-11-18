#include <iostream>
#include <vector>
#include <cstdlib>
#include <omp.h>
#include <time.h>
#include <chrono>
#include <getopt.h>
#include <fstream>
#include <bits/stdc++.h>

#define MAX 1000
#define MIN 1
#define N 20

using namespace std;

int main(int argc, char *argv[]) {
    int nthreads = atoi(argv[1]);
    srand(time(0));

    fstream fout;
    fout.open("sl-times.csv", ios::out | ios::app);
    for(int i=1;i<nthreads+1;i++)
        fout << i << ",";
    fout << endl;

    vector<int> v;
    vector<int> vv[nthreads];
    int arr[N];
    int bite_size = MAX/nthreads;
    for(int i=0;i<N;i++)
        arr[i] = rand() % (MAX - MIN + 1) + MIN;

        for(int l=1;l<nthreads+1;l++)
        {
        auto start = chrono::high_resolution_clock::now();
        // #pragma omp parallel for num_threads(l) reduction(+:prime_count)
        for(int i=0;i<N;i++){
            int thread = omp_get_thread_num();
            if(arr[i] < bite_size*thread+1 && arr[i] > bite_size*(thread))
                vv[thread].push_back(arr[i]);
            vv[thread].push_back(arr[i]);
            sort(vv[thread].begin(), vv[thread].end());
        }
        for(int i=0;i<nthreads;i++)
            v.insert(v.end(), vv[i].begin(), vv[i].end());
        auto finish = chrono::high_resolution_clock::now();
        auto time = chrono::duration<double>(finish - start).count();
        //write time in csv
        fout << time << ",";
    }
    fout << endl;

    for(int i=0;i<v.size();i++)
        cout << v[i] << " ";
}