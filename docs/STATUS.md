# DATE: 26/12/25

## Current status

### *  **Scraping done.** But can ony scrape upto the html data and can't get the actual format for the revire extraction from the actual html-encoded-file need to modify it for extraction of the review for now.
### * Minimilastic fronted made with a switch and a container to input url for the analysis and returning the review to the user as written in the user-review.txt file.
### - Go scraper fetches full HTML correctly.
### - Site detection + pagination works.
### - Reddit path implemented.
### - Python pipeline operational.
### - Canonical analysis schema defined.

## **Next focus:**
### - DOM-level review extraction per site.
### - NLP integration on cleaned review text.
### - Expand scraper for Amazon/Flipkart.


### No architectural changes needed till now.
|
|
|
|
|
# DATE: 29/12/2025

## Current status

### * **Change of architecture for reddit instead of using filter inside of go JSON is being sent directly to python to let python do whatever filtering and classification is needed to do with it**
### First try to do a file I/O system through the colly (HTTP) failed and now trying to implement streaming service to tackle with the problem.