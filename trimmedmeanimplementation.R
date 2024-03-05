# Generate sample data
set.seed(123) # for reproducibility
integer_sample <- sample(1:1000, 100, replace = TRUE)
float_sample <- runif(100, min = 1, max = 1000)

# Calculate trimmed mean for integers
trimmed_mean_integers <- mean(integer_sample, trim = 0.05)

# Calculate trimmed mean for floats
trimmed_mean_floats <- mean(float_sample, trim = 0.05)

# Output the results
print(paste("Trimmed mean for integers:", trimmed_mean_integers))
print(paste("Trimmed mean for floats:", trimmed_mean_floats))
