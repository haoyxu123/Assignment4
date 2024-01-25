import pandas as pd
import time 
start_time = time.perf_counter()
N = 100 
with open('C:/Assignment4/example2.txt', 'wt') as outfile:
    for i in range(N):
        houses = pd.read_csv("C:/Assignment4/housesInput.csv")
        outfile.write(houses.describe().to_string(header=True, index=True))
        outfile.write("\n")
end_time = time.perf_counter()
execution_time = end_time - start_time
print(f"Execution time is {execution_time} seconds")




