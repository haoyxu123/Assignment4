start_time <- Sys.time()
N = 100 
sink("C:/Assignment4/example1.txt")
for (i in 1:N) {
    houses = read.csv(file = "C:/Assignment4/housesInput.csv", header = TRUE)
    print(summary(houses)) 
}
sink()

end_time <- Sys.time()
execution_time = end_time - start_time
print(execution_time)
