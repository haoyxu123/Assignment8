library(boot)
library(dplyr)
library(pryr)

start_time <- Sys.time()
initial_mem <- mem_used()
data <- read.csv("C:/Users/haoyx/Desktop/titanic/train.csv")

filtered_data <- data %>% filter(Pclass == 1) %>% select(Survived)

filtered_data$Survived <- as.numeric(filtered_data$Survived)

statistic_function <- function(data, indices) {
  mean(data[indices])
}

set.seed(123) # For reproducibility
boot_results <- boot(data = filtered_data$Survived, statistic = statistic_function, R = 1000)

ci <- boot.ci(boot_results, type = c("perc"))

# Print the bootstrap results and confidence interval
print(boot_results)
cat("95% Confidence Interval for the survival rate of 1st class passengers:", ci$perc[4], ci$perc[5], "\n")
end_time <- Sys.time()
execution_time <- end_time - start_time
final_mem <- mem_used()
memory_usage <- final_mem - initial_mem
execution_time
memory_usage
