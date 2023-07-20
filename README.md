# Overview of Project

This is an assignment for Week 5 in the Northwestern Masters in Data Science [MSDS-431 Data Engineering with Go](https://msdsgo.netlify.app/data-engineering-with-go/) course.

The purpose of this work is to build a web scraper in golang.  For this incarnation, the assignment objective is to take a list of URLs (provided in "testURLs.jl"), download each URL's text as an html file (in the "wikipages" folder), and save the details to a .jl file (the "goItems.jl" file).

Specifically, we were to use an already-created python example using Scrapy and ensure that our go-based output was consistent with what Scrapy provided.

The work here strongly leveraged [gocolly](github.com/gocolly/colly) as the main engine for engaging with websites.

As always, this program was developed on a Mac although both Mac and Windows executables are provided.  This is because the Canvas website which manages assignments will only accept a .exe and won't accept a Mac executable.  The Mac executable has been tested and works; the Windows executable has not been tested.

# Program Structure

Command line execution of the program is as simple as running the executable

```bash
./msds431wk5
```

The program will then complete the following steps:
1. Read the testURLs.jl file for URLs to attempt to download
2. Download the URL
3. Write the URL data as an .html file
4. Append the appropriate data to the goItems.jl file.

Note that there is logic in the program to delete the wikipages folder and goItems.jl file on each execution.  This functionality allows for better iterative testing although isn't a practical use pattern in real-life application.  In future iterations, a smarter approach would be to check for file existence and only scraping data if the file is non-existent or sufficiently out-dated.

# FYI - assignment details motivating this work

### Management Problem

A technology firm has decided to create its own online library, a knowledge base focused on its current research and development efforts into intelligent systems and robotics. Some of the information for the knowledge base will be collected from the World Wide Web. There will be an initial web crawl, followed by web page scraping and text file parsing. 

The World Wide Web has become our primary information resource. But with the web's immense size and lack of organization, finding our way to the information we need can be a frustrating, time-consuming process, even with the best general-purpose search engines at our disposal.

Fortunately, the firm believes it can collect much of the general information it needs by web scraping Wikipedia pages. Drawing on code from Ryan (2018), they have written a simple crawler/scraper in Python using Scrapy. Unfortunately, the program runs very slowly, requiring sequential searches through ordered lists of web pages.

Hearing that how fast Go is compared with Python, due largely to Go's ability to take advantage of today's multicore processors, the managers have decided to convert from Python Scrapy to Go.  Their thinking is that web crawling, scraping, and parsing can be carried out with concurrent processes across many target websites or web pages. If there are hundreds or thousands of targets to crawl, the data scientists can initiate hundreds or thousands of goroutines. Crawling, scraping, and parsing could be executed concurrently across many websites or web pages.

### Assignment Requirements 

We take on the role of the company's data scientists. We have a list of target web pages relating to the firm's research and development interests in intelligent systems and robotics. Our job is to develop a Go-based web crawler/scraper to obtain text information from the target web pages. There are four alternatives for satisfying management's information needs, ranked here from easiest to hardest:

1.  We can use the Colly framework, drawing on code and examples on GitHub: https://github.com/gocolly/colly Links to an external site.
2. We can develop a crawler/scraper web client using modules from the Go standard library, following methods described by Donovan and Kernighan (2015) or by Saha (2022), with code, exercises, and examples on GitHub: https://github.com/practicalgo Links to an external site. 
3. We can develop an HTTP client as in [2] but use it to access the Common Crawl Links to an external site.. See Patel (2020) for Python examples of this.
4. We can use the Go WebDriver for Selenium with code on GitHub: https://github.com/tebeka/selenium Links to an external site. 

Whichever path we choose, we need to get results similar to what would be provided by a Python/Scrapy-based program. We need to retrieve information (web pages) about intelligent systems and robotics from Wikipedia. And for each web page retrieved, we need to scrape away the HTML markup codes. For the purpose of this assignment, it is fine to focus on the text from each web page, ignoring images.

### Grading Guidelines (100 Total Points)

* Coding rules, organization, and aesthetics (20 points). Effective use of Go modules and idiomatic Go. Code should be readable, easy to understand. Variable and function names should be meaningful, specific rather than abstract. They should not be too long or too short. Avoid useless temporary variables and intermediate results. Code blocks and line breaks should be clean and consistent. Break large code blocks into smaller blocks that accomplish one task at a time. Utilize readable and easy-to-follow control flow (if/else blocks and for loops). Distribute the not rather than the switch (and/or) in complex Boolean expressions. Programs should be self-documenting, with comments explaining the logic behind the code (McConnell 2004, 777–817).
* Testing and software metrics (20 points). Employ unit tests of critical components, generating synthetic test data when appropriate. Generate program logs and profiles when appropriate. Monitor memory and processing requirements of code components and the entire program. If noted in the requirements definition, conduct a Monte Carlo performance benchmark.
* Design and development (20 points). Employ a clean, efficient, and easy-to-understand design that meets all aspects of the requirements definition and serves the use case. When possible, develop general-purpose code modules that can be reused in other programming projects.
* Documentation (20 points). Effective use of Git/GitHub, including a README.md Markdown file for each repository, noting the roles of programs and data and explaining how to test and use the application.
* Application (20 points). Delivery of an executable load module or application (.exe file for Windows or .app file for MacOS). The application should run to completion without issues. If user input is required, the application should check for valid/usable input and should provide appropriate explanation to users who provide incorrect input. The application should employ clean design for the user experience and user interface (UX/UI).

### Assignment Deliverables

* Text showing the link to the GitHub repository for the assignment
* README.md Markdown text file documentation for the assignment
* Zip archive of the GitHub repository
* Executable load module for the program/application (.exe for Windows or .app for MacOS)