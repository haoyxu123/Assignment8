# Assignment8

This Assignment compares Golang and R with the same package. Evaluating the possibilities of Go for the firm to use Go in place of R and determining how much the firm can save in cloud computing costs. 
The computer-intensive statistical method I selected is bootstrapping. Bootstrapping is a statistical method used to estimate the sampling distribution of an estimator by sampling with replacement from the original dataset. In R, the package called 'boot' provides extensive functionality for bootstrapping and related resampling methods. This is the URL of the 'boot' package.
https://CRAN.R-project.org/package=boot

In order to improve the performance of the Go implementation of the selected statistical method, I employed Testing, benchmarking, and software profiling. In the Golang program,  there are functions (TestBootstrapSample,TestCalculateConfidenceInterval and BenchmarkBootstrapSample) provided for testing and benchmarking the BootstrapSample function. These can be used with Go's built-in testing framework to ensure the function behaves as expected and to measure its performance. The code also utilizes CPU profiling and memory profiling to measure the performance of the program. 

#Compare the R and Golang to see the difference between execution time and memory usage

I have uploaded the final results to the repository. In R, the result is 0.5648148 for 2.5% and 0.69444 for 97.5%. The execution time is 4.049 secs and memory usage is 8.62MB. In the Golang program, the 95% CI is 0.5648 and 0.69 which is the same as R. The execution time is 176.2312ms which is 0.1762312secs, Alloc is 7MiB, TotalAlloc is 1MiB, Sys is 7MiB, and numGC is 0. By comparing these, we can see Golang performs better than R both in execution time and memory usage. 

#Describe your efforts in finding R and Go packages for the method. Review your process of building the Go implementation. Review your experiences with testing, benchmarking, software profiling, and logging.

The reason I chose Bootstrapping is when I was in my internship, I used this package many times and in last quarter's statistical class, I just reviewed this statistical method. So I am very familiar with this. In other way, I think bootstrapping is very common to use in data science so I want to know how to perform that in Golang. 
I do not think I had any troubles when I wrote the bootstrapping by using Golang. I found the dataset in Kaggle performs bootstrapping for 1st class passengers and calculated a 95% confidence interval. Then I did testing and benchmarking to evaluate the performance of the program and also implemented profiling to aim for program optimization. 

#Under what circumstances would it make sense for the firm to use Go in place of R for the selected statistical method? Select a cloud provider of infrastructure as a service (IaS). Note the cloud costs for virtual machine (compute engine) services. What percentage of cloud computing costs might be saved with a move from R to Go?

The cloud provider I chose is AWS Lambda, the price of it is 0.0000166667 for every GB-second. The cost of R is the Cost of R=GB-seconds for R×Price per GB-second which is approximately $0.0000006184 per execution. The cost of using Golang is 0.0017×0.0000166667 which is approximately $0.0000000283 per execution. So the percentage of money saved is approximately 95.42% savings in costs per execution. ​When a firm is facing large datasets or very complicated algorithms that might take a long execution time or take up a lot of memory usage, the firm should use Go in place of R and this will help to save a lot of money. 
