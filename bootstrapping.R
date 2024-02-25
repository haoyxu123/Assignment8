library(readr)
library(pryr)
start_time <- Sys.time()
initial_mem <- mem_used()
titanic_data <- read_csv("C:/Users/haoyx/Desktop/titanic/train.csv")

str(titanic_data)

table(titanic_data$Survived)
table(titanic_data$Pclass)

bootstrap_mean <- function(data, statistic, R) {
  n <- length(data)
  means <- numeric(R)
  for (i in 1:R) {
    sample_indices <- sample(1:n, size = n, replace = TRUE)
    resampled_data <- data[sample_indices]
    means[i] <- statistic(resampled_data)
  }
  return(means)
}

proportion_survived <- function(data) {
  mean(data)
}


survived_class_1 <- titanic_data$Survived[titanic_data$Pclass == 1]

# Perform bootstrapping for 1st class passengers
R <- 1000  
bootstrap_results <- bootstrap_mean(survived_class_1, proportion_survived, R)

# Calculate 95% confidence interval
CI <- quantile(bootstrap_results, probs = c(0.025, 0.975))
print(CI)
end_time <- Sys.time()
execution_time <- end_time - start_time
final_mem <- mem_used()
memory_usage <- final_mem - initial_mem
execution_time
memory_usage
